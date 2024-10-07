package file

type ReadFileFunc func(string) ([]byte, error)
type UnmarshalFunc func([]byte, any) error

func ReadFromYAML(path string, target any, readFile ReadFileFunc, yamlUnmarshal UnmarshalFunc) error {
	yf, err := readFile(path)
	if err != nil {
		return err
	}
	return yamlUnmarshal(yf, target)
}
