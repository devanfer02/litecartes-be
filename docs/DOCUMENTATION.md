# API Documentation

#### This is Litecartes Backend API Usage Documentatiton. All of the response body only displayed case API successfully proccess the request. Any other than that have almost same structure with different status code and message.
 
### 1. [Usecase Diagram](#usecase-diagram)
### 2. [Entity Relationship Diagram](#entity-relationship-diagram)
### 3. [API Endpoints](#api-endpoints)

## Usecase Diagram
![img](./usecase.drawio.png)

## Entity Relationship Diagram
##### NOTE: not all design are implemented in this api since we had to make some changes to make sure our system is fully integrated with client side
![img](./erd_desain.png)

## API Endpoints
### 1. [User](#user)
### 2. [Question](#question)

## User
### 1. Fetch All Users    
Method : ```GET```     
Endpoint : ```/user```   
HTTP Response :    
- ```200 OK```   
- ```500 Internal Server Error```   

Query
Field | Datatype | Description | Required |
--- | --- | --- | --- |
limit | integer | rows requested per page | default set to 10
cursor | string | encoded data to navigate page | required to get to next page |
next | boolean | indicating go to previous page if false, go to next page otherwise | default set to true

Response 
```
{
    "status": 200,
    "message": "successfully fech users",
    "data": [
        {
            "uid": "firebase_uniqueid",
            "username": "user",
            "email": "user@gmail.com",
            "subscription_id": 1,
            "school_id": 0,
            "total_exp": 0,
            "gems": 0,
            "streaks": 0,
            "last_active": "2024-01-11",
            "role": "_litecartes-app-user",
            "created_at": 1704987294,
            "updated_at": 1704987294
        }
    ],
    "pagination": {
        "prev_cursor": "eyJjcmVhdGVkX2F0IjoiMjAyNC0wMS0xMiIsInBvaW5fbmV4dCI6ZmFsc2UsImxpbWl0X2RhdGEiOjEwfQ==",
        "next_cursor": "eyJjcmVhdGVkX2F0IjoiMjAyNC0wMS0yMCIsInBvaW5fbmV4dCI6dHJ1ZSwibGltaXRfZGF0YSI6MTB9",
        "limit": 10
    }
}
```

### 2. Fetch User By UID
Method : ```GET```     
Endpoint : ```/user/:uid```   
HTTP Response :    
- ```200 OK```   
- ```500 Internal Server Error```   

Response Body
```
{
    "uid": "firebase_uniqueid",
    "username": "user",
    "email": "user@gmail.com",
    "subscription_id": 1,
    "school_id": 0,
    "total_exp": 0,
    "gems": 0,
    "streaks": 0,
    "last_active": "2024-01-11",
    "role": "_litecartes-app-user",
    "created_at": 1704987294,
    "updated_at": 1704987294
}
```

### 3. Register User
NB : this endpoint will store firebase authenticated user to mysql database, call the API whenever user registered with ```createUserWithEmailAndPassword``` method or ```signInWithGoogle`` once.
Method : ```POST```
Endpoint : ```/user/:uid```   
HTTP Response :    
- ```200 OK```   
- ```404 Not Found```
- ```409 Conflict```
- ```500 Internal Server Error```   

Response Body : 
```
{
    "status": 200,
    "message": "successfully register user"
}
```

### 4. Update User
Method : ```PUT```    
Endpoint : ```/user/:uid```   
HTTP Response :    
- ```200 OK``` 
- ```400 Bad Request```
- ```404 Not Found```
- ```500 Internal Server Error``` 

Json Required Payload  : 
Field | Datatype | 
--- | --- | 
username | string | 
email | string | 
subscription_id | integer | 
school_id | integer | 
total_exp | integer | 
gems | integer | 
streaks | integer | 


Response Body : 
```
{
    "status": 200,
    "message": "successfully update user"
}
```

### 5. Delete User
Method : ```DELETE```    
Endpoint : ```/user/:uid```   
HTTP Response :  
- ```200 OK```
- ```404 Not Found```
- ```500 Internal Server Error```

Response Body :
```
{
    "status": 200,
    "message": "successfully delete user"
}
```

## Question