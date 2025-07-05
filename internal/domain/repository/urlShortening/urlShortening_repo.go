package urlShortening

import (
	"fmt"
	"url_shortening/infra/db/postgres"
)

type UrlShorteningRepository struct {
	db *postgres.Postgres
}

func NewUrlShorteningRepository(db *postgres.Postgres) *UrlShorteningRepository {
	return &UrlShorteningRepository{db: db}
}

func (r *UrlShorteningRepository) RegisterUrl(url *string) error {

	query := `INSERT INTO url_shortening (url_original) VALUES ($1)`
	_, err := r.db.Db.Raw(query, url).Rows()
	if err != nil {
		return err
	}

	fmt.Println(url)

	return nil
}
