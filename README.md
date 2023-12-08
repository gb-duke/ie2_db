# Insight Engine 2.0 Database
 Insight Engine 2 API and PostgreSQL DB

 This project is still an early work in-progress and thus creates test data on every startup with a clean database.

 Data entered while in use will NOT persist between sessions.

## Project Structure
    - iac
        - aws
            - ecr
            - packer
            - rds
            - route53
            - vpc
    - src
        - consumers
        - db
        - decorators
        - dtos
        - handlers
        - interfaces
        - validators

## Running the Code

*Requirements*
- _Docker_

**STEPS**
1. Open a terminal at the root of this project.
2. Run the following command:
    - `docker compose up --build -d`
    - App startup takes about 10 seconds while the API waits for the database and message queue services to startup
    - To view output from the container logs, remove the `-d` from the command above or view them in Docker Desktop
3. To stop everything run the following command:
    - `docker compose down`

## Testing the API
With the application running, you can use a tool such as Postman to send API requests to `http://localhost:8080`

Routes supported can be viewed in the `main.go` file

DTOs can be found in `./src/dtos/..`