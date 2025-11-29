package schema

type Tool struct {
	Name        string
	Description string
	Func        func(map[string]any) (any, error)
}
