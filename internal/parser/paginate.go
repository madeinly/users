package parser

import (
	"fmt"
	"strconv"
)

const (
	FormUserPage  = "user_page"
	FormUserLimit = "user_limit"
)

func (v *UserParseErrors) ValidateUserPage(page string) int64 {
	const minLen = 1

	if len(page) < minLen {
		v.AddError(FormUserPage, fmt.Sprintf("must be at least %d characters", minLen))
	}

	pageInt, err := strconv.ParseInt(page, 10, 64)

	if err != nil {
		v.AddError(FormUserPage, err.Error())
		return 1
	}

	return pageInt
}

func (v *UserParseErrors) ValidateUserLimit(limitSTR string) int64 {
	const minLen = 1

	if len(limitSTR) < minLen {
		v.AddError(FormUserLimit, fmt.Sprintf("must be at least %d characters", minLen))
	}

	limit, err := strconv.ParseInt(limitSTR, 10, 64)

	if err != nil {
		v.AddError(FormUserLimit, err.Error())
	}

	return limit
}
