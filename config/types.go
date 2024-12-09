package config

// Represents the server configuration present in the config.toml file.
type Server struct {
	MaxWorkers int `toml:"max-workers"`
	Motd       string
}

// Represents a basic worker configuration.
type Worker struct {
	Listen  string
	Forward string
}

// Represents the forwarder backend.
type Forwarder struct {
	Addrs  string
	Weight int
}

// Represents the worker configuration with an extra load-balancing config.
type WorkerWithLoadBalancing struct {

	// Determines which algorith will be used with the load-balancer.
	Algorithm string

	// A list of backend forwarders.
	Forwarders []Forwarder
}

// Represents a list of workers.
type Workers = []Worker

// Represents the 'config.toml' file.
type Configuration struct {
	Server
	Workers
}
