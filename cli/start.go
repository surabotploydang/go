package cli

import (
	"git.onespace.co.th/project-v/project-v-go-backend/route"
	"git.onespace.co.th/project-v/project-v-go-backend/service"
	"git.onespace.co.th/project-v/project-v-go-backend/setting"
	"github.com/patcharp/golib/server"
	"github.com/sirupsen/logrus"
)

func StartServer(s *server.Server) error {
	if err := service.InitServices(); err != nil {
		return err
	}
	route.RegisterApiEp(s.Ctx())
	logrus.Infoln("Server start ( Production:", setting.Production, ")")
	return s.Run()
}
