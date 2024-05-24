package model

type BasketRetrieveError struct {
	error
}

type BasketUpdateError struct {
	error
}

type BasketDeleteError struct {
	error
}

type BasketNotFoundError struct {
	error
}

type BasketCompletedError struct {
	error
}

type BasketInvalidDataError struct {
	error
}
