#!/bin/bash

curl -i -N -H "Connection: keep-alive, Upgrade" \
	-H "Upgrade: websocket" \
	-H "Sec-WebSocket-Version: 13" \
	-H "Sec-WebSocket-Extensions: deflate-stream" \
	-H "Sec-WebSocket-Key: WIY4slX50bnnSF1GaedKhg==" \
	-H "Host: localhost:1333" \
	-H "Origin:http://localhost:1333" http://localhost:1333/ws
