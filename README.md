## Execute the application
### Dependencies
First, install the required dependencies to run the application. You have two options for installing Goose:

**Globally (available for all projects):**
```go
go install github.com/pressly/goose/v3/cmd/goose@latest
```
**Locally (only for the current project):**
```go
go get github.com/pressly/goose/v3
```
**Before proceeding, ensure all project dependencies are downloaded and installed correctly by running:**

```go
go mod tidy
```
This will download the dependencies listed in the go.mod file and remove any unused ones.

### Database
**To run the database migrations, use the following command:**
```bash
goose -dir=src/infra/database/migrations sqlite app.db up
```
This command will apply all migrations to the app.db database using the migration files located in the src/infra/database/migrations directory.