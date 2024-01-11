package firebase

import (
    "context"
    "log"

    "firebase.google.com/go/v4"
    "firebase.google.com/go/v4/auth"
    "google.golang.org/api/option"
)


func GetAuthClient() *auth.Client {
    opt := option.WithCredentialsFile("config/litecartes-firebase-sdk.json")
    app, err := firebase.NewApp(context.Background(), nil, opt)

    if err != nil {
        log.Fatalf("Failed to create firebase app. ERR:%s\n", err.Error())
    }

    auth, err := app.Auth(context.Background())
    if err != nil {
        log.Fatalf("Failed to create auth instance. ERR:%s\n", err.Error())
    }

    return auth 
}