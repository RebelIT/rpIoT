+++
title = "System Information"
+++

## Headers
  * `Content-Type:Application/json`
  * `Accept:Application/json`
  * auth - see auth doco if auth is enabled

## URI Methods
* GET `/api/system` - _returns system information_

    ```
    GET http://0.0.0.0:6661/api/system
    response:
    {
        "host": {
            "uptime": "5316",
            "kernel_version": "4.14.79-v7+",
            "platform": "raspbian",
            "platform_ver": "9.6",
            "boottime": "1554780785",
            "hostname": "piDev",
            "host_id": "a336c86e-879f-442d-bf62-00d9f8102c79"
        },
        "cpu": {
            "cpu_stats": [
                {
                    "cpu_num": 0,
                    "mhz": 1500,
                    "model_name": "ARMv7 Processor rev 4 (v7l)",
                    "cores": 1
                },
                {
                    "cpu_num": 1,
                    "mhz": 1500,
                    "model_name": "ARMv7 Processor rev 4 (v7l)",
                    "cores": 1
                },
                {
                    "cpu_num": 2,
                    "mhz": 1500,
                    "model_name": "ARMv7 Processor rev 4 (v7l)",
                    "cores": 1
                },
                {
                    "cpu_num": 3,
                    "mhz": 1500,
                    "model_name": "ARMv7 Processor rev 4 (v7l)",
                    "cores": 1
                }
            ]
        },
        "mem": {
            "total": "984707072",
            "free": "804564992"
        },
        "disk": {
            "total": 15606296576,
            "used": 2544054272,
            "free": 12395008000
        }
      }
      ```
---
