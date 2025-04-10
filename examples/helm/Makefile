install:
ifndef HOST
	$(error HOST is undefined)
endif
	helm upgrade --install \
		signpost-example \
		--repo https://helm.sikalabs.io \
		signpost \
		--namespace signpost-example \
		--create-namespace \
		--values values.yaml \
		--set host=${HOST} \
		--wait

uninstall:
	helm uninstall \
		signpost-example \
		--namespace signpost-example
