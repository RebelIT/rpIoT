+++
title = "Display Control"
+++

## Headers
  * `Content-Type:Application/json`
  * `Accept:Application/json`
  * auth - see auth doco if auth is enabled

## URI Methods
* GET `/api/display` - _returns power state of the HDMI out_
  * hdmi_state = 0 Off|1 On

      ```
      GET http://0.0.0.0:6661/api/display
      response:
      {
          "hdmi_state": 1
      }
      ```
---

* POST `/api/display/{action}` - _initiates a HDMI power state control_
  * {action} = on|off

      ```
      POST http://0.0.0.0:6661/api/display/off
      response:
      {
          "hdmi_state": 0
      }
      ```
---
