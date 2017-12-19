Per Score Server!
===================

This service is a primary interface that provide Routing between services and act as a controller. It implements both **REST** and **GRPC** calls in order to interacts between clients and its internal services

----------


Description
-------------

This application uses Go lang **NET/HTTP** package for starting server and making **REST** call between client and the server. These are following features of the application in terms of functionality .

> **Note:**
> - You must load ```.sh``` file before running any of the go cobra command  so that it can load all the environment variable and setup postgres at the same time.
> - Sometimes the bash does not load the environment ```.env``` file form the bash ```.sh``` file, if there is any issue while running the code than you must fetch the environment file by the following command ```source.env``` before running any of the go cobra command.
> - Make sure you have all the required and correct environment variable available before running the service.
> - Make sure that you have postgres installed in your machine.

#### <i class="icon-file"></i> Login

This is one of the route that is used by the client to check its authenticity in order to get access as a questioner or responder . The user will fill in the details  which will hit the login controller using **REST** call and after unmarshalling the json the unmarshalled data will be forwarded to the Authantication service (PerSocreAuth) using **GRPC** call . The response from the Authantication service contaning token will be  forwarded back to Per Score Server where it will be saved in the local database and marshalled so that it can be converted into json and given back to client as json object.

#### <i class="icon-folder-open"></i> Register

This route is used by the client to register himself as admin, questioner or responder . The user will fill in the details  which will hit the Register controller using **REST** call and after unmarshalling the unmarshalled data will be forwarded to the Authantication service (PerSocreAuth) using **GRPC** call . The response from the Authantication service will be given back to Per Score Server where it will be marshalled and converted into json and given back to client as json object.

#### <i class="icon-pencil"></i> Create question

Once the Questioner is logged in, it will have a right to create a question which will be later approved by the administrator. When a user will fill in the details and submit a question it will hit the create_question controller in per score server using a **REST** call  which will again  unmarshalling the json and the unmarshalled data will be forwarded to Per score cal server using **GRPC** call . The response from per score cal service will be given back to Per Score Server where it will be marshalled and converted into json and given back to client as json object.

#### <i class="icon-sort"></i> Fetch question

Once the Respondent is logged in, it will have a right to attempt for the question on the basis of the choosen category. When a user will submit an answer it will hit the get_question controller in per score server using a **REST** call  which will  unmarshalling the json and the unmarshalled data will be forwarded to Per score cal server using **GRPC** call . The response from per score cal service will be given back to Per Score Server where it will be marshalled and converted into json and given back to client as json object.

#### <i class="icon-file"></i>   Approve Category and Question

Once the Administrator is logged in, it will have a right to see all the questions and the category created by the questioner. Administrator will have a right to accept the question as well as the category . When an Administrator will approve the question or the cateory it will hit the approve_entries controller in per score server using a **REST** call  which will  unmarshalling the json and the unmarshalled data will be forwarded to Per score cal server using **GRPC** call . The response from per score cal service will be given back to Per Score Server where it will be marshalled and converted into json and given back to client as json object.

> **Github URL:**  [<i class="icon-download"></i> perScoreServer](#https://github.com/praveenraturi94/perScore)


----------


Dependencies
-------------------
1> Development packages

* [Postgresql](https://wiki.westfieldlabs.com/display/WL/PostgreSQL)
* [GORM](https://github.com/jinzhu/gorm)
* [Go](https://wiki.westfieldlabs.com/display/WL/Go)
* [GRPC](https://google.golang.org/grpc)

2> Testing packages

* [GINKO](https://github.com/onsi/ginkgo)
* [GOMEGA](https://github.com/onsi/gomega)

----------


Build and run this project
-------------

>1. To give privilege to ur .sh file
    ```
    chmod +x setupPostgres.sh
    ```
2. Run .sh file to create role and database
    ```
    ./setupPostgres.sh
    ```
3. Run command to migrate database
    ```
    go run main.go createDB
    ```
4. Run command to start server
    ```
    go run main.go serve
    ```

----------


Project Folder Structure Description
--------------------

1> APP :- The APP'S based directory will contain all the business logic including routing and controller
* Routes :-  It contain all the routes that are used by client as a rest call.
* Controller :- Controllers process incoming requests, handle user input and interactions, and execute appropriate application logic.
* Model :- Models contain the struct that represent data relation (table) and potentially bridge tables (e.g. for many to many relations).
* Service :- It will contain all the business logic of the application.
* Shared :- this will contain the shared resources that will be used by the application.

2> CMD :-  This is created while initializing the cobra. It will contain all the file related to the cobra cli commands including the root file generated by cobra by default.

3> SERVER :- It will start the server of this go service on the provided host and port.

4> perScoreProto :-  It will contain all the proto and the compiled file used by the application.

5> `setupPostgres.sh` : - this file will automatically set up the postgres in your system including the ROLE AND DATABASE in one go.

### Tables

**it will have one table  which will contain following columns** :

Item           | Value
--------       | ---
ID		       | int
Token          |string
ExpirationTime | int
Email          | string


### Basic Flow diagrams

client server interaction:

```sequence
perscore->PerScoreServer: Request
Note right of perscore: REST Calls
PerScoreServer-->perscore:Response!
```

Login and registration flow :

```flow
st=>start: start
p=>operation: perScore
s=>operation: perscoreserver
a=>operation: perscoreauth
e=>end
op=>operation: if status == "success"
cond=>condition: Yes or No?

st->p->s->a->op->cond
cond(yes)->p
cond(no)->p

```
