# Get correct version

brew services stop postgresql@14 // or any
brew unlink postgresql@14
brew uninstall postgresql@14

---

brew services restart postgresql@16
brew link --overwrite --force postgresql@16
brew services list

# Connect to Postgres

createdb x_clone_development

pgcli x_clone_development

CREATE ROLE xuser WITH LOGIN PASSWORD 'xpassword';
GRANT ALL PRIVILEGES ON DATABASE x_clone_development TO xuser;
REVOKE ALL PRIVILEGES ON DATABASE x_clone_development FROM insidious;
GRANT ALL PRIVILEGES ON DATABASE x_clone_development TO xuser;
ALTER DATABASE x_clone_development OWNER TO xuser;


# Generate GraphQL 
* Initialize Model in resolver.go
* add corresponding functions in their own resolvers

Default resolvers are Query and Mutation, and if you have anything extra, then add it inside `gqlgen.yml` for your new domain to resolve

* Run:
    go get github.com/vektah/gqlparser/v2/validator@v2.5.11
    go get github.com/99designs/gqlgen@v0.17.45
    go run github.com/99designs/gqlgen



# Playground


`register`

```
mutation{
  register(
    input: {
      email: "bob2@gmail.com",
      username: "bob2",
      password: "password",
      confirmPassword: "password",
    }
  ) {
    accessToken
    user{
      id
      email
      username
      createdAt
    }
  }
}
```

`login`

```
mutation{
  login(
    input: {
      email: "bob2@gmail.com",
      password: "password",
    }
  ) {
    accessToken
    user{
      id
      email
      username
      createdAt
    }
  }
}
```
