package datasourcemodel

import (
	tfconfig "github.com/hashicorp/terraform-plugin-testing/config"
)

func (d *DatabasesModel) WithLimit(rows int) *DatabasesModel {
	return d.WithLimitValue(
		tfconfig.ObjectVariable(map[string]tfconfig.Variable{
			"rows": tfconfig.IntegerVariable(rows),
		}),
	)
}

func (d *DatabasesModel) WithRowsAndFrom(rows int, from string) *DatabasesModel {
	return d.WithLimitValue(
		tfconfig.ObjectVariable(map[string]tfconfig.Variable{
			"rows": tfconfig.IntegerVariable(rows),
			"from": tfconfig.StringVariable(from),
		}),
	)
}
