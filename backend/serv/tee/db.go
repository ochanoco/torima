package tee

import (
	"fmt"

	"github.com/ochanoco/proxy/core"
	"github.com/ochanoco/proxy/ent"
	"github.com/ochanoco/proxy/ent/hashchain"
)

func CreateHashChain(db *core.Database, hash, signature []byte) *ent.HashChainCreate {
	chain := db.Client.HashChain.
		Create().
		SetHash(hash).
		SetSignature(signature)

	return chain
}

func SaveHashChain(db *core.Database, l *ent.HashChainCreate) (*ent.HashChain, error) {
	code, err := l.Save(db.Ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to save save log: %v", err)
	}

	return code, err
}

func FindLastHashChain(db *core.Database) (*ent.HashChain, error) {
	hash, err := db.Client.HashChain.
		Query().
		Order(ent.Desc(hashchain.FieldID)).
		Limit(1).
		All(db.Ctx)

	if err != nil {
		return nil, err
	}

	if len(hash) == 0 {
		return nil, nil
	}

	return hash[0], nil
}

func SelectAllHashChains(db *core.Database) ([]*ent.HashChain, error) {
	return db.Client.HashChain.
		Query().
		All(db.Ctx)
}

func SelectAllServiceLogs(db *core.Database) ([]*ent.ServiceLog, error) {
	return db.Client.ServiceLog.
		Query().
		All(db.Ctx)
}
