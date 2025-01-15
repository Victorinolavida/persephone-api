package auth

import "github.com/uptrace/bun"

type AuthRepo struct {
	DB *bun.DB
}
