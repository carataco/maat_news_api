package query

/*
Files: builder.go, filters.go, pagination.go

Role: Query construction layer.

Responsibilities:

Validates requested columns against an allowlist.

Builds safe SQL queries (column selection, date filtering, pagination).

Encapsulates query logic separately from repository/execution.

Why: Prevents SQL injection and keeps SQL generation logic centralized and testable.
*/
