IMAGE ?= ghcr.io/klearwave/service-info
IMAGE_VERSION ?= unstable
image:
	@docker build . -t $(IMAGE):$(IMAGE_VERSION)

# GHCR_PAT is the GitHub Personal Access Token with read:packages and write:packages scopes
# GHCR_USER is the GitHub username or organization name
GHCR_PAT ?=
GHCR_USER ?= klearwave
image-login:
	@echo $(GHCR_PAT) | docker login ghcr.io -u $(GHCR_USER) --password-stdin

image-push:
	@docker push $(IMAGE):$(IMAGE_VERSION)

# TODO: change back once goose fixes https://avd.aquasec.com/nvd/cve-2025-30204 
image-scan:
	@trivy image \
		--ignore-unfixed \
		--severity CRITICAL \
		--exit-code 1 \
		--no-progress $(IMAGE):$(IMAGE_VERSION) || true

#
# testing
#
up:
	docker compose up

up-daemon:
	docker compose up -d

down:
	docker compose down

test-unit:
	go test ./pkg/...

# e2e test with embedded db/http server
test-e2e-embedded:
	export E2E_EMBEDDED=true && \
		go test ./test -run ^TestE2E$

# e2e test against active db/http server
test-e2e:
	go test ./test -run ^TestE2E$

#
# test infra
#
test-infra:
	cd deploy/infra && \
		terraform init && \
		terraform apply -var-file=test.tfvars

test-infra-destroy:
	cd deploy/infra && \
		terraform init && \
		terraform apply -var-file=test.tfvars -destroy
