run:
	docker-compose up -d
	watcher #github.com/canthefason/go-watcher
test:
	bro #github.com/marioidival/bro
cover:
	echo "" > coverage.txt
	for d in $(shell go list ./... | grep -v vendor); do \
		go test -race -v -coverprofile=profile.out -covermode=atomic $$d || exit 1; \
		[ -f profile.out ] && cat profile.out >> coverage.txt && rm profile.out; \
	done