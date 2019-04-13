+++
title = "Power Control"
+++

## Headers
  * `Content-Type:Application/json`
  * `Accept:Application/json`
  * auth - see auth doco if auth is enabled

## URI Methods
* POST `/api/power/{acton}` - _initiates a system power control_
  * {action} = reboot|shutdown

      ```
      POST http://0.0.0.0:6661/api/power/reboot
      response:
      {
        "status": "async reboot in progress..."
      }
      ```
---
