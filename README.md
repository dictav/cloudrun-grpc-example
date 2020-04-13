# CloudRun gRPC Auth Example

## Local test

```
$ make run
```

```
$ make test-local
```

## Deployment

```
$ make deploy PROJECT=YOUR-GCP-PROJECT-NAME
$ make test-gcp GRPC_ADDR=grpc-example-REPLACE_YOUR_URL-uc.a.run.app:443
```
