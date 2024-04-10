package routers

import "github.com/kmou424/resonance-dataserver/server/handlers"

func registerAdminRouters() {
	admin := apiGroup.Group("/admin")
	admin.PUT("/uuid", handlers.AddUUID)
	admin.PATCH("/uuid", handlers.UpdateUUID)
	admin.DELETE("/uuid", handlers.DeleteUUID)
}
