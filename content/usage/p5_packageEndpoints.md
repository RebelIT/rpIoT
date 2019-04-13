+++
title = "Package Manager"
+++

## Headers
  * `Content-Type:Application/json`
  * `Accept:Application/json`
  * auth - see auth doco if auth is enabled

## URI Methods
* GET `/api/apt` - _returns package install history_

      ```
      GET http://0.0.0.0:6661/api/apt
      response:
      {
          "log": [
              "",
              "Start-Date: 2019-04-05  23:00:42",
              "Commandline: apt-get install jq",
              "Requested-By: pi (1000)",
              "Install: jq:armhf (1.5+dfsg-1.3), libjq1:armhf (1.5+dfsg-1.3, automatic), libonig4:armhf (6.1.3-2, automatic)",
              "End-Date: 2019-04-05  23:00:51",
              ""
          ]
      }
      ```
---

* POST `/api/apt/{action}` - _performs a system apt action_
  * {action} = update|upgrade

    ```
    POST http://0.0.0.0:6661/api/apt/update
    response:
    {
      "status": "update async action in progress, GET - /api/apt for status"
    }
    ```
---

* POST `/api/apt/{package}/{action}` - _initiates apt-get to install a package_
  * {package} = any package in APT Package Manager Repositories

    ```
    POST http://0.0.0.0:6661/api/apt/jq/install
    Response:
    {
        "status": "install jq async action in progress, GET - /api/apt for status"
    }
    ```
