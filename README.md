# MQTT2Cribl

Little utility app for reading MQTT messages and sending them off to Cribl. Made this specifically for pulling MQTT messages from dsmr and Shelly devices.

## Variables
| Environment variable name | Description                                                | Example                                           |
| ------------------------- | :--------------------------------------------------------: | ------------------------------------------------: |
| MQTT2CRIBL_ENDPOINT       | endpoint for the Cribl host or cluster                     | http://cribl.service.consul:2300/cribl/_bulk      |
| NQTT2CRIBL_TOPICS         | topics to subscribe to, separated by a ';'                 | dsmr/json;shellies/+/+/power;shellies/+/+/+/power |
| MQTT2CRIBL_BROKER         | MQTT broker url                                            | tcp://overseer.homelab:1883                       |
| MQTT2CRIBL_CLIENTID       | client name the app will use when connecting to the broker | go_mqtt2cribl                                     |

## Runtime
Build the binary and use this in a Nomad job for easy scheduling. App keeps running and logs to stdout.
