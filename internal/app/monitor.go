package app

import (
	"github.com/linzhengen/gin-admin-with-frontend-sample/internal/app/config"
	"github.com/google/gops/agent"
)

// InitMonitor 初始化服务监控
func InitMonitor() error {
	if c := config.GetGlobalConfig().Monitor; c.Enable {
		err := agent.Listen(agent.Options{Addr: c.Addr, ConfigDir: c.ConfigDir, ShutdownCleanup: true})
		if err != nil {
			return err
		}
	}
	return nil
}
