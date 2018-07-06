package content

import "context"

//Service...
type service struct {
	Repository Repository
}

//NewService ...
func NewService(r Repository) Service {
	return service{Repository: r}
}

//Get ...
func (s service) Get(ctx context.Context, obj contentRequest) (contentResponse, error) {
	return s.Repository.Get(obj.ID)
}

//Get ...
func (s service) Delete(ctx context.Context, obj contentRequest) (contentResponse, error) {
	return s.Repository.Delete(obj.ID)
}

func (s service) Save(ctx context.Context, obj contentInput) (contentResponse, error) {
	return s.Repository.Save(obj.ID, obj.Body)
}
