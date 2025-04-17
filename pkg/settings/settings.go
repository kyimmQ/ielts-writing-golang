package settings

type Config struct {
	Server  ServerConfig  `mapstructure:"server" json:"server"`
	MongoDB MongoDBConfig `mapstructure:"mongodb" json:"mongodb"`
	JWT     JWTConfig     `mapstructure:"jwt" json:"jwt"`
}

type ServerConfig struct {
	Port string `mapstructure:"port" json:"port"`
}

type MongoDBConfig struct {
	URI          string `mapstructure:"uri" json:"uri"`
	DatabaseName string `mapstructure:"database_name" json:"database_name"`
}

type JWTConfig struct {
	SecretKey string `mapstructure:"secret_key"`
	Expiry    string `mapstructure:"expiry"`
}
