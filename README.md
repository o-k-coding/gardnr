# gardnr

CLI used for my custom productivity workflow, including tending my digital gardens in the form of private notes as well as a "blog" style website.

This is evolving, and is very opinionated currently to fit my use cases!

Named in honor of how we pronounce Gardener in the midwest 🧑‍🌾

## Building

In this directory run

```bash
./internal/version/get_version.sh
go build -o .bin/gardnr ./cmd/gardnr
```

this will create a binary `gardnr`

## Installing

to install the binary in your local go bin

```bash
go install ./cmd/gardnr
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
- [] parse md files into a tree like data structure for easier querying
- [] use llm to generate description for posts
- [] allow user to define tags for posts
- [] use llm with RAG to generate tags for posts
- [] create a convention for saving files to cloud and templating Image components in the posts
  - [x] using obsidian image syntax
  - [] using normal markdown image syntax
- [] create an rpc server to allow for remote control of gardnr
- [] create VS Code plugin to work with gardnr-control
- [] use charm for more interactive cli

### Digital Garden

Create a post in the configured garden repo content folder GARDNR_GARDEN_REPO_PATH/GARDNR_GARDEN_REPO_CONTENT_PATH

examples:

```bash
gardnr create garden-post --description="Garden update 05-30-2024" --post-path garden --note life/garden/update-05-30-2024.md
gardnr create garden-post --description="Kesh Character Description" --post-path vennelos/characters/NPCs --note life/games/dnd/vennelos/NPCs/Edge of Night/Kesh.md
```

### Notes

Create a daily note from a template

```bash
gardnr create daily-note
```

### Translate

Uses the gcloud translate API

To set up in gcloud console <https://console.cloud.google.com/>

Run all commands in cloud shell

Create service account for authentication

```bash
gcloud iam service-accounts create \
    gardnr-translate --project \
    gardnr
```

Grant the service account the translate user role

```bash
gcloud projects \
    add-iam-policy-binding gardnr \
    --member \
    serviceAccount:gardnr-translate@gardnr.iam.gserviceaccount.com \
    --role roles/cloudtranslate.user
```

Generate a key

```bash
gcloud iam service-accounts keys \
    create translate-key.json \
    --iam-account \
    gardnr-translate@gardnr.iam.gserviceaccount.com
```

save the key json file locally - it is ignored in this repo

set the key file as your google app creds env var

```bash
export GOOGLE_APPLICATION_CREDENTIALS="$NOTES/development/gardnr/translate-key.json"
```

Now Locally

```bash
go get -u \
    cloud.google.com/go/translate/apiv3
```

### Count

Generate word count for a file

```bash
gardnr count words -f /some/file
```
