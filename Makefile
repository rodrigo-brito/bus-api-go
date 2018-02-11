run:
	cd docker
	mkdir -p ./docker/volumes/mysql
	docker-compose -f docker/docker-compose-dev.yaml -p bus-api up --force-recreate
	docker logs -f --tail=50 busapi_api_1
cover:
	echo "" > coverage.txt
	for d in $(shell go list ./... | grep -v vendor); do \
		go test -race -v -coverprofile=profile.out -covermode=atomic $$d || exit 1; \
		[ -f profile.out ] && cat profile.out >> coverage.txt && rm profile.out; \
	done