package hashgenerator

type hashGeneratorFactory struct {
}

func Factory() *hashGeneratorFactory {
	return &hashGeneratorFactory{}
}

func (f *hashGeneratorFactory) SHA256(key string) *sha256Generator {
	return &sha256Generator{key}
}
