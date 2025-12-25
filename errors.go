package absurd_reason

import "errors"

var (
	ErrDiceSaidNo          = errors.New("I rolled the dice and it came up 'no'.")
	ErrForduneCookieSaidNo = errors.New("The fortune cookie said 'No' just last night.")
	ErrIHateThis           = errors.New("I’ve reached an age where I just say 'no' to things I hate. And I hate this.")
)
