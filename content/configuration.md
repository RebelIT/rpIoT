+++
title = "Configuring the API"
+++

## Application Configs
service starts up with a set of default configs set to high security.  If you wish to override
any of these values you must provide a `-config=/path/to/api_config.json` file with override values.

* **Default hard coded values**
  ```
  "install_dir": "/etc/api/"
  "api_port": "8080",
  "statsd_host": "",
  "slack_webhook": "",
  "default_user": "admin",
  "default_user_token": "myPiApi",
  "auth_enabled": true,
  "ssl_enabled": true,
  "ssl_cert_name": "cert.cert",
  "ssl_key_name": "key.key",
  "endpoint_power": true,
  "endpoint_packages": true,
  "endpoint_service": true,
  "endpoint_display": true,
  "endpoint_gpio": true
  ```

* **Options breakdown**
  * `install_dir` - _application installation directory_
  * `api_port` - _application http port_
  * `statsd_host` - _udp listener for statsd metrics_
  * `slack_webhook` - _slack webhook uri for alerting_
  * `default_user` - _user for authentication_
  * `default_user_token` - _user password for authentication_
  * `auth_enabled` - _enable or disable authentication. if false default\_user\* values are ignored_
  * `ssl_cert_name` - _certificate file for SSL_
  * `ssl_key_name` - _private key file for SSL_
  * `ssl_enabled"` - _enable or disable SSL. if false ssl\_cert\* values are ignored_
  * `endpoint_power` - _enables or disables the endpoint to control rPi power_
  * `endpoint_packages` - _enables or disables the endpoint to install/upgrade apt packages_
  * `endpoint_service` - _enables or disables the endpoint to control installed services state_
  * `endpoint_display` - _enables or disables the endpoint to control rPi hdmi port_
  * `endpoint_gpio` - _enables or disables the endpoint to control rPi gpio pins_

* Any of the above default values can be overridden by using the api_config.json. **example** [api_config.json](https://github.com/RebelIT/rpIoT/blob/master/examples/example_api_config.json)
* I am not providing instructions on generating a SSL certificate. You can self sign or use a public CA signed cert as long as you have the key.
