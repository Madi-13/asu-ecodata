package repositories

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"mdm-auth/internal/domain"
)

type Auth interface {
	CreatePermissions(ctx context.Context, title *string, code, role, resourceCode string, scopeCodes []string) error
	GetPermissions(ctx context.Context, roleNames []string) ([]domain.PermissionOut, error)
	CheckPermission(ctx context.Context, roleNames []string, permissionCode, scope string) (*bool, error)
}

type AuthRepo struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepo {
	return &AuthRepo{
		db: db,
	}
}

func (r *AuthRepo) CreatePermissions(ctx context.Context, title *string, code, role, resourceCode string, scopeCodes []string) error {
	var insertPermissionSql = `insert into mdm_auth.permissions(code, title, role_name, resource_code, scopes)
            VALUES ($1, $2, $3, $4, $5)`
	if _, err := r.db.ExecContext(ctx, insertPermissionSql, code, title, role, resourceCode, pq.Array(scopeCodes)); err != nil {
		return err
	}
	return nil
}

func (r *AuthRepo) GetPermissions(ctx context.Context, roleNames []string) ([]domain.PermissionOut, error) {
	var permissions = make([]domain.PermissionOut, 0)
	sql := `select r.code                                                           resource_code,
				   json_agg(json_build_object(p.code, json_build_object('role_name', p.role_name, 'title', p.title, 'scopes',
																		p.scopes))) permissions
			from mdm_auth.resources r
					 join mdm_auth.permissions p on r.code = p.resource_code
				and p.role_name = ANY ($1)
			group by r.code;`
	if err := r.db.SelectContext(ctx, &permissions, sql, pq.Array(roleNames)); err != nil {
		return nil, err
	}
	return permissions, nil
}

func (r *AuthRepo) CheckPermission(ctx context.Context, roleNames []string, permissionCode, scope string) (*bool, error) {
	var hasPermission *bool
	sql := `select exists(select *
					from mdm_auth.permissions p
					where role_name = ANY ($1)
					  and p.code = $2
					  and $3 = any (p.scopes))`
	if err := r.db.GetContext(ctx, &hasPermission, sql, pq.Array(roleNames), permissionCode, scope); err != nil {
		return nil, err
	}
	return hasPermission, nil
}
