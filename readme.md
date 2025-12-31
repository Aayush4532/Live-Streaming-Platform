# Live Streaming Web Server (MVP)

This repository contains the web server for our live streaming platform.  
Its primary responsibility is to authenticate users/hosts and securely provide access to live streams.

This project is currently being built as an **MVP** for the **Prashant Advait Foundation**.  
In the coming months, it will be upgraded to a full **production-grade system**.

---

## What This Server Does

- Issues short-lived signed URLs for live stream playback
- Works with CDN (Cloudflare) for secure and scalable delivery
- Acts as an access-control layer (does not serve media directly)

Future additions planned:
- Live chat
- Coding / practice platform
- Production hardening and monitoring

---

## Basic Architecture

Client → Web Server → CDN → Media Server

The web server only validates access and generates signed playback URLs.

---

## How to Use

### Environment Variables

```bash
export STREAM_URI="https://live.example.com"
export STREAM_SIGNING_SECRET="your_secret_key"
The signing secret must match the CDN configuration.
```

Run the Server
```bash
go run main.go
```

----
Stream Access API
POST /join-seminar
----------------------------------------------------

Returns a short-lived signed HLS URL for playback.

Example response:
```bash
{
  "playback": {
    "url": "https://live.example.com/hls/stream.m3u8?exp=...&sig=..."
  },
  "refresh_in": 120
}
```


---