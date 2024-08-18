package interfaces

type HashGenerator interface {
	Generate(data []byte) string
}
