package checkset

// Check struct
type Check struct {
	checkset    *Checkset
	code        int
	category    string
	description string
	Result      bool
	resources   []string
}
