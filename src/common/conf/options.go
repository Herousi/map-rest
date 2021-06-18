package conf

var (
	Options *Config
)

// 公共配置
type Config struct {
	DbURL      string
	DbDriver   string
	RedisURL   string
	RedisDB    int
	RedisTag   string
	Prefix     string
	ArgBool    bool
	ArgInt     int
	PgDbURL    string
	PgDbDriver string
}
