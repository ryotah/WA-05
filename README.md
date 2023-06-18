# WA-05

Panic app

## Run

```sh
go run main.go
```

## Deploy

```sh
# Create your demo project
gcloud projects create <PROJECT_ID>
gcloud config set project <PROJECT_ID>

# gcloud services enable cloudbuild.googleapis.com
# ERROR: (gcloud.services.enable) FAILED_PRECONDITION: Billing account for project 'xxx' is not found.
# => I enabled the API using Google Cloud console.

gcloud builds submit --tag gcr.io/<PROJECT_ID>/panic-demo

gcloud run deploy panic-demo --image gcr.io/<PROJECT_ID>/panic-demo --allow-unauthenticated

# Delete your demo project
gcloud projects delete <PROJECT_ID>
```