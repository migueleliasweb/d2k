# Using MITMProxy to debug API calls

After enabling your docker daemon to listen to port 2375, you can run:

```
docker run -it --rm --net host mitmproxy/mitmproxy:7.0.2 mitmproxy --mode reverse:127.0.0.1:2375
```

This will run a MITMProxy listening to power 8080 and proxying requests to port 2375. It lets you check the daemon responses easily.

It's super useful for debugging.

```
docker run -it --rm --net host mitmproxy/mitmproxy:7.0.2 mitmdump --mode reverse:127.0.0.1:2375
```