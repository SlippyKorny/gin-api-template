package cat

type CatService struct {
	catRepository CatRepository
}

func NewCatService() CatService {
	return CatService{
		catRepository: NewCatRepository(),
	}
}

func (cs CatService) GetAllCats() ([]CatDTO, error) {
	cats, err := cs.catRepository.GetAllCats()
	if err != nil {
		return nil, err // usually you would change this to a custom message
	}

	// TODO: Some kind of mapping for quick and easy conversion to DTO
	catDTOs := make([]CatDTO, len(cats))
	for i := 0; i < len(cats); i++ {
		catDTOs[i] = CatDTO{
			Name: cats[i].Name,
		}
	}

	return catDTOs, nil
}
