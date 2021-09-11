# Configuring docker to use TCP socket

1) Configure the `daemon.json`:

```
# /etc/docker/daemon.json
{"hosts": ["tcp://127.0.0.1:2375", "unix:///var/run/docker.sock"]}
```

2) If using `systemd`, edit the service file and remove the `-H` flag from the `ExecStart`

3) If using `systemd` reload the service with

```
systemctl reload docker.service
systemctl restart docker.service
```

