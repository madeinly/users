package parser

import (
	"errors"
	"fmt"
	"strconv"
)

type pagination struct {
	Limit int64
	Page  int64
}

const (
	FormUserPage  = "user_page"
	FormUserLimit = "user_limit"
)

func NewPagination() pagination {
	pagination := pagination{
		Limit: -1,
		Page:  1,
	}

	return pagination
}

func (pagination *pagination) AddPage(page string) error {
	const minLen = 1

	var err error

	if len(page) < minLen {
		return errors.New(fmt.Sprintf("must be at least %d characters", minLen))
	}

	pagination.Page, err = strconv.ParseInt(page, 10, 64)

	if err != nil {
		return err
	}

	return nil
}

func (pagination *pagination) AddLimit(limit string) error {
	const minLen = 1

	var err error
	if len(limit) < minLen {
		errors.New(fmt.Sprintf("must be at least %d characters", minLen))
	}

	pagination.Limit, err = strconv.ParseInt(limit, 10, 64)

	if err != nil {
		return err
	}

	return nil
}
