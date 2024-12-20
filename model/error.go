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

type UserCreateError struct {
	error
}

type UsernamePasswordMismatchError struct {
	error
}

type UserRetreiveError struct {
	error
}

type TokenGenerationError struct {
	error
}

type UserNotFoundError struct {
	error
}

type UserDeleteError struct {
	error
}
