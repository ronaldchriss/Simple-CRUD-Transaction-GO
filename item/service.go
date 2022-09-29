package item

type Service interface {
	GetItems() ([]Item, error)
	GetItemByID(input GetItemDetailInput) (Item, error)
	DeleteItem(input GetItemDetailInput) (Item, error)
	CreateItem(input CreateItemInput) (Item, error)
	Update(inputID GetItemDetailInput, input CreateItemInput) (Item, error)
}

type service struct {
	reprository Reprository
}

func NewService(reprository Reprository) *service {
	return &service{reprository}
}

func (s *service) GetItems() ([]Item, error) {

	items, err := s.reprository.FindAll()
	if err != nil {
		return items, err
	}
	return items, nil

}

func (s *service) GetItemByID(input GetItemDetailInput) (Item, error) {
	items, err := s.reprository.FindByID(input.ID)
	if err != nil {
		return items, err
	}
	return items, nil
}

func (s *service) DeleteItem(input GetItemDetailInput) (Item, error) {
	items, err := s.reprository.Delete(input.ID)
	if err != nil {
		return items, err
	}
	return items, nil
}

func (s *service) CreateItem(input CreateItemInput) (Item, error) {
	items := Item{}
	items.Name = input.Name
	items.Price = input.Price
	items.Cost = input.Cost

	newInput, err := s.reprository.Save(items)
	if err != nil {
		return newInput, err
	}

	return newInput, nil
}

func (s *service) Update(inputID GetItemDetailInput, input CreateItemInput) (Item, error) {
	Item, err := s.reprository.FindByID(inputID.ID)
	if err != nil {
		return Item, err
	}

	Item.Name = input.Name
	Item.Price = input.Price
	Item.Cost = input.Cost

	updateItem, err := s.reprository.Update(Item)
	if err != nil {
		return updateItem, err
	}

	return updateItem, nil
}
