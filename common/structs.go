package common

//Config File functions toggle on/off
type Function struct{
	Name		string `json:"name"`
	Enabled 	bool `json:"enabled"`
}

type Config struct{
	ApiPort				string `json:"api_port"`
	Statsd_host			string `json:"statsd_host"`
	SlackWebhook 		string `json:"slack_webhook"`
	DefaultUser			string `json:"default_user"`
	DefaultUserToken	string `json:"default_user_token"`
	Ssl 				Ssl `json:"ssl"`
	Functions			[]Function
}

type Ssl struct{
	Enabled 	bool   `json:"enabled"`
	CertFile 	string `json:"cert_file"`
	KeyFile 	string `json:"key_file"`
}

type SlackMsg struct{
	Text		string `json:"text"`
	Username 	string `json:"username"`
	IconPath	string `json:"icon_path"`
}