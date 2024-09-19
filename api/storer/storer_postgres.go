package storer

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type PostgesStorer struct {
	db *sqlx.DB
}

func NewPostgresStorer(db *sqlx.DB) *PostgesStorer {
	return &PostgesStorer{
		db: db,
	}
}

func (ps *PostgesStorer) CreateProduct(ctx context.Context, p *Product) (*Product, error) {
	query := "INSERT INTO products (name, image, category, description, rating, num_reviews, price, count_in_stock) VALUES (:name, :image, :category, :description, :rating, :num_reviews, :price, :count_in_stock) RETURNING id"

	// Prepare the query with named parameters
	stmt, err := ps.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error preparing query: %w", err)
	}
	defer stmt.Close()

	// Execute the query and return the ID
	err = stmt.GetContext(ctx, &p.ID, p)
	if err != nil {
		return nil, fmt.Errorf("error creating product: %w", err)
	}

	return p, nil
}

func (ps *PostgesStorer) GetProduct(ctx context.Context, id int64) (*Product, error) {
	var p Product
	query := "SELECT * FROM products WHERE id=?"
	err := ps.db.SelectContext(ctx, p, query, id)
	if err != nil {
		return nil, fmt.Errorf("error listing products: %w", err)
	}

	return &p, nil
}

func (ps *PostgesStorer) ListProducts(ctx context.Context) ([]Product, error) {
	var products []Product
	err := ps.db.SelectContext(ctx, &products, "SELECT * FROM products")
	if err != nil {
		return nil, fmt.Errorf("error listing products: %w", err)
	}

	return products, nil
}

func (ps *PostgesStorer) UpdateProduct(ctx context.Context, p *Product) (*Product, error) {
	query := "UPDATE products SET name=:name, image=:image, category=:category, description=:description, rating=:rating, num_reviews=:num_reviews, price=:price, count_in_stock=:count_in_stock, updated_at=:updated_at WHERE id=:id"
	_, err := ps.db.NamedExecContext(ctx, query, p)
	if err != nil {
		return nil, fmt.Errorf("error updating product: %w", err)
	}

	return p, nil
}

func (ps *PostgesStorer) DeleteProduct(ctx context.Context, id int64) error {
	query := "DELETE FROM products WHERE id=?"
	_, err := ps.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("error deleting product: %w", err)
	}

	return nil
}
