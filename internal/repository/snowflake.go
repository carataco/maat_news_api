package repository

/*
Files: snowflake_repository.go

Role: Data access layer (Snowflake-specific).

Responsibilities:

Opens and manages Snowflake connections.

Sets the session variable for row-level security (ALTER SESSION SET CLIENT_ID).

Executes queries safely with parameter binding.

Returns raw rows to service layer for processing.

Why: Encapsulates all Snowflake-specific behavior in one place.

Makes it easier to replace the data source in the future or mock it for tests.
*/
