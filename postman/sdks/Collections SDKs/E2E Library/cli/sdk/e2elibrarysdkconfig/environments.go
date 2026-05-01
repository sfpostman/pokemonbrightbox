package e2elibrarysdkconfig

// Environment type defines the available API environments.
type Environment string

// Environment constants define the available base URLs for different deployment environments.
// Use these constants when configuring the SDK client.
const (
	DefaultEnvironment    Environment = "https://library-api.postmanlabs.com"
	LibraryApiEnvironment Environment = "https://library-api.postmanlabs.com"
)
