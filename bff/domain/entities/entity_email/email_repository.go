package entity_email

type Repository interface {
	First(emailID uint64) (*Email, error)
	Create(userID uint64, emailAddr string) (*Email, error)
	Update(emailID uint64, emailAddr string) (*Email, error)
}
