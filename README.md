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





