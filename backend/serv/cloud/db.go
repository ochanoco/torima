package cloud

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ochanoco/proxy/core"
	"github.com/ochanoco/proxy/ent"
	"github.com/ochanoco/proxy/ent/serviceprovider"
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

func CreateServiceProvider(db *core.Database, host string, destinationIP string) *ent.ServiceProviderCreate {
	proj := db.Client.ServiceProvider.
		Create().
		SetHost(host).
		SetDestinationIP(destinationIP)

	return proj
}

func CreateWhiteList(db *core.Database, path string) *ent.WhiteListCreate {
	wl := db.Client.WhiteList.
		Create().
		SetPath(path)

	return wl
}

func MigrateWhiteList(db *core.Database) error {
	var urls []string

	b, err := os.ReadFile(WHITELIST_PATH)
	if err != nil {
		return fmt.Errorf("failed to read white list (%v)\n=> %v", WHITELIST_PATH, err)
	}

	err = json.Unmarshal(b, &urls)

	if err != nil {
		return fmt.Errorf("failed to parse white list (%v)\n=> %v", WHITELIST_PATH, err)
	}

	projc := CreateServiceProvider(db, core.AUTH_HOST, core.PROXY_CALLBACK_URL)
	_, err = projc.Save(db.Ctx)

	if err != nil {
		return fmt.Errorf("failed creating project: %v", err)
	}

	for _, url := range urls {
		wlc := CreateWhiteList(db, url)

		_, err := wlc.Save(db.Ctx)

		if err != nil {
			return fmt.Errorf("failed add white list to project: %v", err)
		}
	}

	return nil
}

func SaveServiceProvider(db *core.Database, spc *ent.ServiceProviderCreate) (*ent.ServiceProvider, error) {
	sp, err := spc.Save(db.Ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to save service provider: %v", err)
	}

	return sp, nil
}

func SaveWhiteList(db *core.Database, projc *ent.ServiceProvider, wlc *ent.WhiteListCreate) (*ent.ServiceProvider, error) {
	wl, err := wlc.Save(db.Ctx)

	if err != nil {
		return nil, err
	}

	proj, err := projc.
		Update().
		AddWhitelists(wl).
		Save(db.Ctx)

	return proj, err
}

func SaveAuthorizationCode(db *core.Database, codeC *ent.AuthorizationCodeCreate) (*ent.AuthorizationCode, error) {
	code, err := codeC.Save(db.Ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to save authorization code: %v", err)
	}

	return code, err
}

func FindServiceProviderByHost(db *core.Database, host string) (*ent.ServiceProvider, error) {
	return db.Client.ServiceProvider.
		Query().
		Where(serviceprovider.HostEQ(host)).
		Only(db.Ctx)
}
