package xapi

import (
	"go.uber.org/fx"
	"lab-api/api/xapi/admin"
	"lab-api/api/xapi/devops"
	"lab-api/api/xapi/page"
	"lab-api/api/xapi/role"
	"lab-api/api/xapi/system"
)

var Provides = fx.Options(
	system.Provides,
	page.Provides,
	role.Provides,
	admin.Provides,
	devops.Provides,
)