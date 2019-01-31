package common

//Config File functions toggle on/off
type Function struct{
	Name		string `json:"name"`
	Enabled 	bool `json:"enabled"`
}

type Config struct{
	ApiPort			string `json:"api_port"`
	Statsd_host		string `json:"statsd_host"`
	Ssl 			Ssl `json:"ssl"`
	Functions		[]Function
}
type Ssl struct{
	Enabled 	bool   `json:"enabled"`
	CertFile 	string `json:"cert_file"`
	KeyFile 	string `json:"key_file"`
}