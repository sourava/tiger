# Tigerhall Kittens

Tigerhall Kittens is a small web app for tracking the population of tigers in the wild

## Prerequisite
1. Installation instructions are for MacOS, please install go and make for working with the repo.

Install Golang
```bash
brew update&& brew install golang
```

Install Make
```bash
brew install make
```

2. Sendgrid is used to send email notifications.
   Please create a sendgrid account, single sender in the dashboard.
   And set these env variables in a .env file
```
SENDGRID_API_KEY=<YOUR_API_KEY>
SENDGRID_SENDER_EMAIL=<YOUR_SINGLE_SENDER_NAME>
SENDGRID_SENDER_NAME=<YOUR_SINGLE_SENDER_EMAIL>
```


## Usage

```
# run the app
make up

# run e2e tests
make e2e-tests

# generate fresh swagger docs
make generate-swagger-doc

```

## Swagger
This app uses swagger for documenting REST API's.
```
http://www.localhost:8080/swagger/index.html
```

## Design decisions made

1. This app accepts base64 encoded string as image in create tiger sightings api, and resizes and saves it in the tiger_sightings table.
   Ideally this resizing request needs to be moved to a queue which asynchronously resizes it and pushes it to s3 and updates the file location in db. But due to time constraints this was not implemented, and can be taken up for future improvements.

2. This app uses Golang's channels for PUB/SUB tasks. Ideally this can me moved to a centralised queue's like SQS, RabbitMQ.

3. Directory Structure:
- app: Idea behind app directory is to keep all handlers and server.go which is the separate from business logic, it does db initialisation and processes notification messages.
- business: Idea behind business directory is to keep all business specific files. It contains models and services for creating tigers, users, auth
- external: Idea behind external directory is to keep all client wrappers and helpers that does not contain any business logic.
