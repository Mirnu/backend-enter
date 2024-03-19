package store

import "database/sql"

type Store struct {
	Config         *Config
	db             *sql.DB
	userRepository *UserRepository
}

func New(config *Config) *Store {
	return &Store{
		Config: config,
	}
}

func (store *Store) Open() error {
	db, err := sql.Open("postgres", store.Config.DatabaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	store.db = db
	return nil
}
