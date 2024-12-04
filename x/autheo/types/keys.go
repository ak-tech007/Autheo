package types

const (
	// ModuleName defines the module name
	ModuleName = "autheo"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_autheo"
)

var (
	ParamsKey = []byte("p_autheo")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
