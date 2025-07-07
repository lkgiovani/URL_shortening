package postgres

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"sort"
	"url_shortening/config/environment"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	Db  *gorm.DB
	Ctx context.Context
}

func NewPostgres(config *environment.Config) (*Postgres, error) {
	db, err := gorm.Open(postgres.Open(config.DB.DataSource), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	postgres := &Postgres{Db: db}

	if err := postgres.migrate(); err != nil {
		return nil, err
	}

	return postgres, nil
}

//go:embed migration/*.sql
var migrationFS embed.FS

func (p *Postgres) migrate() error {

	if err := p.Db.Exec(`CREATE TABLE IF NOT EXISTS migrations (name TEXT PRIMARY KEY);`).Error; err != nil {
		return fmt.Errorf("cannot create migrations table: %w", err)
	}

	// Read migration files from our embedded file system.
	// This uses Go 1.16's 'embed' package.
	names, err := fs.Glob(migrationFS, "migration/*.sql")
	if err != nil {
		return err
	}
	sort.Strings(names)

	fmt.Println(names)

	for _, name := range names {
		if err := p.migrateFile(name); err != nil {
			return fmt.Errorf("migration error: name=%q err=%w", name, err)
		}
	}

	return nil
}

func (p *Postgres) migrateFile(name string) error {

	tx := p.Db.Begin()

	defer tx.Rollback()

	var count int64
	if err := tx.Raw(`SELECT COUNT(*) FROM migrations WHERE name = ?`, name).Scan(&count).Error; err != nil {
		return err
	} else if count != 0 {
		return nil
	}

	if buf, err := migrationFS.ReadFile(name); err != nil {
		return err
	} else if err := tx.Exec(string(buf)).Error; err != nil {
		return err
	}

	if err := tx.Exec(`INSERT INTO migrations (name) VALUES (?)`, name).Error; err != nil {
		return err
	}

	return tx.Commit().Error
}
