package plugin

type Config struct {
	Say string `mapstructure:"say"`
}

func (c *Config) InitDefaults() {
	if c.Say == "" {
		c.Say = "hello community 🔥"
	}
}
