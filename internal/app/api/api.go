package api

import (
	"github.com/gin-gonic/gin"

	"github.com/kavladst/ads_api/internal/app/configuration"
	"github.com/kavladst/ads_api/internal/app/storage"
)

type Api struct {
	router  *gin.Engine
	Storage *storage.Storage
	Config  *configuration.Configuration
}

func New() (*Api, error) {
	apiConfig, err := configuration.New()
	if err != nil {
		return nil, err
	}

	apiStorage, err := storage.New(apiConfig)
	if err != nil {
		return nil, err
	}

	newApi := &Api{Storage: apiStorage, Config: apiConfig}
	newApi.initRouter()

	return newApi, nil
}

func (a *Api) Run() error {
	return a.router.Run(a.Config.AppHost + ":" + a.Config.AppPort)
}
