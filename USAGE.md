# Usage

## Docker

### Building

```bash
docker build -t shippy-service-consignment .
```

### Running

```bash
docker run -p 50051:50051 aetheryte-service-consignment
```

This can go through a differnet local port if we want it too. Changing `50051:50051` to `8080:50051` would open it on `8080` locally.
