package config

type Mysql struct {
	User     string
	Password string
	Host     string
	Post     int
	Database string
}
type Nacos struct {
	ServerAddr string
	Namespace  string
	Port       int
	DataId     string
	Group      string
}
type AppConfig struct {
	Mysql
	Nacos
}
