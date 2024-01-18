# Python API

This python API was made to help us create question with the help of vertex AI 

## Setup

1. Create a project in google cloud
2. Copy the project id and set it to environment variables
3. Make sure you enabled Vertex API to be used in your google cloud project
4. Set the ```GO_API_URL``` environment, make sure it match the golang API endpoint that this project use
5. make a directory called ```config``` in ```python directory```
6. Store credential needed ```litecartes-gcloud.json``` to the config directory
7. Install packages needed in ```requirements.txt```
8. Run the python API with ```python3 main.py``` or ```python main.py``` if you used windows

## API DOCS

### 1. Generate Question to DB
Method : ```POST```      
Endpoint : ```/create/question/option```   
HTTP Response : 
- ```200 OK``   
- ```500 Internal Server Error```  

Response Body if OK
```
{
    "data": {
        "answer": "0",
        "category_id": "LTC-APP-generated1",
        "created_at": "0001-01-01T00:00:00Z",
        "literacy": "NFT merupakan singkatan dari non-fungible token, yaitu aset digital yang memiliki nilai unik dan tidak dapat ditukarkan dengan aset digital yang lain. NFT biasanya berupa karya seni digital, musik, atau video game. NFT dapat diperjualbelikan di berbagai platform online.",
        "option_list": null,
        "options": "platform online|pasar tradisional|pasar swalayan|toko online",
        "question": "NFT dapat diperjualbelikan di mana?",
        "task_uid": "",
        "title": "Apa itu NFT?",
        "uid": "LTC-APP-8d9619db-9307-4b59-b80c-c9458d8b59e6-QST",
        "updated_at": "0001-01-01T00:00:00Z"
    },
    "message": "successfully insert question",
    "payload": {
        "answer": "0",
        "category_id": "LTC-APP-generated1",
        "literacy": "NFT merupakan singkatan dari non-fungible token, yaitu aset digital yang memiliki nilai unik dan tidak dapat ditukarkan dengan aset digital yang lain. NFT biasanya berupa karya seni digital, musik, atau video game. NFT dapat diperjualbelikan di berbagai platform online.",
        "options": "platform online|pasar tradisional|pasar swalayan|toko online",
        "question": "NFT dapat diperjualbelikan di mana?",
        "title": "Apa itu NFT?"
    },
    "status": 200
}
```
