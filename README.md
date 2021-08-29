# D2K - Docker to Kubernetes
A proxy that implements the Docker `v1.14` API ([defined here](https://docs.docker.com/engine/api/v1.41/)) allowing Kubernetes access from a Docker-like API.

## Milestones (in order)

| Compatibility with | Capabilities supported | Achieved at | Note
| --- | --- | --- | --- |
| `basic` | `run(detach)`, `entrypoint`, `command`, `cpus`, `memory` | `v0.0.1` | --- |
| `process status` | `ps`, `status` | --- | --- |
| `docker stop\|kill` | `stop`, `kill` | --- | --- |
| `docker logs` | `logs`, `stream logs (-f)` | --- | --- |
| `run interactive` | `shell (-it)` | --- | --- |
| `networking 1` | `port forwarding (-p 80:8080)`, `dynamic port forwarding (-P)` | --- | Might need to dynamically bind to local ports from `B2K` |
| `volumes` | `volumes` | --- | This might require some local integration with `inotify` and remote file copy similar to `kubectl cp` |

## Goal

Provide a transparent way to spawn Kubernetes PODs using the Docker API.