
# Prerequisities

## OS

Ubuntu 16.04

### User 

Username : Cellarstone

Password : Cllrs123IoT456

Add Sudo rights : `usermod -aG sudo Cellarstone`

## Docker

Docker CE

## Ngrok link

Install ngrok
```Shell
sudo ./ngrok service install --config=./ngrok.yml
```

Config for ngrok
```Yaml
authtoken: 6mRpb1ZoPHJ6ro1KfqAPq_4SUZciW7QnuJSNo6U9Tiy
tunnels:
  ssh:
    proto: tcp
    addr: 22
  dashboard:
    proto: http
    addr: 10001
```

Start service
```Shell
sudo ./ngrok service start
```


Stop service
```Shell
sudo ./ngrok service stop
```


Restart service
```Shell
sudo ./ngrok service restart
```


Uninstall service
```Shell
sudo ./ngrok service uninstall
```

### Status

http://localhost:4040/status


List of service
```Shell
ps -ef | grep ngrok
```

```Shell
sudo systemctl -a
```

```Shell
service ngrok status
```



## Deamon service - systemd

Always-running process by systemd
https://fabianlee.org/2017/05/21/golang-running-a-go-binary-as-a-systemd-service-on-ubuntu-16-04/

### Create config file
`sudo gedit /lib/systemd/system/cellarhubmanager.service`

```Shell
[Unit]
Description=Cellar.Hub.Manager service
ConditionPathExists=/home/cellarstone/Apps/Cellar.Hub.Manager
After=network.target

[Service]
Type=simple
User=root
Group=root
LimitNOFILE=1024

Restart=on-failure
RestartSec=10
startLimitIntervalSec=60

WorkingDirectory=/home/cellarstone/Apps/Cellar.Hub.Manager
ExecStart=/home/cellarstone/Apps/Cellar.Hub.Manager/cellarhubmanager

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

### Start and monitor service

Enable service (connect config file with systemd)

```Shell
sudo systemctl enable cellarhubmanager.service
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




### Systemclt way commands

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

`ngrok authtoken 6mRpb1ZoPHJ6ro1KfqAPq_4SUZciW7QnuJSNo6U9Tiy`


REST API token

test1
`RkdUfZiKNoaWLTdnZF4w_2r1S1nBNMXyAERATUcnGS`

test2
`5SdaauXmgd6tNhLyP4KL9_4nJr8wfGh9XPLSRPXnm4m`


# Dropbox 

for unhandled situation



# Connection to the Cellarstone Cloud

Hub is connected to the Cloud and sending status data each minute.

