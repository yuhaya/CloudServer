package conf

var GlobalConfig map[string]string

func init() {
	GlobalConfig = make(map[string]string)
	GlobalConfig["mysql"] = "root:@tcp(localhost:3306)/cloud?charset=utf8"
}
