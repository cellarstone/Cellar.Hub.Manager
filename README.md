
# Prerequisities

## OS

Ubuntu 16.04 LTS

### User 

Username : cellarstone

Password : Cllrs123IoT456

Add Sudo rights : `usermod -aG sudo Cellarstone`

Add Docker rights : `usermod -aG docker Cellarstone`

## Docker

Docker CE 

and enable Docker Swarm with command `docker swarm init`

## Google Cloud

Add into path `/home/cellarstone/Apps/GoogleCloudKeys/cellarhubmanager` file `GoogleCloud-cellarhubmanager.json`

### User way
Set environment variable `GOOGLE_APPLICATION_CREDENTIALS` by editing `.bashrc` user file. Use this command 

```Shell
gedit /home/cellarstone/.bashrc
```

and add this row at the end of file

`export GOOGLE_APPLICATION_CREDENTIALS="/home/cellarstone/Apps/GoogleCloudKeys/cellarhubmanager/GoogleCloud-cellarhubmanager.json"`


### Environment way


```Shell
gedit /etc/environment
```


and add this row at the end of file

`GOOGLE_APPLICATION_CREDENTIALS="/home/cellarstone/Apps/GoogleCloudKeys/cellarhubmanager/GoogleCloud-cellarhubmanager.json"`


## Ngrok link

All necessary will be deliver with app itself.

Connect via SSH `ssh cellarstone@0.tcp.ngrok.io -p36428`

## Equinox

All necessary will be deliver with app itself.


## Cellarhubmanager - Deamon service - systemd 

Always-running process by systemd
https://fabianlee.org/2017/05/21/golang-running-a-go-binary-as-a-systemd-service-on-ubuntu-16-04/

### Create config file
`sudo gedit /lib/systemd/system/cellarhubmanager.service`

```Shell
[Unit]
Description=Cellar.Hub.Manager service

[Service]
Type=simple
User=root
Group=root
WorkingDirectory=/home/cellarstone/Apps/Cellar.Hub.Manager
ExecStart=/home/cellarstone/Apps/Cellar.Hub.Manager/cellarhubmanager
Restart=always
Environment=GOOGLE_APPLICATION_CREDENTIALS=/home/cellarstone/Apps/GoogleCloudKeys/cellarhubmanager/GoogleCloud-cellarhubmanager.json

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







# Cellarstone Cloud

Hub is connected to the Cloud and sending status data each minute.


## Cloud Storage

check newer version of docker-stack.yml

## PubSub

PUBLISH      - (each hour) send info about device (cellarDeviceID, cellarHostName, cellarMACaddress,  Wifi, IP ... etc.)
PUBLISH      - (each minute) send status message


# Useful commands

## Cellarhubmanager logs

```Shell
journalctl -u cellarhubmanager
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