package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"go-auth-service/internal/domain/models"
	"go-auth-service/internal/storage"

	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	const op = "sqlite.New"

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) SaveUser(ctx context.Context, email string, passwordHash []byte) (int64, error) {
	const op = "sqlite.SaveUser"

	stmt, err := s.db.Prepare("INSERT INTO users(email, password_hash) VALUES (?, ?)")
	if err != nil {
		return -1, fmt.Errorf("%s: %w", op, err)
	}

	result, err := stmt.ExecContext(ctx, email, passwordHash)
	if err != nil {
		return -1, fmt.Errorf("%s: %w", op, storage.ErrUserExists)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func (s *Storage) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	const op = "sqlite.GetUserByEmail"

	stmt, err := s.db.Prepare("SELECT id, email, password_hash FROM users WHERE email = ?")
	if err != nil {
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	row := stmt.QueryRowContext(ctx, email)

	var user models.User
	err = row.Scan(&user.ID, &user.Email, &user.PasswordHash)
	if err != nil {
		return models.User{}, fmt.Errorf("%s: %w", op, storage.ErrUserNotFound)
	}

	return user, nil
}

func (s *Storage) GetUserById(ctx context.Context, userID int64) (models.User, error) {
	const op = "sqlite.GetUserById"

	stmt, err := s.db.Prepare("SELECT email, password_hash FROM users WHERE id = ?")
	if err != nil {
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	row := stmt.QueryRowContext(ctx, userID)

	user := models.User{ID: userID}
	err = row.Scan(&user.Email, &user.PasswordHash)
	if err != nil {
		return models.User{}, fmt.Errorf("%s: %w", op, storage.ErrUserNotFound)
	}

	return user, nil
}

func (s *Storage) IsAdmin(ctx context.Context, userID int64) (bool, error) {
	const op = "sqlite.IsAdmin"

	stmt, err := s.db.Prepare("SELECT is_admin FROM users WHERE id = ?")
	if err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}

	row := stmt.QueryRowContext(ctx, userID)

	var isAdmin bool
	err = row.Scan(&isAdmin)
	if err != nil {
		return false, fmt.Errorf("%s: %w", op, storage.ErrUserNotFound)
	}

	return isAdmin, nil
}

func (s *Storage) GetAppById(ctx context.Context, appID int) (models.App, error) {
	const op = "sqlite.GetAppById"

	stmt, err := s.db.Prepare("SELECT name, secret FROM apps WHERE id = ?")
	if err != nil {
		return models.App{}, fmt.Errorf("%s: %w", op, err)
	}

	row := stmt.QueryRowContext(ctx, appID)

	app := models.App{ID: appID}
	err = row.Scan(&app.Name, &app.Secret)
	if err != nil {
		return models.App{}, fmt.Errorf("%s: %w", op, storage.ErrUserNotFound)
	}

	return app, nil
}
