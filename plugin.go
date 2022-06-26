package plugin

import (
	"github.com/roadrunner-server/api/v2/plugins/config"
	"go.uber.org/zap"
)

type Plugin struct {
	log *zap.Logger
}

func Init(cfg config.Configurer, log *zap.Logger) error {
	return nil
}

func (p *Plugin) Serve() chan error {
	errCh := make(chan error, 1)
	p.log.Info("foo")
	return errCh
}

func (p *Plugin) Stop() error {
	return nil
}
