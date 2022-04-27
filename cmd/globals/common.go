package globals

type Common struct {
	Bot struct {
		Ws        string `yaml:"ws"`
		VerifyKey string `yaml:"key"`
		Debug     bool   `yaml:"debug"`
	}
	DB string `yaml:"db"`
}
