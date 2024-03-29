# Litecartes Backend

![mockup](./docs/mockup1.png)

Litecartes is an application made to aim address Sustainable Development Goals point 4, Quality Education. With our gamification application product, we aim to increase the literacy level and critical thinking of the future generations

## Team Members
#### Team Name : ```Kusushi```

Name | Role
-- |--
Devan Ferrel | Hacker
I Gusti Ngurah Ryo Adi Tarta | Hacker
Wilson Hamonangan Ariyanto Hutapea  | Hustler
Sarah Safira Millati | Hipster

## Server Side Development Setup
To get started with setup litecartes backend, you can follow the following steps.

1. Clone this repository
2. Navigate to the directory with ```$ cd litecartes-be```

You can use the manual here or bash script to automatically process some steps

#### Bash Script
3. Change ```cmd/init.sh``` to be executeable with command ```$ chmod +x cmd/init.sh```
4. Run command ```$ cmd/init.sh```
5. After done, configure the rest like ```.env``` and firebase admin sdk config in ```config/litecartes-firebase-sdk.json```
6. If done, then run the server with command ```$ go run app/main.go```

#### Manual
3. Install the dependencies needed with ```$ go mod download```
4. Setup a firebase project
5. Generate Firebase Admin SDK private key and put it in file ```config/litecartes-firebase-sdk.json```
6. Clone ```.env.example``` and rename it to ```.env```
7. Configure the env
8. Run the server with command ```$ go run app/main.go``` 

## Credentials Needed
1. ```config/litecartes-firebase-sdk.json``` : firebase sdk credential
2. ```python/config/litecartes-gcloud.json``` : google cloud key for vertex api

## Documentation
To read more about API documentation and system design, you can read more through this [documentation](./docs/DOCUMENTATION.md)

## Client Side
For client side repository, you can take a look in this [github repository](https://github.com/ryoaditarta/litecartesProject-FrontEnd)

## CMS
Litecartes also provides content management system to interact with server like adding item, editing item, etc easier, you can take a look at our CMS in this [folder](./cms) 

Sidenote: the CMS supposed to be used in desktop environment since its not developed for mobile and tablet.

## Tech Stack
1. Golang
2. Gin
3. MySQL
4. Firebase Authentication
