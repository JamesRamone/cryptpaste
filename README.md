# cryptpaste
A sample pastebin-like app with client side-encryption.

## Server
The server is built in Go and uses Chi router.  
The *pastes* are stored in BoltDB instance.

## Web frontend
The frontend is built with Vue.js and Nuxt.js.  
Tailwind CSS is used fo styling.  
Client side encryption relies on `crypto-js`.  

## Deployment
Both the server and web frontend are deployed as Docker containers behind a reverse proxy.
