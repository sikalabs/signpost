FROM sikalabs/ca-certificates as yq
COPY --from=sikalabs/slu:v0.63.0-dev-1 /usr/local/bin/slu /usr/local/bin/slu
RUN slu install-bin yq

FROM sikalabs/signpost:v0.2.0
COPY --from=yq /usr/local/bin/yq /usr/local/bin/yq
COPY config.yml /
RUN yq config.yml -o json > config.json
CMD [ "signpost", "server", "-c", "/config.json" ]
EXPOSE 8000
