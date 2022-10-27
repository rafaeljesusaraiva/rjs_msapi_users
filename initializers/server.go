package initializers

import (
	"github.com/julienschmidt/httprouter"
	"github.com/rafaeljesusaraiva/rjs_msapi_users/token"
)

type ServerStruct struct {
	// DB Store
	Config     Config
	TokenMaker token.Maker
	Router     *httprouter.Router
}

var SRV *ServerStruct

func GetSRV() *ServerStruct {
	return SRV
}

func InitializeServer(config Config, tokenMaker token.Maker, router *httprouter.Router) {
	SRV = &ServerStruct{
		Config:     config,
		TokenMaker: tokenMaker,
		Router:     router,
	}
}
