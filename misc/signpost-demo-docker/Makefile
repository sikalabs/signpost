dev:
	yq config.yml -o json > config.json
	signpost server -c config.json

docker-build:
	docker build --platform linux/amd64 -t sikalabs/signpost-demo .

docker-push:
	docker push sikalabs/signpost-demo

docker-run:
	docker run -p 8000:8000 sikalabs/signpost-demo

setup-git-hooks:
	rm -rf .git/hooks
	(cd .git && ln -s ../.git-hooks hooks)
