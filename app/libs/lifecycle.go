package libs

import (
	"cn/com/hengwei/commons/types"

	_ "github.com/lib/pq"
	"github.com/three-plus-three/modules/web_ext"
	"github.com/three-plus-three/web_example/app/models"
)

// Lifecycle 表示一个运行周期，它包含了所有业务相关的对象
type Lifecycle struct {
	*web_ext.Lifecycle

	Definitions *types.TableDefinitions
	DB          models.DB
	DataDB      models.DB
}
