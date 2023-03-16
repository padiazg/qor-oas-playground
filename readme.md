# QOR OAS support

This project is for testing and developing the OAS document auto-generation from QOR resources declaration as it happens with the API

## Getting started
First you need to clone Portal Admin fork
```bash
cd ~/go/src/github.com/TykTechnologies/ 
git clone git@github.com:TykTechnologies/portal-admin.git
```

Next clone this repo
> Use your own folder if you preferr, but make sure you update the replace to match the folders accordingly
```bash
git clone git@github.com:padiazg/qor-oas-playground.git
```

Update go.mod to so the replace can reach the `portal-admin` module
```
replace github.com/TykTechnologies/portal-admin => ../../TykTechnologies/portal-admin
```

Now you should be ready to run the playground

## Play with it

Run the project with
```bash
go run main.go
```

The admin dashboard is at http://localhost:7100/admin

The auto-generated OAS document is at http://localhost:7100/api/oas

A StopLight viewer is at http://localhost:7100/oas/stoplight.html
> Just click refresh to load the document, the page is prepared to provide a token if the `/api` endpoint is protected