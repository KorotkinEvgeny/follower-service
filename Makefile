load-install:
	go install github.com/tsenart/vegeta@latest

load-lazy: load-install
	jq -ncM 'while(true; .+1) | {method: "POST", url: "http://localhost:3000/api/v1/events", body: {timestamp: now, value: range(100), id:.|tostring} | @base64}'  | \
    vegeta attack -lazy --format=json -duration=60s | tee results.bin | vegeta report

load-high-rate: load-install
	jq -ncM 'while(true; .+1) | {method: "POST", url: "http://localhost:3000/api/v1/events", body: {timestamp: now, value: range(100), id:.|tostring} | @base64}'  | \
    vegeta attack --format=json -duration=60s -rate=50 | tee results.bin | vegeta report

load-retrieve:
	jq -ncM '{method: "GET", url: "http://localhost:3000/api/v1/events/average"}'  | \
    vegeta attack --format=json -duration=60s -rate=50 | tee results.bin | vegeta report

read-load-report:
	cat results.bin | vegeta report -type='hist[0,2ms,4ms,6ms]'

test:
	go test -v ./...

test-race:
	go test -race ./...