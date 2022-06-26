package plugin

import (
	"github.com/roadrunner-server/api/v2/plugins/config"
	"github.com/roadrunner-server/errors"
	"go.uber.org/zap"
)

const name = "custom_plugin"

type Plugin struct {
	log *zap.Logger
	cfg *Config
}

func (p *Plugin) Init(cfg config.Configurer, log *zap.Logger) error {
	if !cfg.Has(name) {
		return errors.E(errors.Disabled)
	}

	err := cfg.UnmarshalKey(name, p.cfg)
	if err != nil {
		return err
	}

	p.log = new(zap.Logger)
	*p.log = *log

	return nil
}

func (p *Plugin) Serve() chan error {
	errCh := make(chan error, 1)

	p.log.Info(p.cfg.Say)

	return errCh
}

func (p *Plugin) Stop() error {
	return nil
}

func (p *Plugin) Name() string {
	return name
}
