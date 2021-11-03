# make influx1_up
influx1_up:
	docker compose -f ./influxdb/influxdb-1.yml up --build

# make influx2_up
influx2_up:
	docker compose -f ./influxdb/influxdb-2.yml up --build

# make relay1_up
relay1_up:
	influxdb-relay -config ./relay/relay-1.toml

# make relay2_up
relay2_up:
	influxdb-relay -config ./relay/relay-2.toml

# make relay_up
relay_up:
	docker compose -f ./relay/relay.yml up

# make loadbalancer_up
loadbalancer_up:
	docker compose -f ./loadbalancer/nginx.yml up --build
