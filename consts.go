package parcels


const (
	lenBuffer = 64 // We use 64 to match the size of a cache line,
	               // with the hope that this makes improves the
	               // performance .
)
