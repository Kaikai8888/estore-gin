package commons

//TODO: use generic to accept different db data schema
type Repository interface {
	Find(any) []any
	FindById(string) any
	Create(any) any 
}
