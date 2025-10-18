package config

type ZugangConfig struct {
	Server struct {
		Host string
		Port int
	}
	OAuth struct {
		GrantTypesAll bool
		GrantTypes    []string
	}
}
