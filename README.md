# spot-api

This API provides simple restful API access to SpotInst Spot market services.

## Endpoints

```
GET /v1/spot/ping
GET /v1/spot/version
GET /v1/spot/metrics

GET /v1/spot/{account}/elastigroups
POST /v1/spot/{account}/elastigroups
GET /v1/spot/{account}/elastigroups/{elastigroup}
PUT /v1/spot/{account}/elastigroups/{elastigroup}
DELETE /v1/spot/{account}/elastigroups/{elastigroup}

GET /v1/spot/{account}/instances
POST /v1/spot/{account}/instances
GET /v1/spot/{account}/instances/{instance}
PUT /v1/spot/{account}/instances/{instance}
DELETE /v1/spot/{account}/instances/{instance}
```

## Examples

### Elastigroups

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

### Managed Instances

Example request body of new managed instance (POST):

```json
{
  "region": "us-east-1",
  "name": "mi-test",
  "description": "My new Spot Instance",
  "strategy": {
    "revertToSpot": {
      "performAt": "always"
    }
  },
  "persistence": {
    "persistPrivateIp": true,
    "persistBlockDevices": true,
    "persistRootDevice": true,
    "blockDevicesMode": "onLaunch"
  },
  "healthCheck": {
    "type": "EC2",
    "autoHealing": true,
    "gracePeriod": 120,
    "unhealthyDuration": 120
  },
  "compute": {
    "product": "Linux/UNIX",
    "subnetIds": [
      "subnet-0123456789abcdef0"
    ],
    "vpcId": "vpc-0123456789abcdef0",
    "launchSpecification": {
      "instanceTypes": {
        "preferredType": "t3a.micro",
        "types": [
          "t3a.micro",
          "t3.micro",
          "t2.micro"
        ]
      },
      "securityGroupIds": [
        "sg-0123456789abcdef0"
      ],
      "imageId": "ami-02354e95b39ca8dec",
      "keyPair": null,
      "tags": [
        {
          "tagKey": "CreatedBy",
          "tagValue": "me"
        }
      ],
      "networkInterfaces": [
        {
          "associateIpv6Address": false,
          "deviceIndex": 0
        }
      ],
      "creditSpecification": {
        "cpuCredits": "standard"
      },
      "shutdownScript": null,
      "userData": null
    },
    "privateIp": null
  },
  "integrations": {},
  "scheduling": {}
}
```

Example request body of managed instance update (PUT):

```json
{
  "compute": {
    "launchSpecification": {
      "tags": [
        {
          "tagKey": "Environment",
          "tagValue": "dev"
        }
      ]
    }
  }
}
```

## Authentication

Authentication is accomplished using a pre-shared key (hashed string) in the `X-Auth-Token` header.

## Author

E Camden Fisher <camden.fisher@yale.edu>
Tenyo Grozev <tenyo.grozev@yale.edu>

## License

GNU Affero General Public License v3.0 (GNU AGPLv3)  
Copyright (c) 2020 Yale University
