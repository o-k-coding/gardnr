# grdnr

CLI used for my custom productivity workflow.

This is evolving, and is very opinionated currently to fit my use cases!

## Building

In this directory run

```bash
./internal/version/get_version.sh
go build -o .bin/grdnr ./cmd/grdnr
```

this will create a binary `grdnr`

## Installing

to install the binary in your local go bin

```bash
go install ./cmd/grdnr
```

this will build the binary and then provide it in your go bin

## Testing

### Running unit/integration tests

```bash
go test ./...
```

to run an individual test

```bash
go test ./... -run "TestTranslateText"
```

Be warned!! the translate test calls the translate api and there is a free tier limit of 100k chars/month

## Usage

## Project structure

Based off
<https://github.com/golang-standards/project-layout>

and

<https://github.com/spf13/cobra/blob/master/user_guide.md>

## Tasks

- [] Makefile

## Bugs

- [] If you don't specify a post-name for create garden-post and the template file does not follow the format x.ext.tmpl then it will not work.
- [] Translate tests are failing

## Features

### Wishlist

- [] setup command to help get local environment setup for managing notes and digital garden
- [] use llm to generate description for posts
- [] allow user to define tags for posts
- [] use llm with RAG to generate tags for posts
- [] create a convention for saving files to cloud and templating Image components in the posts
- [] create an rpc server to allow for remote control of grdnr
- [] create VS Code plugin to work with grdnr-control
- [] use charm for more interactive cli

### Digital Garden

Create a post in the configured garden repo content folder GRDNR_GARDEN_REPO_PATH/GRDNR_GARDEN_REPO_CONTENT_PATH

```bash
grdnr create garden-post --description="Garden update 05-30-2024" --post-path garden --note life/garden/update-05-30-2024.md
```

### Notes

Create a daily note from a template

```bash
grdnr create daily-note
```

### Translate

Uses the gcloud translate API

To set up in gcloud console <https://console.cloud.google.com/>

Run all commands in cloud shell

Create service account for authentication

```bash
gcloud iam service-accounts create \
    grdnr-translate --project \
    grdnr
```

Grant the service account the translate user role

```bash
gcloud projects \
    add-iam-policy-binding grdnr \
    --member \
    serviceAccount:grdnr-translate@grdnr.iam.gserviceaccount.com \
    --role roles/cloudtranslate.user
```

Generate a key

```bash
gcloud iam service-accounts keys \
    create translate-key.json \
    --iam-account \
    grdnr-translate@grdnr.iam.gserviceaccount.com
```

save the key json file locally - it is ignored in this repo

set the key file as your google app creds env var

```bash
export GOOGLE_APPLICATION_CREDENTIALS="$NOTES/development/grdnr/translate-key.json"
```

Now Locally

```bash
go get -u \
    cloud.google.com/go/translate/apiv3
```

### Count

Generate word count for a file

```bash
grdnr count words -f /some/file
```