# AGENTS.md

## Build Requirements

- **Go 1.26.1** with `GOEXPERIMENT=jsonv2` environment variable required
- Build command: `GOEXPERIMENT=jsonv2 go build -v -o build_assets/v2node -trimpath -ldflags "-X 'github.com/wyx2685/v2node/cmd.version=$version' -s -w -buildid="`
- Docker build uses same `GOEXPERIMENT=jsonv2` flag
- No Makefile, linter config, or test files exist

## Dependencies

- Uses modified xray-core: `github.com/xtls/xray-core` replaced with `github.com/wyx2685/xray-core` via go.mod replace directive
- Run `go mod download` before building

## Architecture

- `cmd/`: Cobra CLI (server command entry point)
- `conf/`: Viper-based config loading (`/etc/v2node/config.json` default)
- `core/`: Xray-core integration and management
- `node/`: Node controller and API interaction
- `api/v2board/`: V2board panel API client
- `common/`: Utility packages (crypto, file, rate limiting)
- `limiter/`: Dynamic rate limiting
- `script/`: Install/management shell scripts

## Config Format

JSON with structure:
```json
{
  "Log": {"Level": "info", "Output": "", "Access": "none"},
  "Nodes": [{"ApiHost": "", "NodeID": 0, "ApiKey": "", "Timeout": 15}],
  "PprofPort": 0
}
```

## Key Entry Points

- `main.go` -> `cmd.Run()` -> `serverCommand` in `cmd/server.go`
- `core.New()` creates V2Core instance
- `node.New()` fetches node info from V2board API

## Runtime

- Watches config file changes and hot-reloads
- Uses logrus for logging
- pprof support via PprofPort config
- Graceful shutdown on SIGINT/SIGTERM