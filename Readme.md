
# Sertis - Blog APIs
Create by : Atid Chalermsong 
Email : jak.ka.ge@gmail.com

## Installation
Create docker image & container
```sh
docker-compose up -d
```
Create simple data
```sh
docker exec -it postgres psql -U postgres
```
```sh
\i data.sql
```
```sh
\q
```

## Postman collections
 https://www.getpostman.com/collections/1489019eb5534a1bd9d6


## API Spec

**1. Login API** -  get token to access blog api. 

```mermaid
sequenceDiagram
Login(Page) ->> Login(API): POST: /login 
Database -->> Login(API): Select password by username
Login(API) ->> Login(API): Validate username, password
alt Login data is incorrect
	Login(API) -->> Login(Page): 401 (Unauthorized)
else Login data is correct
	Database -->> Login(API): Select user info
	Login(API) -->> Login(Page): Send Token
end  
```

**Endpoint :** http://localhost/login

**Method :** POST

**Request header :** none

**Request body :** from-data
| Key  | Required | Remark (Use this value to login) |
| ------------- | ------------- | ------------- |
| username | required | jakub |
| password | required | 1234 |

**Request body Example :**
```sh
curl --location --request POST 'http://localhost/login' \
--form 'username="jakub"' \
--form 'password="1234"'
```
**Response body :**

| Key  | Remark |
| ------------- | ------------- |
| token | jwt token with user data | 

**Response body Example :**
```json
{
	"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjMiLCJuYW1lIjoiSmFrdWIgQW50bGsiLCJleHAiOjE2NTE5OTQ1MTJ9.VDIDHwCKBvk8ApiwHfoP0vUTihDc_ii7fKg9Oo2K71w"
}
```
\
**2. Create Card API** 

```mermaid
sequenceDiagram
Create Blog(Page) ->> Blog(API): POST: /blog/card 
Blog(API) ->> Database: Insert blog data
Blog(API) ->> Create Blog(Page): Status success/fail
```

**Endpoint :** http://localhost/blog/card

**Method :** POST

**Authorization:** Bearer Token 

**Request header :** 
| Key  | Required | Remark |
| ------------- | ------------- | ------------- |
|  Authorization | required | Get token form login api |

**Request body :** json
| Key  | Required | Remark  |
| ------------- | ------------- | ------------- |
| category | required | Category ID of card |
| title | required | Title of card |
| content | required | Content of card |

**Request body Example :**
```json
{
	"category": "1",
	"title": "Rising seas could submerge Rio **and** Jakarta **by** 2100",
	"content": "Aminath knows this all too well. As the environment and climate change minister for the Maldives, she is part of a community of politicians and scientists trying to work out how quickly sea levels will rise, if this can be slowed and what it means for us all. In some places, new ways of holding back the tide may buy us a few decades. Elsewhere, this won''t be possible. We are facing a disaster unfolding in slow motion. Responding effectively means a sea change in the way we think."
}
```
**Response body :**

| Key  | Remark |
| ------------- | ------------- |
| code | error code | 
| message | error reason | 
| requestId | request id for issue tracking | 

**Response body Example :**
Success case:
```json
{
	"code": "S000",
	"message": "Success",
	"requestId": "9af8d66d-7282-4bbf-9ebc-3b452ecadc21"
}
```
Fail case:
```json
{
	"code": "E000",
	"message": "Error xxx",
	"requestId": "9af8d66d-7282-4bbf-9ebc-3b452ecadc21"
}
```
\
**3. Get card list API**
```mermaid
sequenceDiagram
Blog Dashboard(Page) ->> Blog(API): GET: /blog/card/list 
Database -->> Blog(API): Select blog data
Blog(API) -->> Blog Dashboard(Page): Array list of card
```

**Endpoint :** http://localhost/blog/card

**Method :** GET

**Authorization:** Bearer Token 

**Request header :** 
| Key  | Required | Remark |
| ------------- | ------------- | ------------- |
|  Authorization | required | Get token form login api |

**Request body :** none

**Response body :**

| Key  | Remark |
| ------------- | ------------- |
| code | error code | 
| message | error reason | 
| requestId | request id for issue tracking | 
| result | array of card in json format | 

**Response body Example :**
Success case:
```json
{
	"code": "S000",
	"message": "Success",
	"requestId": "843a204a-5932-4a5d-b799-ff86d4c34f8b",
	"result": [
		{
			"id": "54926313-8755-4749-8ec4-0b8756448d27",
			"category": "Biology",
			"title": "How a soil microbe could rev up artificial photosynthesis",
			"status": "publish",
			"content": "\nPlants rely on a process called carbon fixation -- turning carbon dioxide from the air into carbon-rich biomolecules - for their very existence. \nThat's the whole point of photosynthesis, and a cornerstone of the vast interlocking system that cycles carbon through plants, \nanimals, microbes and the atmosphere to sustain life on Earth.",
			"author": "Jana Novkova",
			"create_timestamp": "2022-05-05 09:14:41",
			"update_timestamp": "2022-05-05 09:14:41"
		},
		{
			"id": "706ef181-7418-442a-a512-c2b92dc981ed",
			"category": "Engineering",
			"title": "Rising seas could submerge Rio and Jakarta by 2100",
			"status": "publish",
			"content": "\nAminath knows this all too well. As the environment and climate change minister for the Maldives, \nshe is part of a community of politicians and scientists trying to work out how quickly sea levels will rise, \nif this can be slowed and what it means for us all. In some places, new ways of holding back the tide may buy us a few decades. \nElsewhere, this won't be possible. We are facing a disaster unfolding in slow motion. Responding effectively means a sea change in the way we think. …",
			"author": "Terrza Konecna",
			"create_timestamp": "2022-05-05 09:14:41",
			"update_timestamp": "2022-05-05 09:14:41"
		}
	]
}
```
Fail case:
```json
{
	"code": "E000",
	"message": "Error xxx",
	"requestId": "9af8d66d-7282-4bbf-9ebc-3b452ecadc21"
}
```
\
**4. Update Card API** 

```mermaid
sequenceDiagram
Update Blog(Page) ->> Blog(API): PUT: /blog/card 
Blog(API) ->> Database: Update card data by cardId
Blog(API) ->> Update Blog(Page): Status success/fail
```

**Endpoint :** http://localhost/blog/card

**Method :** POST

**Authorization:** Bearer Token 

**Request header :** 
| Key  | Required | Remark |
| ------------- | ------------- | ------------- |
|  Authorization | required | Get token form login api |

**Request body :** json
| Key  | Required | Remark  |
| ------------- | ------------- | ------------- |
| id | required | ID of card |
| title | optional | Title of card |
| content | optional | Content of card |

**Request body Example :**
```json
{
	"id": "f529a8d3-475f-4490-b6c9-b87d7e17baa6",
	"title": "test222",
	"content": "test333"
}
```
**Response body :**

| Key  | Remark |
| ------------- | ------------- |
| code | error code | 
| message | error reason | 
| requestId | request id for issue tracking | 

**Response body Example :**
Success case:
```json
{
	"code": "S000",
	"message": "Success",
	"requestId": "9af8d66d-7282-4bbf-9ebc-3b452ecadc21"
}
```
Fail case:
```json
{
	"code": "E000",
	"message": "Error xxx",
	"requestId": "9af8d66d-7282-4bbf-9ebc-3b452ecadc21"
}
```
\
**5. Delete Card API** 

```mermaid
sequenceDiagram
Dashboard Blog(Page) ->> Blog(API): DELETE: /blog/card/:cardId
Blog(API) ->> Database: Delete card by cardId
Blog(API) ->> Dashboard Blog(Page): Status success/fail
```

**Endpoint :** http://localhost/blog/card/:cardId

**Method :** DELETE

**Authorization:** Bearer Token 

**Request header :** 
| Key  | Required | Remark |
| ------------- | ------------- | ------------- |
|  Authorization | required | Get token form login api |



**Request body Example :**
```json
{
	"id": "f529a8d3-475f-4490-b6c9-b87d7e17baa6",
	"title": "test222",
	"content": "test333"
}
```
**Response body :**

| Key  | Remark |
| ------------- | ------------- |
| code | error code | 
| message | error reason | 
| requestId | request id for issue tracking | 

**Response body Example :**
Success case:
```json
{
	"code": "S000",
	"message": "Success",
	"requestId": "9af8d66d-7282-4bbf-9ebc-3b452ecadc21"
}
```
Fail case:
```json
{
	"code": "E000",
	"message": "Error xxx",
	"requestId": "9af8d66d-7282-4bbf-9ebc-3b452ecadc21"
}
```
