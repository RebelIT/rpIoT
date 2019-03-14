package config

type Config struct{
	InstallDir			string `json:"install_dir"`
	ApiPort				string `json:"api_port"`
	StatsdHost			string `json:"statsd_host"`
	SlackWebhook 		string `json:"slack_webhook"`
	DefaultUser			string `json:"default_user"`
	DefaultUserToken	string `json:"default_user_token"`
	AuthEnabled 		bool `json:"auth_enabled"`
	SslEnabled 			bool `json:"ssl_enabled"`
	SslCertName 		string `json:"ssl_cert_name"`
	SslKeyName			string `json:"ssl_key_name"`
	EndpointPower		bool `json:"endpoint_power"`
	EndpointPackages 	bool `json:"endpoint_packages"`
	EndpointService 	bool `json:"endpoint_service"`
	EndpointDisplay 	bool `json:"endpoint_display"`
	EndpointGpio 		bool `json:"endpoint_gpio"`
}