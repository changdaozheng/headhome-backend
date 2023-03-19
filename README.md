# HeadHome 

## Project Description

### The Team

| Members                                             | 
| --------------------------------------------------- | 
| [Chay Hui Xiang](https://github.com/chayhuixang)    | 
| [Chang Dao Zheng](https://github.com/changdaozheng) |
| [Marc Chern Di Yong](https://github.com/Trigon25)   | 
| [Ong Jing Xuan](https://github.com/ongjx16)         |

## HeadHome (Backend)
This repo contains the backend application for HeadHome.
Go `(Version 1.19+)` must be installed to run the application.

### Tech Stack
1. [Go](https://go.dev/)
2. [Firebase](https://firebase.google.com/)

### Quick Start:
1. Run the following code in bash to install the required dependencies
```
go get all
```
2. Create a `.env` file and insert your Firebase Admin SDK private key and Maps API api key. 
<br>
<font color="#888888">
    Note: Place the entire Firebase Admin SDK private key json object on a single line and escape all `\` and `\n` characters with `\`. Lastly, surround the json object with double quotations `""`.
</font>

```css
/*.env file*/
FIREBASE_ADMIN_PRIVATE_KEY=<your inline firebase admin private key>
MAPS_API_KEY=<your maps api key>
```

### Scripts
Run `go run ./cmd` to launch the server.

### File structure
```tree
├── cmd
│   └──main.go
├── controllers
│   ├── care_giver_controller.go
│   ├── care_receiver_controller.go
│   ├── map_controller.go
│   ├── sos_controller.go
│   ├── travel_log_controller.go
│   └── volunteers_controller.go
├── database
│   ├── care_giver_collection.go
│   ├── care_receiver_collection.go
│   ├── database.go
│   ├── sos_log_collection.go
│   ├── travel_log_collection.go
│   └── volunteers_collection.go
├── logic
│   ├── direction.go
│   └── util.go
├── models
│   ├── care_giver.go
│   ├── care_receiver.go
│   ├── sos_log.go
│   ├── travel_log.go
│   └── volunteers.go
├── routes
│   └── routes.go
├── websocket.go
│   ├── client.go
│   ├── msg_pump.go
│   ├── websocket.go
│   └── ws_hub.go
├── .env (not included in github repo)
├── .gitignore
├── .replit
├── go.mod
├── go.sum
├── Makefile
├── README.md
└── replit.nix
```
