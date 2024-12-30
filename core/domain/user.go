package domain

type (
	User struct {
		Id    string
		Name  string
		Email string
		Age   string
	}
	Opt func(*User)
)

func NewUser(opts ...Opt) *User {
	u := &User{}
	for _, opt := range opts {
		opt(u)
	}
	return u
}

func WithName(name string) Opt {
	return func(u *User) {
		u.Name = name
	}
}

func WithEmail(email string) Opt {
	return func(u *User) {
		u.Email = email
	}
}

func WithAge(age string) Opt {
	return func(u *User) {
		u.Age = age
	}
}
