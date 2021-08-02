package internal

type Column interface {
	Depends() []string
	MinimumValues() int
}
