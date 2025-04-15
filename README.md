# Greenlight

## Introduction
Greenlight is an API for handling events and movies

## Setup 

Setting up the database

Create the volume
```bash
mkdir -p ${HOME}/postgresql/data
```

Create network
```bash
docker network create greenlight-net
```

Start postgres
```bash
docker run -it -d --name greenlightdb -e POSTGRES_USER=greenlight -e POSTGRES_DB=greenlight -e POSTGRES_PASSWORD=greenlight  -v ${HOME}/postgresql/data:/var/lib/postgresql -p 5432:5432 --net greenlight-net postgres:latest
```

## Events API

```bash
curl -X POST http://localhost:8000/v1/events \
  -H "Content-Type: application/json" \
  -d '{
        "start_time": "2025-04-15T14:00:00Z",
        "end_time": "2025-04-15T16:00:00Z",
        "title": "Sample Event",
        "description": "This is a description of the event.",
        "venue": "Sample Venue",
        "tags": ["Tag1", "Tag2"],
        "cover": "http://example.com/cover.jpg"
      }'
```
