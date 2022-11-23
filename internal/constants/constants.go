package constants

const (
	EnvConfigPath     = "CONFIG"
	DefaultConfigPath = "/Users/rayyanfaisalhunerkar/workspace/scrummy/scrummyy-api/configs"
	DatabaseUsername  = "database.username"
	DatabasePassword  = "database.password"
	DatabaseName      = "database.name"
	DatabasePort      = "database.port"
)

const (
	KeyHttpClientInstance   = "http_client"
	KeyDecisionDBInstance   = "decision_database"
	KeyCacheInstance        = "cache"
	HeaderRequestID         = "X-Request-Id"
	HeaderRequestIDResponse = "Request-Id"
	IsCorsEnabled           = "cors.enabled"
	CorsAllowedHeaders      = "cors.allowed_headers"
	CorsAllowedMethods      = "cors.allowed_methods"
	CorsAllowAllOrigins     = "cors.allow_all_origins"
	AppHost                 = "app.host"
	AppPort                 = "app.port"
	AppEnvironment          = "app.environment"
	DebugLevel              = "debug.level"
	DebugSQLQueries         = "debug.sql_queries"
	AppCookieDomain         = "app.cookie_domain"
	AppRegion               = "app.region"
	ApiReadTimeout          = "timeouts.api_read_timeout"
	ApiWriteTimeout         = "timeouts.api_write_timeout"
)

const (
	AuthenticationEnabled      = "authentication.enabled"
	AuthenticationType         = "authentication.type"
	AuthenticationExcludePaths = "authentication.exclude_paths"
	AuthorizationType          = "authorization.type"
	AuthorizationExcludePaths  = "authorization.exclude_paths"
)
