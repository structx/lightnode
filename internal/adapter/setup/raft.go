package setup

// Raft
type Raft struct {
	LocalID   string `hcl:"local_id"`
	BaseDir   string `hcl:"base_dir"`
	Bootstrap bool   `hcl:"bootstrap"`
}
