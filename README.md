# Go-FlexGet #

[![wercker status](https://app.wercker.com/status/43149a32ca352251a19a7cbdfdaba20f/m "wercker status")](https://app.wercker.com/project/bykey/43149a32ca352251a19a7cbdfdaba20f)

## Introduction ##

**Go-FlexGet** is tiny web application displaying status, logs and configuration from [FlexGet](http://flexget.com/), on a different server.

## Usage ##

- Compile **Go-FlexGet**: `go build .`
- Modify FlexGet server configuration in `application.properties`
- Run **Go-FlexGet**: `./go-flexget -logtostderr=true`

## Components ##

- **Go**, version `1.4.2`
- **AngularJS**, version `1.3.14`
- **Angular Material**, version `0.8.3`
