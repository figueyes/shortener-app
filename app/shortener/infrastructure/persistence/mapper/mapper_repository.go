package mapper

import "github.com/figueyes/shortener-app/app/shortener/domain/entities"

type mapperRepository struct {
	shorts []*entities.Short
}

func NewMapperRepository() *mapperRepository {
	return &mapperRepository{
		shorts: make([]*entities.Short, 0),
	}
}

func (r *mapperRepository) Create(short *entities.Short) (*entities.Short, error) {
	r.shorts = append(r.shorts, short)
	return short, nil
}

func (r *mapperRepository) FindByShortUrl(shortUrl string) (*entities.Short, error) {
	for _, value := range r.shorts {
		if value.ShortUrl == shortUrl {
			return value, nil
		}
	}
	return nil, nil
}

func (r *mapperRepository) Update(shortUrl string, short *entities.Short) (*entities.Short, error) {
	for _, value := range r.shorts {
		if value.ShortUrl == shortUrl {
			if short.ShortUrl != "" {
				value.ShortUrl = short.ShortUrl
			}
			if short.OriginalUrl != "" {
				value.OriginalUrl = short.OriginalUrl
			}
			if short.User != "" {
				value.User = short.User
			}
			if short.IsEnable != nil {
				value.IsEnable = short.IsEnable
			}
			return value, nil
		}
	}
	return nil, nil
}
