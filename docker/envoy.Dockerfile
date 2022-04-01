FROM envoyproxy/envoy-dev:d4e28efb56a79703faffc18d6af1a578770af2f1

RUN apt-get update && apt-get install iproute2 -y

COPY ./entry.sh .
RUN chmod +x /entry.sh

ENTRYPOINT ["/entry.sh"]
CMD ["envoy", "-c", "/envoy/envoy.yaml"]
