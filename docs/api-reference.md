# API Reference

The `api` package provides the primary interface for interacting with ESENT.

## Core Components

### Instance (`JET_INSTANCE`)
The instance is the top-level object in ESENT. Each instance operates as a separate process environment.
- `JetCreateInstance2`: Creates a new instance.
- `JetInit2`: Initializes the instance (recovers the database).
- `JetTerm2`: Terminates the instance.

### Session (`JET_SESID`)
A session is the context for all database operations. Transactions are scoped to a session.
- `JetBeginSession`: Starts a session.
- `JetEndSession`: Ends a session.

### Database (`JET_DBID`)
- `JetCreateDatabase`: Creates a new database file.
- `JetAttachDatabase`: Attaches an existing database file to the instance.
- `JetOpenDatabase`: Opens an attached database for use.

### Table (`JET_TABLEID`)
- `JetCreateTable`: Creates a new table.
- `JetOpenTable`: Opens an existing table.
- `JetCloseTable`: Closes a table.

## Data Operations

### Columns
- `JetAddColumn`: Adds a column definition to a table.
- `JetRetrieveColumn`: Reads data from a column for the current record.
- `JetSetColumn`: Writes data to a column for the current record.

### Records
- `JetUpdate`: Inserts or updates a record.
- `JetDelete`: Deletes the current record.
- `JetMove`: Moves the cursor (First, Last, Next, Previous).
- `JetSeek`: Seeks to a record matching the current index key.

## Transactions

ESENT supports ACID transactions.
- `JetBeginTransaction`: Starts a transaction.
- `JetCommitTransaction`: Commits changes.
- `JetRollback`: Rolls back changes.

## System Parameters
You can configure the engine using system parameters via `JetSetSystemParameter`. Common parameters include:
- Database page size
- Log file path
- Circular logging options

For a full list of functions, refer to the `IJetApi` interface in `api/interface.go`.

