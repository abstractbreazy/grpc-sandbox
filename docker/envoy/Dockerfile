FROM envoyproxy/envoy-dev:0264041ec7a271c0291aa2d761856045c90aef7e

RUN apt-get update && apt-get install iproute2 -y

COPY /docker/envoy/entry.sh .
RUN chmod +x /entry.sh

ENTRYPOINT ["/entry.sh"]
CMD ["envoy", "-c", "/etc/envoy/envoy.yaml"]

