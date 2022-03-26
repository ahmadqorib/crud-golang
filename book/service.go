package book

type Service interface {
	FindAll() ([]Book, error)
	FindById(ID int) (Book, error)
	Create(book BookRequest) (Book, error)
	Update(ID int, book BookUpdateRequest) (Book, error)
	Delete(ID int) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	books, err := s.repository.FindAll()
	return books, err
}

func (s *service) FindById(ID int) (Book, error) {
	book, err := s.repository.FindById(ID)
	return book, err
}

func (s *service) Create(request BookRequest) (Book, error) {
	price, _ := request.Price.Int64()
	book := Book{
		Title:       request.Title,
		Price:       int(price),
		Description: request.Description,
		Rating:      request.Rating,
	}

	newBook, err := s.repository.Create(book)
	return newBook, err
}

func (s *service) Update(ID int, request BookUpdateRequest) (Book, error) {
	book, err := s.repository.FindById(ID)
	price, _ := request.Price.Int64()

	book.Title = request.Title
	book.Price = int(price)
	book.Description = request.Description
	book.Rating = request.Rating

	newBook, err := s.repository.Update(book)
	return newBook, err
}

func (s *service) Delete(ID int) (Book, error) {
	book, err := s.repository.FindById(ID)
	newBook, err := s.repository.Delete(book)
	return newBook, err
}
