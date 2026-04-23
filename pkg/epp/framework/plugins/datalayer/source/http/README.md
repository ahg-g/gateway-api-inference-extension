# HTTP Data Source

The HTTP Data Source is a base implementation for data layer sources that retrieve data via HTTP/HTTPS. It is designed to be embedded or used by specific data source plugins to handle the mechanics of HTTP polling and parsing.

## Features

-   Supports both `http` and `https` schemes.
-   Configurable TLS certificate verification (skip verification).
-   Pluggable response parsers.
-   Directly polls endpoints based on their addressable metadata.

## Usage in Other Plugins

The `metrics-data-source` uses `HTTPDataSource` as its underlying implementation, providing it with a Prometheus-specific parser.

## Implementation Details

The `HTTPDataSource` implements the `fwkdl.DataSource` interface:

-   `Poll(ctx, ep)`: Fetches data for a specific endpoint.
-   `OutputType()`: Returns the type of data produced after parsing.
-   `TypedName()`: Returns the plugin type and name.

It uses a `Client` interface to perform the actual HTTP requests, which allows for mocking in tests.
