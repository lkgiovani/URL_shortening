package urlShortening_repo

import (
	"fmt"
	"url_shortening/infra/config/environment"
	"url_shortening/infra/db/postgres"

	"github.com/google/uuid"
)

type UrlOriginal struct {
	UrlOriginal  string `gorm:"column:url_original"`
	UrlShortened string `gorm:"column:url_shortened"`
	Slug         string `gorm:"column:slug"`
}

type UrlListItem struct {
	ID           string `gorm:"column:id"`
	UrlOriginal  string `gorm:"column:url_original"`
	UrlShortened string `gorm:"column:url_shortened"`
	Slug         string `gorm:"column:slug"`
	CreatedAt    string `gorm:"column:created_at"`
}

type UrlShorteningRepository struct {
	db     *postgres.Postgres
	config *environment.Config
}

func NewUrlShorteningRepository(db *postgres.Postgres, config *environment.Config) *UrlShorteningRepository {
	return &UrlShorteningRepository{db: db, config: config}
}

func (r *UrlShorteningRepository) RegisterUrl(url *string, idUser string) (UrlOriginal, error) {

	uniqueID, err := uuid.NewV7()
	if err != nil {
		return UrlOriginal{}, err
	}

	urlShortened := r.config.URL_SHORTENED_PREFIX + "/" + uniqueID.String()[len(uniqueID.String())-8:]

	query := `SELECT url_original, url_shortened, slug FROM url_shortening WHERE id_user = $1 AND url_original = $2`
	response, err := r.db.Db.Raw(query, idUser, *url).Rows()
	if err != nil {
		return UrlOriginal{}, err
	}
	defer response.Close()

	var urlOriginal UrlOriginal

	fmt.Println(response)
	if response.Next() {
		err = response.Scan(&urlOriginal.UrlOriginal, &urlOriginal.UrlShortened, &urlOriginal.Slug)
		if err != nil {
			return UrlOriginal{}, err
		}

		fmt.Println(urlOriginal)

		return urlOriginal, nil
	}

	query = `INSERT INTO url_shortening (id,id_user,url_original,url_shortened, slug) VALUES ($1,$2,$3,$4,$5)`
	response, err = r.db.Db.Raw(query, uniqueID, idUser, url, urlShortened, uniqueID.String()[len(uniqueID.String())-8:]).Rows()
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

	query := `SELECT url_original, url_shortened, slug FROM url_shortening WHERE slug = $1 LIMIT 1`
	response, err := r.db.Db.Raw(query, urlShortened).Rows()
	if err != nil {
		return UrlOriginal{}, err
	}

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

func (r *UrlShorteningRepository) GetUserUrls(idUser string) ([]UrlListItem, error) {
	query := `SELECT id, url_original, url_shortened, slug, created_at FROM url_shortening WHERE id_user = $1 ORDER BY created_at DESC`

	rows, err := r.db.Db.Raw(query, idUser).Rows()
	if err != nil {
		return []UrlListItem{}, err
	}
	defer rows.Close()

	var urls []UrlListItem
	for rows.Next() {
		var url UrlListItem
		err = rows.Scan(&url.ID, &url.UrlOriginal, &url.UrlShortened, &url.Slug, &url.CreatedAt)
		if err != nil {
			return []UrlListItem{}, err
		}
		urls = append(urls, url)
	}

	// Garantir que sempre retorne um array, mesmo que vazio
	if urls == nil {
		urls = []UrlListItem{}
	}

	return urls, nil
}
