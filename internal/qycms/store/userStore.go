package store

import (
	"context"
	metav1 "gitee.com/windcoder/qingyucms/internal/pkg/qycms-common/meta/v1"
	v1 "gitee.com/windcoder/qingyucms/internal/qycms/models/v1"
)

type UserStore interface {
	Create(ctx context.Context, user *v1.User, opts metav1.CreateOptions) error
	Update(ctx context.Context, user *v1.User, opts metav1.UpdateOptions) error
	Delete(ctx context.Context, username string, opts metav1.DeleteOptions) error
	// DeleteByWhere(ctx context.Context, whereStr string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, usernames []string, opts metav1.DeleteOptions) error
	Get(ctx context.Context, username string, opts metav1.GetOptions) (*v1.User, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.UserList, error)
	ListOptional(ctx context.Context, opts metav1.ListOptions) (*v1.UserList, error)
}
