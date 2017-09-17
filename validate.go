package vld

import "bytes"

type Validate []FieldError

func (e Validate) Ok(fieldName string, err error) bool {
	if err != nil {
		e = append(e, &fieldError{
			fieldName: fieldName,
			err:       err,
		})
		return false
	}
	return true
}

func (e Validate) HasError() bool {
	return len(e) > 0
}

func (e Validate) Error() string {
	var buffer bytes.Buffer
	for _, err := range e {
		buffer.WriteString(err.Error())
		buffer.WriteString("\n")
	}
	return buffer.String()
}
