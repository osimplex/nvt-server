package server

type Server interface {
	End()
	ListenAndServe() error
}
