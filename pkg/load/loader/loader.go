package loader

type Loader interface {
	Load(string) (string, error)
	GetRPM() int
}
