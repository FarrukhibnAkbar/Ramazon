package constants

type Sentinel string

func (s Sentinel) Error() string {
	return string(s)
}

const (
	ErrPasswordTooShort      = Sentinel("password is too short")
	ErrPasswordTooLong       = Sentinel("password is too long")
	ErrMustContainDigit      = Sentinel("password must contain at least 1 digit")
	ErrMustContainAlphabetic = Sentinel("password must contain at least 1 alphabetic")
	ErrEmailAddress          = Sentinel("invalid email address. valid e-mail can contain only latin letters, numbers, '@' and '.'")
	ErrInvalidRequestBody    = Sentinel("invalid request body")
	ErrIncorrectNameValue    = Sentinel("name must include minumum 3 and maximum 255 charachters")
	ErrAuthIncorrect         = Sentinel("auth incorrect")
	ErrAuthNotGiven          = Sentinel("auth not given")
)

const (
	// PGForeignKeyViolationCode is used to check foriegn key violation in database
	PGForeignKeyViolationCode = "23503"
	// PGUniqueKeyViolationCode is used to check unique key violation in database
	PGUniqueKeyViolationCode = "23505"
)

const (
	// ErrRowsAffectedIsZero indicates that sql command didn't work
	ErrRowsAffectedIsZero = Sentinel("no rows affected after sql command")
	//ErrAlreadyExists ...
	ErrAlreadyExists = Sentinel("already exists")
	//ErrNotExists ...
	ErrDistrictNotExists = Sentinel("district no exists")
)
