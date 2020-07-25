# KEPSearch

This repo is just sample for accessing Neo4J through Go.

KEP(https://github.com/kubernetes/enhancements#enhancements-tracking-spreadsheet) can be retrieved via this tool.

# Usage

Launch Neo4j locally and then they mount KEP.csv as docker volume.
It is better to retrieve KEP from Google Spreadsheet directry.

```
cd docker_neo4j
docker-compose up -d
```

Get KEPs
```
go run main.go

0 :      Pod Security Policy
1 :      Addon management layer
2 :      CronJobs (previously ScheduledJobs)
3 :      Add AppArmor support
4 :      Add sysctl support
5 :      PodDisruptionBudget and /eviction subresource
6 :      Arbitrary/Custom Metrics in the Horizontal Pod Autoscaler
7 :      Support node-level user namespace remapping
8 :      Integrate Cluster Bootstrap/Discovery with Kubenetes Core
9 :      Seccomp
10 :     Taint Based Eviction
11 :     Snapshot / Restore Volume Support for Kubernetes (CRD + External Controller)
12 :     Provide RunAsGroup feature for Containers in a Pod
13 :     Kubelet Client TLS Certificate Rotation
14 :     Kubelet Server TLS Certificate Rotation
15 :     Ephemeral Containers
16 :     Limit node access to API
17 :     Dynamic Kubelet Configuration
18 :     Add support for resizing PVs
19 :     Raw block device using persistent volume source
20 :     Ability to create dynamic HA clusters with kubeadm
21 :     Local Ephemeral Storage Capacity IsolationÂ
22 :     Supports Storage as Allocatable Resource
23 :     Support paged LIST queries from the Kubernetes API
24 :     Taint node by Condition
25 :     Redesign Event API
26 :     Support AWS Network Load Balancer
27 :     Kubectl Diff
28 :     Configurable Pod Process Namespace Sharing
29 :     IPv6 support added
30 :     Add support for online resizing of PVs
31 :     Topology aware routing of services
32 :     External client-go credential providers
33 :     TokenRequest API and Kubelet integration
34 :     Enable kubectl to expose `ephemeral-storage` resource information
35 :     Schedule DaemonSet Pods by kube-scheduler
36 :     Node Failure Recover with ForceEvictione Taint
37 :     Deprecate ComponentStatus
38 :     Dynamic Maximum volume count
39 :     Server-side Apply
~
```

# Verification

```
❯ docker-compose logs 
Attaching to testneo4j
testneo4j | Directories in use:
testneo4j |   home:         /var/lib/neo4j
testneo4j |   config:       /var/lib/neo4j/conf
testneo4j |   logs:         /logs
testneo4j |   plugins:      /plugins
testneo4j |   import:       NOT SET
testneo4j |   data:         /var/lib/neo4j/data
testneo4j |   certificates: /var/lib/neo4j/certificates
testneo4j |   run:          /var/lib/neo4j/run
testneo4j | Starting Neo4j.
testneo4j | 2020-07-25 08:51:51.231+0000 INFO  Starting...
testneo4j | 2020-07-25 08:51:54.882+0000 INFO  ======== Neo4j 4.1.1 ========
testneo4j | 2020-07-25 08:51:57.508+0000 INFO  Performing postInitialization step for component 'security-users' with version 2 and status CURRENT
testneo4j | 2020-07-25 08:51:57.511+0000 INFO  Updating the initial password in component 'security-users'  
testneo4j | 2020-07-25 08:51:58.270+0000 INFO  Bolt enabled on 0.0.0.0:7687.
testneo4j | 2020-07-25 08:52:00.031+0000 INFO  Remote interface available at http://localhost:7474/
testneo4j | 2020-07-25 08:52:00.034+0000 INFO  Started.
```

# Configuration

The username and password of Neo4j are hard-coded in cmd/neo4jwrapper/neo4jwrapper.go.
- username: neo4j
- password: neo4j
