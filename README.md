## my-app-go

This is a simple CRUD application written in Go, using a MySQL database. The web interface allows you to manage the following entities::

- Teachers
- Students
- Courses
- StudentsCourses (records linking students to their courses)

## Dependencies

- Go 1.24.6 or newer
- godoc 0.1.0

## Features

Each module (teachers, students, etc.) supports:

- Viewing a list of records
- Creating a new record
- Editing a record (except for StudentsCourses)
- Deleting a record

## Templates

All HTML templates are stored in the `templates/` folder. File names follow the format:

`<module>_<action>.html`

Examples:

- `teachers_index.html`
- `students_create.html`
- `courses_edit.html`
- `students_courses_index.html`

## Running the Application

The application can be run in two modes:

#### Development:

    ENVIRONMENT=dev ./my-app-go

#### Production:

    ENVIRONMENT=prod ./my-app-go

Depending on the selected mode, the application will use the corresponding database configuration and port.

## Example Usage

#### 1. Start the application

    ENVIRONMENT=dev ./my-app-go

#### 2. Open the web interface

Open your browser and go to:

`http://localhost:8080`

#### 3. Database migrations

Flyway will create the required tables and populate them with initial data:

    docker compose -f docker-compose.yml -f docker-compose.dev.yml --env-file .env.dev --profile migrations run --rm flyway migrate