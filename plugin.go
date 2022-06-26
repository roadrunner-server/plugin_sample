package plugin

import (
	"github.com/roadrunner-server/api/v2/plugins/config"
	"go.uber.org/zap"
)

type Plugin struct {
	log *zap.Logger
}

func (p *Plugin) Init(cfg config.Configurer, log *zap.Logger) error {
	p.log = log
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
