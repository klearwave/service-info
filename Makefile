IMAGE ?= ghcr.io/klearwave/service-info
IMAGE_VERSION ?= unstable
image:
	@docker build . -t $(IMAGE):$(IMAGE_VERSION)

# NOTE: ensure the below DB_CONTAINER aligns with that of the docker-compose.yaml file
DB_CONTAINER ?= postgres:16.6-bullseye
db-container:
	docker run -d \
		--name postgres-container \
		-e POSTGRES_USER=postgres \
		-e POSTGRES_PASSWORD=postgres \
		-e POSTGRES_DB=postgres \
		-p 5432:5432 \
		$(DB_CONTAINER)

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
