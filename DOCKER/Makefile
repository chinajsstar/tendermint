build:
	@sh -c "'$(CURDIR)/build.sh'"

push:
	@sh -c "'$(CURDIR)/push.sh'"

build_develop:
	docker build -t "chinajsstar/tendermint:develop" -f Dockerfile.develop .

build_testing:
	docker build --tag tendermint/testing -f ./Dockerfile.testing .

push_develop:
	docker push "chinajsstar/tendermint:develop"

.PHONY: build build_develop push push_develop
