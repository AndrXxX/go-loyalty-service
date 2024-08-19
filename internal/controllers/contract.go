package controllers

type orderNumberChecker interface {
	Check(val string) bool
}

type ormConverter[source any, target any] interface {
	ConvertMany(list []source) []target
}
