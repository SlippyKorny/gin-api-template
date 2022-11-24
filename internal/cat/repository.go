package cat

import "github.com/SlippyKorny/gin-api-template/internal/entity"

type CatRepository struct {
}

func NewCatRepository() CatRepository {
	return CatRepository{}
}

func (cr CatRepository) GetAllCats() ([]entity.Cat, error) {
	return []entity.Cat{
		{ID: 0, Name: "Meowster"},
	}, nil
}
