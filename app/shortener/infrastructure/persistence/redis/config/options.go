package config

type RedisOptions struct {
	Addr     *string
	Password *string
	DB       *int
}

func Config() *RedisOptions {
	return new(RedisOptions)
}

func (o *RedisOptions) SetAddress(address string) *RedisOptions {
	o.Addr = &address
	return o
}

func (o *RedisOptions) SetPassword(password string) *RedisOptions {
	o.Password = &password
	return o
}

func (o *RedisOptions) SetDB(db int) *RedisOptions {
	o.DB = &db
	return o
}

func MergeOptions(options ...*RedisOptions) *RedisOptions {
	option := new(RedisOptions)

	for _, v := range options {
		if v.DB != nil {
			option.DB = v.DB
		}
		if v.Addr != nil {
			option.Addr = v.Addr
		}
		if v.Password != nil {
			option.Password = v.Password
		}
	}
	return option
}
