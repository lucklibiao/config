package config

type Bytes = []byte

type Config struct{
	ZK string
	LogPath string
	Domain string
}

func (c *Config) Init(file string) (err error){
	return Parse(file,c)
}
