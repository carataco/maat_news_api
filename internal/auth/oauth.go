package auth

/*
Files: middleware.go, context.go

Role: Authentication and context management.

Responsibilities:

Validates client authentication (currently a stub with X-Client-ID, later replaceable with real OAuth/JWT).

Stores client_id in the request context for downstream access.

Provides helpers to extract client_id from context (GetClientID).

Why: Decouples authentication from business logic and ensures tenant isolation by passing client_id securely through the context.
*/
