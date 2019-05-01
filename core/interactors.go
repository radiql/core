package core

// Interactor is the interface to be implemented by Interactor components
// from Robert C. Martin's Clean Architecture.
type Interactor interface {
	Process()
}
