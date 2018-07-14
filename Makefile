image=quay.io/cloudflavor/ci-pipelines
vers=v0.1.0-alpha
gobuild=go build -o _output/bin/ci-pipelines
commit=$(shell git rev-parse --short HEAD)
tag=$(shell git describe --abbrev=0 --tags)

build:
	rm -rf _output || true
	mkdir -p _output/bin/
	 $(gobuild) -v -ldflags "-X main.commit=$(commit) -X main.version=$(vers)" ./cmd/cip/main.go

generate:
	echo "test"

publish:
	docker build -t $(image):$(commit) .
	docker push $(image):$(commit)
	docker tag $(image):$(commit) $(image):latest
	docker push $(image):latest

push: tag
	docker build -t $(image):$(commit) .
	docker push $(image):$(commit)
	docker push $(image):$(vers)
	docker push $(image):latest
	git push --tags

tag:
	docker tag $(image):$(commit) $(image):latest
	docker tag $(image):$(commit) $(image):$(vers)
	git tag $(vers)

test:
	go test -race -coverprofile=coverage.txt -covermode=atomic -v ./pkg/pipelines/

gen:
	vendor/k8s.io/code-generator/generate-groups.sh all github.com/PI-Victor/ci-pipelines/pkg/client  github.com/PI-Victor/ci-pipelines/pkg/apis  pipelines:v1
