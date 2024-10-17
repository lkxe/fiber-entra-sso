# Fiber Entra-ID SSO example

This is a simple test application demonstrating Microsoft Entra ID (formerly Azure AD) Single Sign-On (SSO) using Go Fiber with Django templates and a PostgreSQL database connection using GORM.

## Features

- Microsoft Entra ID Single Sign-On
- Go Fiber web framework
- Django templating engine
- PostgreSQL database integration with GORM
- Simple note-taking functionality for authenticated users

## Prerequisites

- Go 1.23 or later (will probably also work on older versions, but haven't tested it)
- PostgreSQL
- Microsoft Entra ID account and registered application [Learn how to create an application](https://docs.ricoh-pmc.com/docs/add-a-custom-app-in-microsoft-entra-id)

## Setup

Clone the repository:
```shell
git clone https://github.com/xlukas1337/fiber-entra-sso.git
cd fiber-entra-sso
```

Install dependencies:
```shell
go mod tidy
```

Set up your PostgreSQL database:

```shellsudo -u postgres psql
CREATE USER root WITH SUPERUSER;
CREATE DATABASE notes;
GRANT ALL PRIVILEGES ON DATABASE notes TO root;
\q
```

Create a .env file in the project root with the following content or edit and rename the existing .env.example:
```
MICROSOFT_CLIENT_ID=your_client_id
MICROSOFT_TENANT_ID=your_tenant_id
MICROSOFT_CLIENT_SECRET=your_client_secret
MICROSOFT_REDIRECT_URL=http://localhost:3000/auth/microsoft/callback
PORT=3000
DATABASE_URL=postgres://root:@localhost:5432/notes
```
Replace your_client_id, your_tenant_id, and your_client_secret with your Microsoft Entra ID application credentials.

Run the application:
```shell
go run main.go
```

Open a web browser and navigate to http://localhost:3000

## Usage

Click the "Login" button to authenticate with your Microsoft account if you don't get redirected automatically
Once logged in, you can add notes using the form provided
Your notes will be displayed on the page and stored in the PostgreSQL database

## Security Note
This is a test application and is not intended for production use. It uses simplified security measures for demonstration purposes. In a production environment, you should implement proper security practices, including:

- Using a non-superuser database account with limited privileges
- Secure password storage and management
- Properly handling and validating user input
- Implementing HTTPS
- Following OWASP security guidelines
- Not use Tailwind's cdn 

## Contributing
This is a test application, but contributions for improvements or bug fixes are welcome. Please open an issue or submit a pull request.

## License
MIT License
