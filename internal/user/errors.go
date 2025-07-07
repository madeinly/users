package user

type UserError struct {
	Code    string
	Message string
	Field   string
}

type UserErrors []*UserError

func NewUserChecker() *UserErrors {
	return &UserErrors{}
}

func (ues *UserErrors) AddError(code, message, field string) {
	*ues = append(*ues, &UserError{
		Code:    code,
		Message: message,
		Field:   field,
	})
}

func (ues UserErrors) HasErrors() bool {
	return len(ues) > 0
}

func (ue *UserError) GetCode() string    { return ue.Code }
func (ue *UserError) GetMessage() string { return ue.Message }
func (ue *UserError) GetField() string   { return ue.Field }
