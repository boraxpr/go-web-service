package internal

type Dao[T any] interface {
	Get(id uint32) (T, error)
	GetAll() ([]T, error)
	Save(t T) error
	Update(t T) error
	Delete(id uint32) error
}
