package ratelimit

/*
Files: middleware.go

Role: Rate limiting per client.

Responsibilities:

Implements per-client token bucket using golang.org/x/time/rate.

Applies rate limiting middleware before handlers.

Prevents abuse or accidental DoS from a single client.

Why: Protects your API and Snowflake from excessive load.

Later, this can be upgraded to a distributed limiter using Redis if horizontally scaled.
*/
