package initialize

import (
	"github.com/cgund98/realdebrid-cli/internal/biz/realdebrid"
	"github.com/spf13/viper"
)

func RealDebridController() *realdebrid.Controller {
	apiToken := viper.GetString("api.token")
	baseUrl := viper.GetString("api.base_url")

	return realdebrid.NewController(apiToken, baseUrl)
}
