package setup

// Server configuration
type Server struct {
	ListenAddr string    `hcl:"listen_addr"`
	Telemetry  Telemetry `hcl:"telemetry,block"`
}

// Telemetry
type Telemetry struct {
	BindPort string `hcl:"bind_port"`
}
