package mapper

import (
	"github.com/accurics/terrascan/mapper/core"
	"github.com/accurics/terrascan/mapper/iac-providers/arm"
)

// NewMapper returns a mapper based on IaC provider.
func NewMapper(iacType string) core.Mapper {
	switch iacType {
	case "arm":
		return arm.Mapper()
	}
	return nil
}
