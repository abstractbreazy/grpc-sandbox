FROM envoyproxy/envoy-dev:f2023ef77bdb4abaf9feef963c9a0c291f55568f

CMD ["envoy", "-c", "/envoy/envoy.yaml"]
