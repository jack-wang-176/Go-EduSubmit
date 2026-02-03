package pkg

type CollectError struct {
	Raw    error
	Code   int
	Status int
	Msg    string
}
