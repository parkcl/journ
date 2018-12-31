package params

// CliParams containing properties relevant to encryption
type CliParams struct {
	SkipEncryption bool
	Key            string
}

// ShouldEncrypt returns true if skip flag is false
func (p CliParams) ShouldEncrypt() bool {
	return p.SkipEncryption == false
}
