package auth

type Auth struct {
	ID         string
	FirstName  string
	LastName   string
	Gender     int
	Email      string
	Password   string
	VerifiedAt string
}

type AccessToken struct {
	AccessToken  string
	RefreshToken string
	ExpiredAt    int64
}
