package domain

type (
	User struct {
		Id    int
		Name  string
		Email string
		Age   int
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

func WithAge(age int) Opt {
	return func(u *User) {
		u.Age = age
	}
}
