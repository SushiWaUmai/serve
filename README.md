# Serve
Simple File Server written in go.

## How to use

```
version: "3"
services:
  serve:
    image: sushiwaumai/serve
    ports:
      - 3000:3000
    volumes: 
      - .:/files # this serves the . path on the file server
```
