package cloud

import (
	"github.com/ochanoco/proxy/core"
	"github.com/ochanoco/proxy/ent"
)

func CreateAuthorizationCode(db *core.Database, token string) *ent.AuthorizationCodeCreate {
	code := db.Client.AuthorizationCode.
		Create().
		SetToken(token)

	return code
}

func CreateRandomAuthorizationCode(db *core.Database) (*ent.AuthorizationCodeCreate, error) {
	token, err := core.RandomString(32)
	if err != nil {
		return nil, err
	}

	code := db.Client.AuthorizationCode.
		Create().
		SetToken(token)

	return code, nil
}
