# Message Service

The message service is a simple a system for exchanging text based messages.

## Getting started

1. Install Consul

	Consul is the default registry/discovery for go-micro apps. It's however pluggable.
	[https://www.consul.io/intro/getting-started/install.html](https://www.consul.io/intro/getting-started/install.html)

2. Run Consul
	```
	$ consul agent -server -bootstrap-expect 1 -data-dir /tmp/consul
	```

3. Run Sync Service. If you use consul for discovery then you can use the same service for synchronization. 

4. Download and start the service

	```shell
	go get github.com/micro/message-srv
	message-srv
	```

	OR as a docker container

	```shell
	docker run microhq/message-srv --registry_address=YOUR_REGISTRY_ADDRESS --sync_address=YOUR_SYNC_ADDRESS
	```

