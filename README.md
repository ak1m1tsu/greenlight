## Setup database

Create database

```sql
CREATE DATABASE <name>
```

Create extensions

```sql
# For version fields
CREATE EXTENSION uuid-ossp;

# For email fields
CREATE EXTENSION citext;

# For indexes
CREATE EXTENSION pg_trgm;
```
