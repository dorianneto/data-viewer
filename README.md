# Media Metadata
API to extract metadata from media using `ffprobe` under the hood.

## Usage
`$ docker compose up -d`

## API
1. `GET /`
2. `POST /metadata`
```
{
  "link": {{ mediaLink }}
}
```

### Examples:

```
$ curl -X GET http://127.0.0.1
```

```
$ curl -X POST -d '{"link":"https://images.rawpixel.com/image_png_800/cHJpdmF0ZS9sci9pbWFnZXMvd2Vic2l0ZS8yMDIyLTA4L3B4ODM3MDQxLWltYWdlLWpvYjg1MV8xLnBuZw.png"}' http://127.0.0.1/metadata
```
