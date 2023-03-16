# QOR OAS support

This project is for testing and developing the OAS document auto-generation from QOR resources declaration as it happens with the API

## Getting started
Clone this repo
```bash
git clone git@github.com:padiazg/qor-oas-playground.git
```
Now you should be ready to run the playground

```bash
go run main.go
```

## Play with it
The admin dashboard is at http://localhost:7100/admin

The auto-generated OAS document is at http://localhost:7100/api/oas

A StopLight viewer is at http://localhost:7100/oas/stoplight.html
> Just click refresh to load the document, the page is prepared to provide a token if the `/api` endpoint is protected