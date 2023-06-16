package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/RoyJoel/LZStudyingProjectBackendBackend/package/config"
	model2 "github.com/RoyJoel/LZStudyingProjectBackendBackend/package/model"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	a "github.com/casbin/xorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
)

var (
	enforcer     *casbin.Enforcer
	adapter      *a.Adapter
	TokenTimeOut = time.Second * 3600
	RDB          = config.RDB
)

func init() {
	var err error
	adapter, err = a.NewAdapter("mysql", "root:12345678@tcp(localhost:3306)/LZStudyingProjectBackend", true)
	if err != nil {
		log.Fatalf("error: adapter: %s", err)
	}
	m, err := model.NewModelFromString(`
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = my_func(r.sub,p.sub,r.obj,p.obj)
		`)
	if err != nil {
		log.Fatalf("error: model: %s", err)
	}
	//自定义匹配函数
	enforcer, err = casbin.NewEnforcer(m, adapter)
	enforcer.AddFunction("my_func", KeyMatchFunc)
	if err != nil {
		log.Fatalf("error: enforcer: %s", err)
	}
}

// 自定义匹配函数
// 规则： (r.sub == p.sub && r.obj == p.obj) || (r.sub == p.sub && p.obj == '*')
func KeyMatchFunc(args ...interface{}) (interface{}, error) {
	KeyMatch := func(rsub, psub, robj, pobj string) bool {
		return (rsub == psub && robj == pobj) || (rsub == psub && pobj == "*")
	}
	return (bool)(KeyMatch(args[0].(string), args[1].(string), args[2].(string), args[3].(string))), nil
}

func AddPolicy(role string, res string, action string) bool {
	result, err := enforcer.AddPolicy(role, res, action)
	if err != nil {
		panic(err)
	}
	return result
}

func CheckEnforce(role string, res string, action string) bool {
	result, _ := enforcer.Enforce(role, res, action)
	return result
}

func DeletePolicy(role string, res string, action string) bool {
	result, _ := enforcer.DeletePermissionForUser(role, action)
	return result
}

func SetToken(ctx context.Context, token string, player model2.Player) bool {
	RDB.Set(ctx, token, player, TokenTimeOut)
	return true
}

func GetToken(ctx context.Context, token string) model2.Player {
	var player model2.Player
	u, err := RDB.Get(ctx, token).Result()
	if len(u) == 0 {
		return player
	}
	err = json.Unmarshal([]byte(u), &player)
	if err != nil {
		panic(err)
	}
	fmt.Println(player)
	return player
}
