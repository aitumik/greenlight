# Greenlight

## Introduction
Greenlight is an API for handling movies


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