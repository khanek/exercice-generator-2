package admin

import (
	"sync"

	"khanek/exercise-generator/config/bindatafs"

	"github.com/jinzhu/gorm"
	"github.com/qor/admin"
	"github.com/qor/qor"
)

type AdminModel interface {
	ConfigureResource(*admin.Resource) error
}

var _admin *admin.Admin
var _once = sync.Once{}
var _models = []AdminModel{}

func Initialize(db *gorm.DB) *admin.Admin {
	_once.Do(func() {
		_admin = admin.New(&qor.Config{DB: db})
		assetFS := bindatafs.AssetFS.NameSpace("admin")
		_admin.SetAssetFS(assetFS)
	})
	return _admin
}

func GetAdmin() *admin.Admin {
	if _admin == nil {
		panic("Admin not initialized")
	}
	return _admin
}

func Register(model AdminModel) error {
	res := _admin.AddResource(model)
	if err := model.ConfigureResource(res); err != nil {
		return err
	}
	return nil
}
