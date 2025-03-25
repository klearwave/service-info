IMAGE ?= ghcr.io/klearwave/service-info
IMAGE_VERSION ?= unstable
image:
	@docker build . -t $(IMAGE):$(IMAGE_VERSION)

#
# testing
#
up:
	docker compose up

down:
	docker compose down

# e2e test with embedded db/http server
test-e2e-embedded:
	export E2E_EMBEDDED=true && \
		go test ./test -run ^TestE2E$

# e2e test against active db/http server
test-e2e:
	go test ./test -run ^TestE2E$
