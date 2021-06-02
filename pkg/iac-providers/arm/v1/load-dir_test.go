/*
    Copyright (C) 2020 Accurics, Inc.

	Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at

		http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.
*/

package armv1

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"syscall"
	"testing"

	"go.uber.org/zap"
	"gopkg.in/src-d/go-git.v4"
	gitConfig "gopkg.in/src-d/go-git.v4/config"
	"gopkg.in/src-d/go-git.v4/plumbing"

	"github.com/accurics/terrascan/pkg/iac-providers/output"
)

const (
	repoURL  = "https://github.com/accurics/KaiMonkey.git"
	branch   = "master"
	basePath = "test_data"
	provider = "arm"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func TestLoadIacDir(t *testing.T) {
	table := []struct {
		name    string
		dirPath string
		armv1   ARMV1
		want    output.AllResourceConfigs
		wantErr error
	}{
		{
			name:    "empty config",
			dirPath: "./testdata/testfile",
			armv1:   ARMV1{},
			wantErr: fmt.Errorf("no directories found for path ./testdata/testfile"),
		},
		{
			name:    "load invalid config dir",
			dirPath: "./testdata",
			armv1:   ARMV1{},
			wantErr: nil,
		},
		{
			name:    "invalid dirPath",
			dirPath: "not-there",
			armv1:   ARMV1{},
			wantErr: &os.PathError{Err: syscall.ENOENT, Op: "lstat", Path: "not-there"},
		},
		{
			name:    "key-vault",
			dirPath: "./testdata/key-vault",
			armv1:   ARMV1{},
			wantErr: nil,
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			_, gotErr := tt.armv1.LoadIacDir(tt.dirPath, false)
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("unexpected error; gotErr: '%v', wantErr: '%v'", gotErr, tt.wantErr)
			}
		})
	}
}

func TestARMMapper(t *testing.T) {
	root := filepath.Join(basePath, provider)
	dirList := make([]string, 0)
	err := filepath.Walk(root, func(filePath string, fileInfo os.FileInfo, err error) error {
		if fileInfo != nil && fileInfo.IsDir() {
			dirList = append(dirList, filePath)
		}
		return err
	})

	if err != nil {
		t.Error(err)
	}

	armv1 := ARMV1{}
	for i := 1; i < len(dirList); i++ {
		dir := dirList[i]
		t.Run(dir, func(t *testing.T) {
			_, err := armv1.LoadIacDir(dir, false)
			if err != nil {
				t.Error(err)
			}
		})
	}
}

func setup() {
	err := downloadArtifacts()
	if err != nil {
		zap.S().Fatal(err)
	}
}

func shutdown() {
	os.RemoveAll(basePath)
}

func downloadArtifacts() error {
	os.RemoveAll(basePath)

	r, err := git.PlainClone(basePath, false, &git.CloneOptions{
		URL: repoURL,
	})
	if err != nil {
		return err
	}

	w, err := r.Worktree()
	if err != nil {
		return err
	}

	err = r.Fetch(&git.FetchOptions{
		RefSpecs: []gitConfig.RefSpec{"refs/*:refs/*", "HEAD:refs/heads/HEAD"},
	})
	if err != nil {
		return err
	}

	err = w.Checkout(&git.CheckoutOptions{
		Branch: plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", branch)),
		Force:  true,
	})
	if err != nil {
		return err
	}
	return nil
}
