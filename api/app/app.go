package app

import (
	"bss-beta-backend/admin"
	auditlog "bss-beta-backend/audit_log"
	"bss-beta-backend/data"
	"bss-beta-backend/home"
	"bss-beta-backend/staff"

	"fx.prodigy9.co/app"
)

func Build() *app.Builder {
	return app.Build().
		Name("BSS Beta API").
		DefaultAPIMiddlewares().
		Mount(home.App).
		Mount(admin.App).
		Mount(staff.App).
		Mount(data.App).
		Mount(auditlog.App)
}
