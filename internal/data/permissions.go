package data

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

type Permissions []string

func (p Permissions) Include(code string) bool {
	for i := range p {
		if code == p[i] {
			return true
		}
	}
	return false
}

type PermissionModel struct {
	DB *sql.DB
}

func (m PermissionModel) GetAllForUser(ctx context.Context, userID int64) (Permissions, error) {
	query := `
		SELECT p.code
		FROM permissions p
		INNER JOIN users_permissions up ON up.permission_id = p.id
		INNER JOIN users u ON u.id = up.user_id
		WHERE u.id = $1`

	rows, err := m.DB.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permissions Permissions

	for rows.Next() {
		var permission string

		err := rows.Scan(&permission)
		if err != nil {
			return nil, err
		}

		permissions = append(permissions, permission)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return permissions, nil
}

func (m PermissionModel) AddForUser(ctx context.Context, userID int64, codes ...string) error {
	query := `
		INSERT INTO users_permissions
		SELECT $1, permissions.id FROM permissions WHERE permissions.code = ANY($2)`

	_, err := m.DB.ExecContext(ctx, query, pq.Array(codes))

	return err
}
