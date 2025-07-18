package config

type JwtConfig struct {
	APISecret     string
	RefreshSecret string
	AccessTTL     int
	RefreshTTL    int
}
