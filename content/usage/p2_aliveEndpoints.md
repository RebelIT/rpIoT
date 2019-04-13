+++
title = "API Alive"
+++

## Headers
  * `Content-Type:Application/json`
  * `Accept:Application/json`
  * auth - see auth doco if auth is enabled

## URI Methods
* GET `/api/alive` - _http status health check endpoint_

      ```
      GET http://0.0.0.0:6661/api/alive
      response:
      {
        "status": "im running fine..."
      }
      ```
---
