package controllers

type orderNumberChecker interface {
	Check(val string) bool
}
