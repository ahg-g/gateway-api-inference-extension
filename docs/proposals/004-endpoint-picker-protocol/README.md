# Endpoint Picker Protocol

The Endpoint Picker, or EPP, is a core component of the inference extension. Ultimately it's
responsible for picking an endpoint from the `InferencePool`. A reference implementation can be
found [here](../../../pkg/epp/).

This doc defines the protocol between the EPP and the proxy (e.g, Envoy).

The EPP MUST implement the Envoy
[external processing service](https://www.envoyproxy.io/docs/envoy/latest/api-v3/service/ext_proc/v3/external_processor) protocol.

## Proxy Request
For each HTTP request, the proxy CAN community to the proxy the subset of endpoints the EPP MUST pick from by setting an unstructured in the [filtered metadata](https://github.com/envoyproxy/go-control-plane/blob/63a55395d7a39a8d43dcc7acc3d05e4cae7eb7a2/envoy/config/core/v3/base.pb.go#L819) field of the ext-proc request. The metadata entry for the subset MUST be wrapped with an outer key (which represents the metadata namespace) with a default of `envoy.lb.subset_hint`.

```go
filterMetadata: {
  "envoy.lb.subset_hint" {
     "x-gateway-destination-endpoint-subset-hint": [<ip:port>, <ip:port>, ...]
  }
}
```

## EPP Response
For each HTTP request, the EPP MUST communicate to the proxy the picked model server endpoint via:

1. Setting the `x-gateway-destination-endpoint` HTTP header to the selected endpoint in <ip:port> format.

2. Set an unstructured entry in the [dynamic_metadata](https://github.com/envoyproxy/go-control-plane/blob/c19bf63a811c90bf9e02f8e0dc1dcef94931ebb4/envoy/service/ext_proc/v3/external_processor.pb.go#L320) field of the ext-proc response. The metadata entry for the picked endpoints MUST be wrapped with an outer key (which represents the metadata namespace) with a default of `envoy.lb`.

The pirmary endpoint MUST be set using the key `x-gateway-destination-endpoint` as follows:
```go
dynamicMetadata: {
  "envoy.lb": {
    "x-gateway-destination-endpoint": <ip:port>
  }
}
```

Fallback endpoints MUST be set using the key `x-gateway-destination-endpoint-fallbacks` as a list as follows:
```go
dynamicMetadata: {
  "envoy.lb" {
     "x-gateway-destination-endpoint-fallbacks": [<ip:port>, <ip:port>, ...]
  }
}
```

Note:
- If the EPP did not communicate the server endpoint via these two methods, it MUST return an error.
- The EPP MUST not set two different values in the header and the inner response metadata value. 
- Setting different value leads to unpredictable behavior because proxies aren't guaranteed to support both paths, and so this protocol does not define what takes precedence.

### Why envoy.lb namespace as a default? 
The `envoy.lb` namesapce is a predefined namespace used for subsetting. One common way to use the selected endpoint returned from the server, is [envoy subsets](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/upstream/load_balancing/subsets) where host metadata for subset load balancing must be placed under `envoy.lb`.



