package middlewares

import "github.com/kmou424/resonance-dataserver/server/engine"

func Register() {
	engine.E.Use(getErrorHandler())
	engine.E.Use(getAuthVerifier())
}
