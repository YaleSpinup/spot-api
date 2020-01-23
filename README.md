# spot-api

This API provides simple restful API access to SpotInst Spot market services.

## Endpoints

```
GET /v1/spot/ping
GET /v1/spot/version
GET /v1/spot/metrics
GET /v1/spot/{account}/elastigroups
POST /v1/spot/{account}/elastigroups
PUT /v1/spot/{account}/elastigroups/{elastigroup}
DELETE /v1/spot/{account}/elastigroups/{elastigroup}
```

## Examples

Example request body of new elastigroup (POST):

```json
{
  "name": "test202018",
  "strategy": {
    "risk": 100,
    "onDemandCount": null,
    "availabilityVsCost": "balanced",
    "drainingTimeout": 120,
    "fallbackToOd": true,
    "scalingStrategy": {
      "terminationPolicy": "newestInstance"
    },
    "lifetimePeriod": "days",
    "revertToSpot": {
      "performAt": "always"
    },
    "persistence": {}
  },
  "capacity": {
    "target": 2,
    "minimum": 2,
    "maximum": 4,
    "unit": "instance"
  },
  "scaling": {
    "up": null,
    "down": null
  },
  "compute": {
    "instanceTypes": {
      "ondemand": "t3.micro",
      "spot": [
	    "t3.micro",
	    "t3a.micro",
	    "t3a.nano",
	    "t3.nano"
      ]
    },
    "availabilityZones": [
      {
        "name": "us-east-1a",
        "subnetId": "subnet-00000000000000000"
      }
    ],
    "product": "Linux/UNIX",
    "launchSpecification": {
      "securityGroupIds": [
        "sg-00000000000000000"
      ],
      "monitoring": false,
      "ebsOptimized": false,
      "imageId": "ami-00000000000000000",
      "keyPair": "user",
      "healthCheckType": null,
      "tenancy": "default",
      "userData": null,
      "shutdownScript": null
    },
    "elasticIps": null,
    "preferredAvailabilityZones": null
  },
  "multai": null,
  "scheduling": {},
  "region": "us-east-1",
  "thirdPartiesIntegration": {}
}
```

Example request body of elastigroup update (PUT):

```json
{
  "id": "sig-000000000",
  "capacity": {
    "target": 2,
    "minimum": 2,
    "maximum": 4
  },
  "compute": {
    "instanceTypes": {
      "ondemand": "t3.micro",
      "spot": [
  	    "t3.micro",
	    "t3a.micro",
	    "t3a.nano",
	    "t3.nano"
      ]
    },
    "launchSpecification": {
      "tags": [
        { 
          "tagKey": "food",
          "tagValue": "good"
        }
      ]
    }
  }
}
```

## Authentication

Authentication is accomplished via a pre-shared key.  This is done via the `X-Auth-Token` header.

## Author

E Camden Fisher <camden.fisher@yale.edu>

## License

GNU Affero General Public License v3.0 (GNU AGPLv3)  
Copyright (c) 2019 Yale University
