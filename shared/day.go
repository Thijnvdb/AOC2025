package shared

type Day struct {
	InputFile   string
	ExampleFile string
	Run         func(inputFile string, part uint) error
}
