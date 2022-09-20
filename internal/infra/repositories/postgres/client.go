package postgres

import "gorm.io/gorm"

type client struct {
	db *gorm.DB
}

// NewClient returns a new instance to use postgres.
func NewClient() *client {
	return &client{
		db: connect(),
	}
}

// Create stores a new record in our db.
func (c *client) Create(value interface{}) error {
	return c.db.Create(value).Error
}

// First finds first record that match with given conditions.
func (c *client) First(dest interface{}, conds ...interface{}) error {
	return c.db.First(dest, conds...).Error
}

// FindInOrder finds and orders results that match with given conditions.
func (c *client) FindInOrder(out interface{}, order string, where ...interface{}) error {
	return c.db.Order(order).Find(out, where...).Error
}
