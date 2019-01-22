# rpIoT
raspberryPi IOT device API

## About this
A configurable API for my raspberryPi IOT devices around the house. Toggle functions on and off in the 
configuration file before starting the service.  Used for integrating/controlling rPI's into your home automation. 

## Deploying it (ansible)
```
cd ansible
ansible-playbook ansible_deploy.yml -i 10.0.0.115s
```