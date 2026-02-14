package config

/*
Role: Handles application configuration.

Responsibilities:

Loads environment variables (like SNOWFLAKE_DSN, PORT, RATE_LIMIT_RPS).

Provides defaults where necessary.

Validates critical values (e.g., Snowflake DSN must be set).

Exposes a Config struct for other packages to consume.

Why: Centralized configuration prevents scattering environment variable lookups throughout your code, making it easier to maintain and test.
*/

type Config struct {
	Logger        int
	Store         int
	RateLimiter   int
	Authenticator int
	Port          int
}
