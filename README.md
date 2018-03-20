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
$ go run example/main.go
# Listening on 239.0.0.0:9999
# '301223baa544' joined the cluster
# map[301223baa544:0xc420056140]
# map[301223baa544:0xc420056180]
# map[301223baa544:0xc42009a080]
# map[301223baa544:0xc42009a0c0]
# map[301223baa544:0xc42009a100]
# '56785de23293' joined the cluster
# map[301223baa544:0xc42009a100 56785de23293:0xc42009a140]
# map[301223baa544:0xc42009a180 56785de23293:0xc42009a140]
# map[301223baa544:0xc42009a180 56785de23293:0xc42009a1c0]
# map[301223baa544:0xc420056200 56785de23293:0xc42009a1c0]
# map[301223baa544:0xc420056200 56785de23293:0xc42009a200]
# map[301223baa544:0xc42009a240 56785de23293:0xc42009a200]
# 2018/03/20 15:13:05 	 	 	56785de23293 left the cluster
# map[301223baa544:0xc42009a340]
```

In a separate terminal window, run the following, also from the root of this repository.

```bash
$ go run example/main.go
# Listening on 239.0.0.0:9999
# '56785de23293' joined the cluster
# map[56785de23293:0xc420092100]
# '301223baa544' joined the cluster
# map[301223baa544:0xc4200560c0 56785de23293:0xc420092100]
# map[56785de23293:0xc420092140 301223baa544:0xc4200560c0]

```

Both nodes know the existence of one another.
