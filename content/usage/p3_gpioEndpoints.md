+++
title = "GPIO Pin Control"
+++

## Headers
  * `Content-Type:Application/json`
  * `Accept:Application/json`
  * auth - see auth doco if auth is enabled

## GPIO Guide

```
Rev 2 and 3 Raspberry Pi                        Rev 1 Raspberry Pi (legacy)
+-----+---------+----------+---------+-----+      +-----+--------+----------+--------+-----+
| BCM |   Name  | Physical | Name    | BCM |      | BCM | Name   | Physical | Name   | BCM |
+-----+---------+----++----+---------+-----+      +-----+--------+----++----+--------+-----+
|     |    3.3v |  1 || 2  | 5v      |     |      |     | 3.3v   |  1 ||  2 | 5v     |     |
|   2 |   SDA 1 |  3 || 4  | 5v      |     |      |   0 | SDA    |  3 ||  4 | 5v     |     |
|   3 |   SCL 1 |  5 || 6  | 0v      |     |      |   1 | SCL    |  5 ||  6 | 0v     |     |
|   4 | GPIO  7 |  7 || 8  | TxD     | 14  |      |   4 | GPIO 7 |  7 ||  8 | TxD    |  14 |
|     |      0v |  9 || 10 | RxD     | 15  |      |     | 0v     |  9 || 10 | RxD    |  15 |
|  17 | GPIO  0 | 11 || 12 | GPIO  1 | 18  |      |  17 | GPIO 0 | 11 || 12 | GPIO 1 |  18 |
|  27 | GPIO  2 | 13 || 14 | 0v      |     |      |  21 | GPIO 2 | 13 || 14 | 0v     |     |
|  22 | GPIO  3 | 15 || 16 | GPIO  4 | 23  |      |  22 | GPIO 3 | 15 || 16 | GPIO 4 |  23 |
|     |    3.3v | 17 || 18 | GPIO  5 | 24  |      |     | 3.3v   | 17 || 18 | GPIO 5 |  24 |
|  10 |    MOSI | 19 || 20 | 0v      |     |      |  10 | MOSI   | 19 || 20 | 0v     |     |
|   9 |    MISO | 21 || 22 | GPIO  6 | 25  |      |   9 | MISO   | 21 || 22 | GPIO 6 |  25 |
|  11 |    SCLK | 23 || 24 | CE0     | 8   |      |  11 | SCLK   | 23 || 24 | CE0    |   8 |
|     |      0v | 25 || 26 | CE1     | 7   |      |     | 0v     | 25 || 26 | CE1    |   7 |
|   0 |   SDA 0 | 27 || 28 | SCL 0   | 1   |      +-----+--------+----++----+--------+-----+
|   5 | GPIO 21 | 29 || 30 | 0v      |     |
|   6 | GPIO 22 | 31 || 32 | GPIO 26 | 12  |
|  13 | GPIO 23 | 33 || 34 | 0v      |     |
|  19 | GPIO 24 | 35 || 36 | GPIO 27 | 16  |
|  26 | GPIO 25 | 37 || 38 | GPIO 28 | 20  |
|     |      0v | 39 || 40 | GPIO 29 | 21  |
+-----+---------+----++----+---------+-----+
```


## URI Methods
* GET `/api/gpio` - _returns the state of all GPIO pins_

      ```
      GET http://0.0.0.0:6661/api/gpio
      response:
      {
          "pins": [
              {
                  "bcm_pin": 2,
                  "state": 1
              },
              {
                  "bcm_pin": 3,
                  "state": 1
              },
              {
                  "bcm_pin": 4,
                  "state": 1
              },
              ... 5-25 (same) ...
              {
                  "bcm_pin": 26,
                  "state": 0
              }
          ]
      }
      ```
---

* GET `/api/gpio/{pin_number}` - _returns the state of a single GPIO pin_
  * pin_number = 2-26 available gpio pins

      ```
      GET http://0.0.0.0:6661/api/gpio/12
      response:
      {
          "bcm_pin": 12,
          "state": 0
      }
      ```
---

* POST `/api/gpio/{pin_number}/pullup` - _performs a GPIO pin pullup_
  * pin_number = 2-26 available gpio pins

      ```
      POST http://0.0.0.0:6661/api/gpio/12/pullup
      response:
      {
          "bcm_pin": 12,
          "state": 1
      }
      ```
---

* POST `/api/gpio/{pin_number}/pulldown` - _performs a GPIO pin pulldown_
  * pin_number = 2-26 available gpio pins

      ```
      POST http://0.0.0.0:6661/api/gpio/12/pulldown
      response:
      {
          "bcm_pin": 12,
          "state": 0
      }
      ```
---

* POST `/api/gpio/{pin_number}/toggle` - _performs a GPIO pin toggle (change state)_
  * pin_number = 2-26 available gpio pins

      ```
      POST http://0.0.0.0:6661/api/gpio/13/toggle
      response:
      {
          "bcm_pin": 13,
          "state": 1
      }
      ```
---

* POST `/api/gpio/{pin_number}/depress/{millisecond}` - _simulates a button depress connected to GPIO pins_
  * pin_number = 2-26 available gpio pins
  * millisecond = any unit of time in ms (200-300 is about right for a button push simulation)

      ```
      POST http://0.0.0.0:6661/api/gpio/13/depress/200
      response:
      {
          "bcm_pin": 13,
          "state": 0
      }
      ```
---
