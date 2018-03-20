# Experiments in UDP Multicasting

This package implements a UDP multicast node discovery mechanism.

In TCP/IP there is a similar mechanism provided for in the UDP broadcast and multicast. This package wraps Golang's UDP functions in the `net` package.

## Background

The following is rephrased from [here](https://www.quora.com/What-is-the-difference-between-broadcasting-and-multicasting).

**Broadcast** sends packets to all devices on a LAN. Unlike multicasting, there is no mechanism to limit the recipients of a broadcast: all packets go to all devices whether they want them or not. There is no mechanism to forward broadcasts between LANs.

**Multicast** sends packets to all devices in a specified group. Membership in a group is set up when devices send "join" packets to an upstream router, and routers and switches keep track of this membership. When multicast packets arrive at a switch, they are only sent to devices or segments (such as WiFi) where at least one device wants them. Multicast can traverse the networks where it has been configured.

## Examples

This package comes with some small command line utilities in the `examples` dir.

In a terminal window, run the following from the root of this repository.

```bash
$ docker run -it --rm -v ${PWD}:/go/src/github.com/andrefsp/go-multicast golang:1.9 go run /go/src/github.com/andrefsp/go-multicast/example/main.go
#   Current nodes in the cluster:::
#            &{Hostname:e54d414943c0 IP:172.17.0.2 LastHeartbeat:2018-03-20 16:47:17.996843836 +0000 UTC m=+4.004034182}
#   Current nodes in the cluster:::
#            &{Hostname:e54d414943c0 IP:172.17.0.2 LastHeartbeat:2018-03-20 16:47:21.994131792 +0000 UTC m=+8.001321338}
#   2018/03/20 16:47:24 'b076198c77e1' joined the cluster
#   Current nodes in the cluster:::
#            &{Hostname:e54d414943c0 IP:172.17.0.2 LastHeartbeat:2018-03-20 16:47:27.997419646 +0000 UTC m=+14.004611191}
#            &{Hostname:b076198c77e1 IP:172.17.0.3 LastHeartbeat:2018-03-20 16:47:28.178230595 +0000 UTC m=+14.185422940}
#   Current nodes in the cluster:::
#            &{Hostname:e54d414943c0 IP:172.17.0.2 LastHeartbeat:2018-03-20 16:47:31.996477616 +0000 UTC m=+18.003722335}
#            &{Hostname:b076198c77e1 IP:172.17.0.3 LastHeartbeat:2018-03-20 16:47:32.174284968 +0000 UTC m=+18.181481111}
#   ^Csignal: interrupt
```

In a separate terminal window, run the following, also from the root of this repository.

```bash
$ docker run -it --rm -v ${PWD}:/go/src/github.com/andrefsp/go-multicast golang:1.9 go run /go/src/github.com/andrefsp/go-multicast/example/main.go
#   2018/03/20 16:47:23 'e54d414943c0' joined the cluster
#   2018/03/20 16:47:24 'b076198c77e1' joined the cluster
#
#   Current nodes in the cluster:::
#            &{Hostname:e54d414943c0 IP:172.17.0.2 LastHeartbeat:2018-03-20 16:47:25.996776668 +0000 UTC m=+3.824877605}
#            &{Hostname:b076198c77e1 IP:172.17.0.3 LastHeartbeat:2018-03-20 16:47:26.179862879 +0000 UTC m=+4.007876560}
#   Current nodes in the cluster:::
#            &{Hostname:e54d414943c0 IP:172.17.0.2 LastHeartbeat:2018-03-20 16:47:31.996546082 +0000 UTC m=+9.824557664}
#            &{Hostname:b076198c77e1 IP:172.17.0.3 LastHeartbeat:2018-03-20 16:47:30.174359331 +0000 UTC m=+8.002370913}
#   ^Csignal: interrupt
#
```

Both nodes know the existence of one another.
