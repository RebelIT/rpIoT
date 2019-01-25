# rpIoT
raspberryPi IOT device API

## About this:
A configurable API for my raspberryPi IOT devices around the house. Toggle functions on and off in the 
configuration file before starting the service.  Used for integrating/controlling rPI's into your home automation. 

## Deploying it to a raspberryPi (ansible & go):
* You need a working local GOLang environment for ansible to build the project
* Update api_config.json 
    * true/false values for what endpoints you want enabled
    * http port number for the api to listen on
    * statsd host if you want statsd metrics to be emitted on api usage
* Update ansible_deploy.yml with your local gopath
   ```
    gopath: "/Users/rebelit/go"
    goroot: "/usr/local/opt/go/libexec"
   ```
* Run it (depending on your setup you may not need `--ask-sudo-pass`)
    ```
    ansible-playbook ansible_deploy.yml -i ansible_host --ask-sudo-pass
    ```
    
## Build/Test/Run it (docker):
Build the image then run docker with one of the scripts, test, build, or run.  no need for GOLang to be configured 
locally!

* Build the image:
    * `docker build -t rpiot .`
* Run all Tests:
    * `docker run -v $PWD:/go/src/github.com/rebelit/rpIoT -i -t --rm rpiottest '/test.sh'`
* Build the program:
    * `docker run -v $PWD:/go/src/github.com/rebelit/rpIoT -i -t --rm rpiottest '/build.sh'`
    * `scp main pi@yourHost:/your/directory/here/`
* Run the program:
    * `docker run -v $PWD:/go/src/github.com/rebelit/rpIoT -i -t --rm rpiottest '/run.sh'`
    

## Using it:
Hit the current default endpoints to test, add more for any other functions you need for your 
devices. 

TIP: If you do not have an API client, you can change these endpoints in the httpHandlers.go to all GET's
and hit them from a browser on your phone or tablet to perform the actions. 

Default endpoints:
* GET /api/alive
    * health check of the API is running
* POST /api/power/{reboot|shutdown}
    * power control the raspberryPi
* POST /api/apt/{update|upgrade}
    * system package updates and installation
* POST /api/service/{service name}/{start|stop|restart}
    * control local services remotely
* add more as you need, or contact me/submit issue to add more. 


#### What else can I do (outside of the default endpoints):
###### Project: 
DAK wall mounted digital calendar & digital magic mirror
* Turn the HDMI display on and off (pi 3B+) - not tested on other models
    * POST /api/display/{on|off}
    
###### Project: 
Hard wired electronics control with GPIO pins
* Control GPIO pins
    * POST /api/pullup/{gpio pin #}
        * returns the pin state readout in the json body `{"message":}`
    * POST /api/pulldown/{gpio pin #}
        * returns the pin state readout in the json body `{"message":}`
    * POST /api/toggle/{gpio pin #}

---
### Notes:
I set it up to check the config file before every API action. While it is less efficient to read the file every 
time, it adds the flexibility to toggle endpoints on and off without having to restart the entire API service.
 
 
## Test it:

```
docker build -t rpiot .
docker run -v $PWD:/go/src/github.com/rebelit/rpIoT -i -t rpiottest '/test.sh'
docker run -v $PWD:/go/src/github.com/rebelit/rpIoT -i -t rpiottest '/build.sh'
docker run -v $PWD:/go/src/github.com/rebelit/rpIoT -i -t rpiottest '/run.sh'

```