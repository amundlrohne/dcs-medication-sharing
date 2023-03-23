
# MedShare: Medication sharing service 

### Built With

[![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)

[![React](https://img.shields.io/badge/react-%2320232a.svg?style=for-the-badge&logo=react&logoColor=%2361DAFB)](https://react.dev/)

![JavaScript](https://img.shields.io/badge/javascript-%23323330.svg?style=for-the-badge&logo=javascript&logoColor=%23F7DF1E)

[![MongoDB](https://img.shields.io/badge/MongoDB-%234ea94b.svg?style=for-the-badge&logo=mongodb&logoColor=white)](https://www.mongodb.com/)

[![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)](https://www.docker.com/)

[![Kubernetes](https://img.shields.io/badge/kubernetes-%23326ce5.svg?style=for-the-badge&logo=kubernetes&logoColor=white)](https://kubernetes.io/)

[![Terraform](https://img.shields.io/badge/terraform-%235835CC.svg?style=for-the-badge&logo=terraform&logoColor=white)](https://www.terraform.io/)

[FHIR](https://www.hl7.org/fhir/) 

## About 

### An easy to use cloud native platform for consent-based sharing medication records across borders using the FHIR standard.

![Project figure](/images/projectFig.jpg)

## API Reference

### Healthcare Provider Service

```http
POST /healthcare-provider
GET /healthcare-provider/name/:name
GET /healthcare-provider/:id
GET /healthcare-provider/all
POST /healthcare-provider/verify
GET /healthcare-provider/current
GET /healthcare-provider/getID/:token
DELETE /healthcare-provider/
``` 

### Medication Records Service 
```http
POST /medication-record
GET /medication-record
POST /medication-record/new
DELETE /medication-record
```

### Consent Service 
```http 
POST /consent/
GET /consent/from/:from_public_key
GET /consent/to/:to_public_key
GET /consent/

```

### Standardization Service 
```http 
GET /standardization/valid/:drugName
GET /standardization/drugNames/all
GET /standardization/:drugName
```



<!-- 
```http
  GET /api/items
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `api_key` | `string` | **Required**. Your API key |

#### Get item

```http
  GET /api/items/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |

#### add(num1, num2)

Takes two numbers and returns the sum. -->



