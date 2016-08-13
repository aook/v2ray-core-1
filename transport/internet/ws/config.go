package ws

type Config struct {
	ConnectionReuse bool
	Path            string
	Pto             string
	Cert            string
	PrivKey         string
}

func (this *Config) Apply() {
	effectiveConfig = this
}

var (
	effectiveConfig = &Config{
		ConnectionReuse: true,
		Path:            "",
		Pto:             "",
	}
)
