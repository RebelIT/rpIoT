+++
title = "Using the Endpoints"
+++

## Headers
  * `Content-Type:Application/json`
  * `Accept:Application/json`
  * auth - see auth doco if auth is enabled

## API endpoints
* **Base endpoints**
  * GET `/api/alive` - _return a json status if the API is alive_
  * GET `/api/system` - _return a json status if the running system stats_
* **Power Control endpoints**
  * POST `/api/power/{acton}` - _return a json status if the API is alive_

        ```
        actions:
        reboot - Reboots the Raspberry Pi
        shutdown - performs a graceful system shutdown
        ```

* **Package manager endpoints**
  * GET `/api/apt` - _return a json response of the var/log/apt/history.log_
  * POST `/api/apt/{action}` - _return a json response of the var/log/apt/history.log_

        ```
        actions:
        update - performs a local apt-get update
        upgrade - performs a local apt-get upgrade
        ```

  * POST `/api/apt/{package}/{action}` - _initiates apt-get to install a package_

        ```
        package:
        any package name in your apt repos, reddis, vim, apache, etc...
        actions:
        install - performs an apt-get install {package}
        remove - performs an apt-get remove {package}
        ```

* **Service Control endpoints**
  * GET `/api/service/{service}` - _returns the status of a service (active/inactive)_
 
        ```
        service:
        any installed systemd service name
        ``` 
        
  * POST `/api/service/{service}/{action}` - _return a json response of the var/log/apt/history.log_

        ```
        service:
        any installed systemd service name
        actions:
        start - performs a systemctl start {service}
        stop - performs a systemctl stop {service}
        restart - performs a systemctl restart {service}
        ```

* **HDMI Control endpoints**
  * GET `/api/display` - _return a status of the HDMI display power 0-off | 1-on_
  * POST `/api/display/{action}` - _turns the HDMI power on or off_

        ```
        actions:
        on - turns on hdmi power
        off - turns off hdmi power
        ```

* **GPOI Pin Control endpoints**
  * POST `/api/gpio/{number}/pullup` - _performs a pin# pullup_
  * POST `/api/gpio/{number}/pulldown` - _performs a pin# pulldown_
  * POST `/api/gpio/{number}/toggle` - _performs a pin# toggle_
  * POST `/api/gpio/{number}/depress/{milliseconds}` - _simulates a button press (toggle x2 for n milliseconds)_
     
        ```
        numbers - any valid GPIO pin # (2-25)
        milliseconds - any measurement of time in ms. 
             (200-300 milliseconds usually suffice)
        
        ```

## HTTP Return Body
* Http codes used: `OK | BadRequest | Forbidden | InternalServerError | Unauthorized`
* All endpoints return a json response on 200 - OK success
  * example display power on return:

        ```
        {
            "namespace": "/api/display/on",
            "message": "display_power=1",
            "data": null
        }
        ```

  * example get system stats return:

        ```
        {
            "host": {
                "uptime": "151296",
                "kernel_version": "4.14.98-v7+",
                "platform": "raspbian",
                "platform_ver": "9.8",
                "boottime": "1552535567",
                "hostname": "piCalendar",
                "host_id": "c850aa6c-693f-4f48-9a3e-1cc068f7b2e0"
            },
            "cpu": {
                "cpu_stats": [
                    {
                        "cpu_num": 0,
                        "mhz": 0,
                        "model_name": "ARMv7 Processor rev 4 (v7l)",
                        "cores": 1,
                        "util_percent": 1200
                    },
                    {
                        "cpu_num": 1,
                        "mhz": 0,
                        "model_name": "ARMv7 Processor rev 4 (v7l)",
                        "cores": 1,
                        "util_percent": 1200
                    },
                    {
                        "cpu_num": 2,
                        "mhz": 0,
                        "model_name": "ARMv7 Processor rev 4 (v7l)",
                        "cores": 1,
                        "util_percent": 1200
                    },
                    {
                        "cpu_num": 3,
                        "mhz": 0,
                        "model_name": "ARMv7 Processor rev 4 (v7l)",
                        "cores": 1,
                        "util_percent": 1200
                    }
                ]
            },
            "net": {
                "total": "968077312",
                "free": "197435392"
            },
            "disk": {
                "total": 13776019456,
                "used": 4469084160,
                "free": 8583548928
            }
        }
      ```
