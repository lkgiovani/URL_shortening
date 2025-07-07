package urlShortening

import (
	"fmt"
	"url_shortening/config/environment"
	"url_shortening/infra/db/postgres"

	"github.com/google/uuid"
)

type UrlOriginal struct {
	UrlOriginal  string `gorm:"column:url_original"`
	UrlShortened string `gorm:"column:url_shortened"`
	Slug         string `gorm:"column:slug"`
}

type UrlShorteningRepository struct {
	db     *postgres.Postgres
	config *environment.Config
}

func NewUrlShorteningRepository(db *postgres.Postgres, config *environment.Config) *UrlShorteningRepository {
	return &UrlShorteningRepository{db: db, config: config}
}

func (r *UrlShorteningRepository) RegisterUrl(url *string) (UrlOriginal, error) {

	uniqueID, err := uuid.NewV7()
	if err != nil {
		return UrlOriginal{}, err
	}

	urlShortened := r.config.URL_SHORTENED_PREFIX + "/" + uniqueID.String()[len(uniqueID.String())-8:]

	query := `SELECT url_original, url_shortened, slug FROM url_shortening WHERE id_user = $1 AND url_original = $2`
	response, err := r.db.Db.Raw(query, "0197e048-f7c3-7fec-974f-6cf3a63f2383", *url).Rows()
	if err == nil {
		var urlOriginal UrlOriginal

		for response.Next() {
			err = response.Scan(&urlOriginal.UrlOriginal, &urlOriginal.UrlShortened, &urlOriginal.Slug)
			if err != nil {
				return UrlOriginal{}, err
			}
		}

		defer response.Close()

		return urlOriginal, nil
	}

	query = `INSERT INTO url_shortening (id,id_user,url_original,url_shortened, slug) VALUES ($1,$2,$3,$4,$5)`
	response, err = r.db.Db.Raw(query, uniqueID, "0197e048-f7c3-7fec-974f-6cf3a63f2383", url, urlShortened, uniqueID.String()[len(uniqueID.String())-8:]).Rows()
	if err != nil {
		return UrlOriginal{}, err
	}

	defer response.Close()

	return UrlOriginal{
		UrlOriginal:  *url,
		UrlShortened: urlShortened,
		Slug:         uniqueID.String()[len(uniqueID.String())-8:],
	}, nil
}

func (r *UrlShorteningRepository) GetUrl(urlShortened string) (UrlOriginal, error) {

	fmt.Println(urlShortened)
	query := `SELECT url_original, url_shortened, slug FROM url_shortening WHERE url_shortened = $1 LIMIT 1`
	response, err := r.db.Db.Raw(query, urlShortened).Rows()
	if err != nil {
		return UrlOriginal{}, err
	}

	var urlOriginal UrlOriginal

	for response.Next() {
		err = response.Scan(&urlOriginal.UrlOriginal, &urlOriginal.UrlShortened)
		if err != nil {
			return UrlOriginal{}, err
		}
	}

	defer response.Close()

	return urlOriginal, nil
}
