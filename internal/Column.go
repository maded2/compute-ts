package internal

type Column interface {
	GetName() string
	Depends() []string
	MinimumValues() int
}
