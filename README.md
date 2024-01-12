# Litecartes Backend

## Development Setup
1. Clone this repository
2. Navigate to the directory with ```$ cd litecartes-be```

You can use the manual here or bash script to automatically process some steps
### Bash Script
3. Change ```cmd/init.sh``` to be executeable with command ```$ chmod +x cmd/init.sh```
4. Run command ```$ cmd/init.sh```
5. After done, configure the rest 

### Manual
3. Install the dependencies needed with ```$ go mod download```
4. Setup a firebase project
5. Generate Firebase Admin SDK private key and put it in file ```config/litecartes-firebase-sdk.json```
6. Clone ```.env.example``` and rename it to ```.env```
7. Configure the env
8. Run the server with command ```$ go run app/main.go``` 

## Documentation
To read more about API documentation and system design, you can read more through this [documentation](./docs/DOCUMENTATION.md)

## Tech Stack
1. Golang
2. Gin
3. MySQL
4. Firebase Authentication