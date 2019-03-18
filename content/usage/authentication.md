+++
title = "Authentication and Security"
+++

## Authentication enabled usage
If `"auth_enabled": true` you must provide auth headers in each API call
* Required headers for all API calls:
  ```
  X-API-User:<default_user config value>
  X-API-Token:<default_user_token config value>
  ```

* Disabling auth `"auth_enabled": false` makes your API wide open to anyone on your
network.

## SSL enabled usage
If `"ssl_enabled": true` you must place your cert and key in the `install_dir` directory.  You can use any CA signed cert or a self signed one, as long as you have the key.
