package productrepository_test

import (
	"math"
	productrepository "project_pos_app/repository/product"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to create GORM database: %v", err)
	}

	return gormDB, mock
}

func TestShowAllProducts(t *testing.T) {
	db, mock := setupTestDB(t)
	sqlDB, err := db.DB()
	if err != nil {
		t.Fatalf("Failed to get raw DB: %v", err)
	}
	defer func() { _ = sqlDB.Close() }()

	logger := zap.NewNop()
	repo := productrepository.NewProductRepo(db, logger)

	// Test with page = 1
	page, limit := 1, 10
	expectedTotalRecords := 30
	expectedTotalPages := int(math.Ceil(float64(expectedTotalRecords) / float64(limit)))

	mock.ExpectQuery(`SELECT count\(\*\) FROM "products"`).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(expectedTotalRecords))

	// Expect query without OFFSET when page = 1
	mock.ExpectQuery(`SELECT \* FROM "products" LIMIT \$1`).
		WithArgs(limit).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "price"}).
			AddRow(1, "Product 1", 100).
			AddRow(2, "Product 2", 200))

	products, totalRecords, totalPages, err := repo.ShowAllProducts(page, limit)

	assert.NoError(t, err)
	assert.Equal(t, expectedTotalRecords, totalRecords)
	assert.Equal(t, expectedTotalPages, totalPages)
	assert.Len(t, *products, 2)
	assert.Equal(t, "Product 1", (*products)[0].Name)
	assert.Equal(t, "Product 2", (*products)[1].Name)

	// Test with page > 1
	page = 2
	mock.ExpectQuery(`SELECT count\(\*\) FROM "products"`).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(expectedTotalRecords))

	// Expect query with OFFSET when page > 1
	mock.ExpectQuery(`SELECT \* FROM "products" LIMIT \$1 OFFSET \$2`).
		WithArgs(limit, (page-1)*limit).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "price"}).
			AddRow(3, "Product 3", 300).
			AddRow(4, "Product 4", 400))

	products, totalRecords, totalPages, err = repo.ShowAllProducts(page, limit)

	assert.NoError(t, err)
	assert.Equal(t, expectedTotalRecords, totalRecords)
	assert.Equal(t, expectedTotalPages, totalPages)
	assert.Len(t, *products, 2)
	assert.Equal(t, "Product 3", (*products)[0].Name)
	assert.Equal(t, "Product 4", (*products)[1].Name)
}

// func TestCreateProduct(t *testing.T) {
// 	db, mock := setupTestDB(t)
// 	defer func() { _ = db }()

// 	logger := zap.NewNop()
// 	repo := productrepository.NewProductRepo(db, logger)

// 	// Contoh data produk dengan kolom lengkap
// 	product := &model.Product{
// 		Name:       "New Product",
// 		Price:      300,
// 		ImageURL:   "http://example.com/product.jpg",
// 		ItemID:     "12345",
// 		Stock:      "10",
// 		CategoryID: 1,
// 		Qty:        5,
// 		Status:     "active",
// 	}

// 	// Expectation pada mock
// 	mock.ExpectBegin()
// 	mock.ExpectExec(`INSERT INTO "products"
// 		\("name", "price", "image_url", "item_id", "stock", "category_id", "qty", "status", "created_at", "updated_at", "deleted_at"\)
// 		VALUES
// 		\(\$1, \$2, \$3, \$4, \$5, \$6, \$7, \$8, \$9, \$10, \$11\)`).
// 		WithArgs(product.Name, product.Price, product.ImageURL, product.ItemID, product.Stock, product.CategoryID, product.Qty, product.Status, product.CreatedAt, product.UpdatedAt, product.DeletedAt).
// 		WillReturnResult(sqlmock.NewResult(1, 1))
// 	mock.ExpectCommit()

// 	// Call to the actual function
// 	err := repo.CreateProduct(product)

// 	// Check for errors
// 	assert.NoError(t, err)
// }

// func TestUpdateProduct(t *testing.T) {
// 	db, mock := setupTestDB(t)
// 	defer func() { _ = db }()

// 	logger := zap.NewNop()
// 	repo := productrepository.NewProductRepo(db, logger)

// 	product := &model.Product{Name: "Updated Product", Price: 500}

// 	mock.ExpectExec("UPDATE \"products\" SET \"name\"=\\$1, \"price\"=\\$2 WHERE id = \\$3").
// 		WithArgs(product.Name, product.Price, 1).
// 		WillReturnResult(sqlmock.NewResult(0, 1))

// 	err := repo.UpdateProduct(1, product)

// 	assert.NoError(t, err)
// }

// func TestDeleteProduct(t *testing.T) {
// 	db, mock := setupTestDB(t)
// 	defer func() { _ = db }()

// 	logger := zap.NewNop()
// 	repo := productrepository.NewProductRepo(db, logger)

// 	mock.ExpectExec("DELETE FROM \"products\" WHERE \"products\"\\.\"id\" = \\$1").
// 		WithArgs(1).
// 		WillReturnResult(sqlmock.NewResult(0, 1))

// 	err := repo.DeleteProduct(1)

// 	assert.NoError(t, err)
// }
