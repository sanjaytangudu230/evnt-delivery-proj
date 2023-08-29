This is a simple event delivery system implemented in Go. This servie maintain FIFO and retries on delivery failures.

Steps to Run the service:

 i)CLone the repository/ Download the zip file
 
 ii) Make sure docker is running
 
 iii) Open the project and go to terminal.
 
 iv) Run 'make deploy' command in termial.

This service has been tested under good throughput using Apache Benchmark

![WhatsApp Image 2023-08-28 at 11 37 13](https://github.com/sanjaytangudu230/evnt-delivery-proj/assets/61742536/43f4f7bd-6070-4df4-b43f-0e116d21bdf9)

Curl to Ingest Event:

curl --location 'http://localhost:8080/api/event-delivery/ingest-event' \
--header 'Content-Type: application/json' \
--data '{
    "user_id" : "2345",
    "payload": "dfghjk"
}'
