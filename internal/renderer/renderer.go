package renderer

type Parser interface {
	SetExpr(string) error
	NextInstances(int) ([]string, error)
	HumanReadableStr() (string, error)
	GetHints(pad, hintIdx int) string
	Validate() error
	IncIter()
	DecIter()
}
