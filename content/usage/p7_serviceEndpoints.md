+++
title = "Service Control"
+++

## Headers
  * `Content-Type:Application/json`
  * `Accept:Application/json`
  * auth - see auth doco if auth is enabled

## URI Methods
* GET `/api/service/{service}` - _returns status of a systemd service_

      ```
      GET http://0.0.0.0:6661/api/service/web-api
      response:
      {
          "status": "active"
      }
      ```
---

* POST `/api/service/{service}/{action}` - _initiates a systemd service control_
  * {action} = start|stop|restart
  
      ```
      POST http://0.0.0.0:6661/api/service/redis/stop
      response:
      {
          "status": "success"
      }
      ```
---
