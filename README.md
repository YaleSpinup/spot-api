# spot-api

This API provides simple restful API access to SpotInst Spot market services.

## Endpoints

```
GET /v1/spot/ping
GET /v1/spot/version
GET /v1/spot/metrics
```

## Authentication

Authentication is accomplished via a pre-shared key.  This is done via the `X-Auth-Token` header.

## Author

E Camden Fisher <camden.fisher@yale.edu>

## License

GNU Affero General Public License v3.0 (GNU AGPLv3)  
Copyright (c) 2019 Yale University
