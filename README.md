# student-registration
This project is a RESTful API for a student course registration system. This serves as the backend and a client side program will leverage this later on once this is complete. This REST API server is to be ran locally in a docker container.

Project is not fully complete yet.

## Instructions for Setup
1. Need to make sure VS Code Go extension is installed
    1. Open VS Code
    2. Go to the extensions tab 
    3. Search for ```Go``` and install it
2. Need to make sure extension REST Client is installed
    1. Open VS Code
    2. Go to the extensions tab 
    3. Search for ```Rest Client``` and install it
3. Need to install Go on your local machine
    1. Go to the Go documentation site located [here](https://go.dev/)
    2. Click on Download and select the right install package for your machine
    3. Once the install is done, restart VS Code 
    4. Open a new terminal in VS Code
    5. Verify that Go is installed by executing the following command: ```go version```
4. This is optional but recommended for convenience: install jq on your local machine. Steps are different depending on your OS. The following steps are for Windows
    1. Open VS Code as an administrator
    2. Open terminal
    3. Run the following command: ```choco install jq```
    4. Select yes for the prompts
    5. Verify that jq is installed using the following command: ```jq --version```

After completing these steps, your machine should be ready. You will need to clone this repo to your local machine by executing the following command: 
    ```
    git clone https://github.com/Jeff-Marm-Al/student-registration.git
    ```

## How to use
### Without a Docker Container
Open this folder in VS Code and run the following command to start up the database server: 
```go run .```

Once the server is running, you can go to the api-test folder and execute the four current API calls there. At the moment, there is only API endpoints to create a student, to list all students (first and last name), create a course, and list all courses.

### With a Docker Container
Pre-req: Must have docker desktop installed and have it open in the background

1. In your VS Code terminal, execute the following command to build a local docker image using the docker file
    ```
    docker build -t <name the image> .
    ```

2. Execute the following command to start up the container in detached mode with the container port 8080 mapped to the host port 8080
    ```
    docker run -d -p 8080:8080 --name <name the container> <image name from previous step>
    ```

3. Verify that the container is running by executing the following command to list the containers running
    ```
    docker ps -a
    ```

4. Verify in the logs for the container that the server is running using the container ID from the previous step's output
    ```
    docker logs <container ID>
    ```

If all is functioning correctly, you can now execute curl commands to the API server running in the container. The following are some examples of curl commands you can make to get you started off:
(if you do not have jq installed, remove "| jq" from the curl commands)

This curl would create a student
```
curl -X POST http://localhost:8080/students \
-H "Content-Type: application/json" \
-d '{"firstname": "Kyle","lastname": "Garcia","email": "kGarcia@dummy.com","password": "fakePassW0rd!"}' | jq
```

This curl would list all students
```
curl -X GET http://localhost:8080/students | jq
```

This curl would list the student created in the first example
```
curl -X GET http://localhost:8080/students/Kyle/Garcia | jq
```