package ports

type Validator interface {
	ValidateStruct(s interface{}) error
}
