package models

import (
	"database/sql"
	"errors"

	"github.com/brunoleitem/gr/data"
)

type Config struct {
	Key string
}

func (c *Config) Restore() (*Config, error) {
	var result Config
	query := `SELECT * FROM config LIMIT 1`
	err := data.DB.QueryRow(query).Scan(&result.Key)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Config) Save() error {
	exists, err := c.exists()
	if err != nil {
		return err
	}
	if exists {
		return c.update()
	}
	return c.insert()
}

func (c *Config) update() error {
	query := `UPDATE config SET key = ?`
	stmt, err := data.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(c.Key)
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) insert() error {
	query := `INSERT INTO config(key)
	VALUES (?)`
	stmt, err := data.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(c.Key)

	if err != nil {
		return err
	}
	return nil
}

func (c *Config) exists() (bool, error) {
	var key string
	query := `SELECT key FROM config LIMIT 1`
	err := data.DB.QueryRow(query).Scan(&key)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
