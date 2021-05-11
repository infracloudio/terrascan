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

package config

import (
	"strings"

	"github.com/accurics/terrascan/pkg/mapper/convert"
	fn "github.com/accurics/terrascan/pkg/mapper/iac-providers/arm/functions"
	"github.com/accurics/terrascan/pkg/mapper/iac-providers/arm/types"
)

const (
	sku              = "sku"
	name             = "name"
	family           = "family"
	capacity         = "capacity"
	enableNonSslPort = "enableNonSslPort"
	patchSchedule    = "Microsoft.Cache/redis/patchSchedules"
	scheduleEntries  = "scheduleEntries"
	dayOfWeek        = "dayOfWeek"
	startHourUtc     = "startHourUtc"
)

// redisCacheConfig holds config for azurerm_redis_cache
type redisCacheConfig struct {
	config

	Capacity         float64 `json:"capacity"`
	Family           string  `json:"family"`
	SKUName          string  `json:"sku_name"`
	EnableNonSSLPort bool    `json:"enable_non_ssl_port"`

	PatchSchedule struct {
		DayOfWeek    string  `json:"day_of_week"`
		StartHourUTC float64 `json:"start_hour_utc"`
	} `json:"patch_schedule,omitempty"`
}

// RedisCacheConfig returns config for azurerm_redis_cache
func RedisCacheConfig(r types.Resource, params map[string]interface{}) interface{} {
	cf := redisCacheConfig{
		config: config{
			Location: fn.LookUp(nil, params, r.Location).(string),
			Name:     fn.LookUp(nil, params, r.Name).(string),
			Tags:     r.Tags,
		},
		EnableNonSSLPort: fn.LookUp(nil, params, convert.ToString(r.Properties, enableNonSslPort)).(bool),
	}

	s := convert.ToMap(r.Properties, sku)
	cf.SKUName = fn.LookUp(nil, params, convert.ToString(s, name)).(string)
	cf.Family = fn.LookUp(nil, params, convert.ToString(s, family)).(string)
	cf.Capacity = fn.LookUp(nil, params, convert.ToString(s, capacity)).(float64)

	for _, rr := range r.Resources {
		if strings.EqualFold(rr.Type, patchSchedule) {
			sch := convert.ToMap(rr.Properties, scheduleEntries)
			cf.PatchSchedule.DayOfWeek = fn.LookUp(nil, params, convert.ToString(sch, dayOfWeek)).(string)
			cf.PatchSchedule.StartHourUTC = fn.LookUp(nil, params, convert.ToString(sch, startHourUtc)).(float64)
		}
	}
	return cf
}
