
# Prerequisities

## OS

Ubuntu 16.04

## Docker

Docker CE

## Deamon service - systemd

Always-running process by systemd
https://fabianlee.org/2017/05/21/golang-running-a-go-binary-as-a-systemd-service-on-ubuntu-16-04/

`sudo gedit /lib/systemd/system/cellarhubmanager.service`

```Shell
[Unit]
Description=Cellar.Hub.Manager service
ConditionPathExists=/home/Apps/Cellar.Hub.Manager
After=network.target

[Service]
Type=simple
User=root
Group=root
LimitNOFILE=1024

Restart=on-failure
RestartSec=10
startLimitIntervalSec=60

WorkingDirectory=/home/Apps/Cellar.Hub.Manager
ExecStart=/home/Apps/Cellar.Hub.Manager/cellarhubmanager

# make sure log directory exists and owned by syslog
PermissionsStartOnly=true
ExecStartPre=/bin/mkdir -p /var/log/echoservice
ExecStartPre=/bin/chown syslog:adm /var/log/echoservice
ExecStartPre=/bin/chmod 755 /var/log/echoservice
StandardOutput=syslog
StandardError=syslog
SyslogIdentifier=cellarhubmanager

[Install]
WantedBy=multi-user.target
```

### Systemclt way commands

Enable a service as daemon service

```Shell
sudo systemctl enable cellarhubmanager.service
```

Service start

```Shell
sudo systemctl start cellarhubmanager
```



Restart a service

```Shell
sudo systemctl restart cellarhubmanager
```


List of all service by systemd

```Shell
sudo systemctl -a
```



Tail the log

```Shell
sudo journalctl -f -u cellarhubmanager
```



### Service way commands


Enable a service as daemon service - ??? work ???

```Shell
service cellarhubmanager enable   
```

Service status - GREAT !!!

```Shell
service cellarhubmanager status
```

Service start

```Shell
service cellarhubmanager start
```

Service restart

```Shell
service cellarhubmanager restart
```


### Other commands

List of the service process

```Shell
ps -ef | grep cellarhubmanager
```


# Libraries

Equinox-io
 - automated creation of cross platform binaries
 - automated downloading new binaries from cloud storage
 - same author as ngrok
 - https://github.com/equinox-io

Facebook Grace
 - Graceful restart for Golang webserver
 - https://github.com/facebookgo/grace



# Unique Id

Hub ID - UNIQUE ID for whole world
 - it is tied up with user account


# Self-updating program

Equinox.io 

```Shell
./equinox release \
  --version="0.3.4" \
  --platforms="darwin_amd64 linux_amd64" \
  --signing-key=equinox.key \
  --app="app_h9SyPnPqLpq" \
  --token="fHeN81JECeiVAxoiJfEyPxBGSdMnBxVjsxZffG7wrHgEvwqJshuF" \
  ../
```

# Graceful restart - without blackout

Facebook Grace


# Expose device to ngrok

Install ngrok token

`ngrok authtoken ASDFASDFERWEFASFAWEFWAFA`

ngrok 


# Dropbox 

for unhandled situation



# Connection to the Cellarstone Cloud

Hub is connected to the Cloud and sending status data each minute.

