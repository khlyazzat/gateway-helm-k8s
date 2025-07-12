package config

type JwtConfig struct {
	RefreshSecret string `mapstructure:"JWT_REFRESH_SECRET"`
	AccessTTL     int    `mapstructure:"JWT_ACCESS_TTL"`
	RefreshTTL    int    `mapstructure:"JWT_REFRESH_TTL"`
}
