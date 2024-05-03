# QUBGradeMe

QUBGradeMe is a grade calculator system which was built using microservice architecture during my final year of university. Although the operations which can be performed by the system are simple, the process of building it was a valuable learning experience which helped me understand the fundamental concepts of cloud computing. 

Click [here](https://qubgrademe.up.railway.app/) to access a live deployment of QUBGradeMe.

## Contents

- [Services](#services)
  - [frontend](#frontend)
  - [proxy](#proxy)
  - [monitor](#monitor)
  - [database](#database)
  - [average](#average)
  - [totalmarks](#totalmarks)
  - [classify](#classify)
  - [classifymodules](#classifymodules)
  - [sortmodules](#sortmodules)
  - [maxmin](#maxmin)
- [What I Learned](#what-i-learned)

# Services

## frontend

The frontend service provides an interface for users to input values, perform calculations, save & retreive data, and analyse test results.

Use the following credentials to access the metrics dashboard within the frontend:

```
Username: admin.boss
Password: supersecret
```

## proxy

Implemented using **Go**, this service acts as a reverse proxy by receiving client requests from the frontend and forwarding them accordingly.

The following features are included:

- The proxy service is configured with a list of known routes to each service within the network. If the proxy is unable to deliver a request to a particular service, it will attempt to deliver the request using other configured routes for that service.
- Protected endpoints allow admins to dynamically add and remove routes to the configuration of a live proxy. Live instances of the proxy service will communicate with each other to ensure that a configuration update to one instance will be applied to all instances.
- When a new proxy spins up, it will reach out to an existing proxy and retrieve a list of known endpoints within the network. This ensures that a new proxy is immediately updated with the most recent configuration.

Below are example cURL commands for dynamically updating the proxy configuration.

Get all of the currently configured endpoints:
```bash
curl -X GET -u admin.boss:supersecret https://qubgrademe-proxy.up.railway.app/admin/routes | json_pp
```

Add a new endpoint to the configuration:
```bash
curl -X POST -u admin.boss:supersecret https://qubgrademe-proxy.up.railway.app/admin/routes\?service\=example\&route\=https://example-service.app
```

Delete an endpoint from the configuration:
```bash
curl -X DELETE -u admin.boss:supersecret https://qubgrademe-proxy.up.railway.app/admin/routes\?service\=example\&route\=https://example-service.app
```

## monitor

Implemented using **Go**, this service is used to run tests against all of the other services within the QUBGradeMe system, providing assurance that the application is healthy at all times.

The following features are included:

- The test suite can be executed manually from the metrics dashboard within the frontend.
- Tests will run automatically on a specified time interval.
- For each test failure, an email alert will be sent to the configured admin email address.
- The output of each test suite execution is stored in a MongoDB cluster.
- From the frontend metrics dashboard, historic test data can be conveniently retrieved and visualised using Chart.js.

## database

Implemented using **Go** and **MongoDB**, the database service allows users to save the data within their frontend input fields and retrieve it at any time using a unique identifier.

Saving with an existing unique identifier will overwrite the associated data.

## average

A **Go** service which calculates the average mark based off those which are provided.

This service also includes automated tests and a CI/CD configuration file from when it was originally orchestrated via GitLab.

## totalmarks

A **Python** service which calculates the total amount of marks obtained based off those which are provided.

This service also includes automated tests and a CI/CD configuration file from when it was originally orchestrated via GitLab.

## classify

A **Node** service which calculates the overall classification based off the marks which are provided.

This service also includes automated tests and a CI/CD configuration file from when it was originally orchestrated via GitLab.

## classifymodules

A **Java** service which calculates the classification for each module based off the marks which are provided.

This service also includes automated tests and a CI/CD configuration file from when it was originally orchestrated via GitLab.

## sortmodules

A **PHP** service which sorts the provided marks from highest to lowest.

This service also includes automated tests and a CI/CD configuration file from when it was originally orchestrated via GitLab.

## maxmin

A **PHP** service which returns the highest and the lowest marks based off those which are provided.

This service also includes automated tests and a CI/CD configuration file from when it was originally orchestrated via GitLab.

# What I Learned

- The fundamentals of microservice architecture.
- Cloud computing concepts such as load balancing, auto-scaling, redundancy, and containerisation.
- Technologies such as Docker and Kubernetes.
- The core principles of CI/CD.

