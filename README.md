# go-influxdb
Go integration for InfluxDB

Send Data:

$ curl -i -X POST localhost:8080/iot -H "Content-Type: application/json" -d '{"gpu_id": "3070", "rig_id": "10", "current_watts": 120}'

Get Data:

$ curl -i -X GET localhost:8080/iot
