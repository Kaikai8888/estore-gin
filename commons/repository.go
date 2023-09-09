package commons

//TODO: use generic to accept different db data schema
type Repository interface {
	Find(any) ([]any, error)
	FindById(string) (any, error)
	Create(any) (error)
}
