# student-registration
Project is not done. Work in progress
## Instructions for Setup
1. Need to make sure VS Code Go extension is installed
    1. Open VS Code
    2. Go to the extensions tab 
    3. Search for 'Go' and install it
2. Need to make sure extension REST Client is installed
    1. Open VS Code
    2. Go to the extensions tab 
    3. Search for 'Rest Client' and install it
3. Need to install Go on your local machine
    1. Go to the Go documentation site located [here](https://go.dev/)
    2. Click on Download and select the right install package for your machine
    3. Once the install is done, restart VS Code 
    4. Open a new terminal in VS Code
    5. Verify that Go is installed by executing the following command: ```go version```

After completing these steps, your machine should be ready. You will need to clone this repo to your local machine by executing the following command: 
    ```
    git clone https://github.com/Jeff-Marm-Al/student-registration.git
    ```

## How to use
Open this folder in VS Code and run the following command to start up the database server: ```go run .```

Once the server is running, you can go to the api-test folder and execute the four current API calls there. At the moment, there is only API endpoints to create a student, to list all students (first and last name), create a course, and list all courses.