# Thumbnail Previewer Backend

This project is a backend API for a thumbnail previewer application. It is deployed as an AWS Lambda function and can be accessed through an API Gateway. The API allows users to fetch YouTube channel information, including the channel name and profile picture URL, by providing the channel handle.

## Table of Contents

- [Features](#features)
- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Running Locally](#running-locally)
  - [Environment Variables](#environment-variables)
- [Deployment](#deployment)
- [API Endpoints](#api-endpoints)
- [License](#license)

## Features

- Fetch YouTube channel information using the channel handle.
- Deployed as an AWS Lambda function.
- Accessed through an API Gateway.

## Project Structure

```
.env
.env.json.example
.github/
    workflows/
        prod-deploy.yml
.vscode/
api/
    main.go
bin/
    thumbnail_previewer_backend
    thumbnail_previewer_backend.zip
Dockerfile
env.json
go.mod
go.sum
internal/
    config/
        config.go
    handlers/
        channel_handler.go
        ping.go
    models/
        channel.go
    responses/
        responses.go
        youtube_response.go
    services/
        youtube_service_test.go
        youtube_service.go
LICENSE
Makefile
template.yaml
```

## Getting Started

### Prerequisites

- Go 1.22 or later
- Make
- AWS CLI (for deploying the Lambda function)
- AWS SAM CLI (for running the Lambda function locally)
- A Youtube API Key (console.cloud.google.com)

### Installation

Clone the repository:

```sh
git clone https://github.com/pedropassos06/thumbnail-previewer-backend.git
cd thumbnail-previewer-backend
```

Install dependencies:

```sh
go mod tidy
```

### Running Locally

Create a `.env.json` file in the root directory with the following content:

```sh
{ 
    "GetThumbnailPreviewFunction": { 
        "YOUTUBE_API_KEY": "YOUR_YOUTUBE_API_KEY"
    }
}
```

Build and run the application:

```sh
make run
```

The application will be running locally, and you can access the endpoints.

## Deployment

The application is deployed as an AWS Lambda function using GitHub Actions. The deployment process is defined in the `prod-deploy.yml` file.

Ensure you have the necessary AWS credentials set up in your GitHub repository secrets:

- `AWS_ACCESS_KEY_ID`
- `AWS_SECRET_ACCESS_KEY`
- `AWS_REGION`
- `YOUTUBE_API_KEY`

Push your changes to the main branch, and the GitHub Actions workflow will automatically build, test, and deploy the application.

## API Endpoints

### Get Channel

- **URL:** `/channel`
- **Method:** `GET`
- **Query Parameters:**
  - `handle`: The YouTube channel handle.
- **Description:** Fetches the YouTube channel information, including the channel name and profile picture URL.

### Ping

- **URL:** `/ping`
- **Method:** `GET`
- **Description:** Returns a simple "pong" message to check if the API is running.

## License

This project is licensed under the MIT License. See the License file for details.