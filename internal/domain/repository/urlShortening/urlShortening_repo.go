package urlShortening

import (
	"url_shortening/config/environment"
	"url_shortening/infra/db/postgres"

	"github.com/google/uuid"
)

type UrlShorteningRepository struct {
	db     *postgres.Postgres
	config *environment.Config
}

func NewUrlShorteningRepository(db *postgres.Postgres, config *environment.Config) *UrlShorteningRepository {
	return &UrlShorteningRepository{db: db, config: config}
}

func (r *UrlShorteningRepository) RegisterUrl(url *string) (string, string, error) {

	uniqueID, err := uuid.NewV7()
	if err != nil {
		return "", "", err
	}

	urlShortened := r.config.URL_SHORTENED_PREFIX + "/" + uniqueID.String()[:8]

	query := `INSERT INTO url_shortening (id,id_user,url_original,url_shortened) VALUES ($1,$2,$3,$4)`
	_, err = r.db.Db.Raw(query, uniqueID, "0197e048-f7c3-7fec-974f-6cf3a63f2383", url, urlShortened).Rows()
	if err != nil {
		return "", "", err
	}

	return urlShortened, uniqueID.String(), nil
}

func (r *UrlShorteningRepository) GetUrl(urlShortened string) (string, error) {
	query := `SELECT url_original FROM url_shortening WHERE url_shortened = $1`
	_, err := r.db.Db.Raw(query, urlShortened).Rows()
	if err != nil {
		return "", err
	}

	return "", nil
}
