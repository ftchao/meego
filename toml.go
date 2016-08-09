package meego

/**
 * 读取配置文件
 * 配置文件使用TOML规范
 *
 * example:
 * 1.
 *    meego.Loadconf()
 * 2.
 *    import "github.com/ftchao/meego/toml"
 *
 *    // 读取Int类型
 *    r_auth_connect_timeout := meego.ConfGet("redis.auth.name")
 *
 *    // 读取String类型
 *    r_auth_connect_timeout := meego.ConfGet("redis.auth.connect.timeout")
 *
 *    // 读取数组类型
 *    r_auth_servers := meegol.ConfGet("redis.auth.server").([]interface{})
 */

import (
	"github.com/pelletier/go-toml"
	"log"
)

var _instance_toml_conf *toml.TomlTree

func LoadConf(path string) {
	var err error
	if _instance_toml_conf == nil {
		_instance_toml_conf, err = toml.LoadFile(path)
		if err != nil {
			log.Fatal("LoadConf: ", err)
		}
	}
}

func ConfGet(key string) interface{} {
	if _instance_toml_conf != nil {
		return _instance_toml_conf.Get(key)
	}
	return nil
}

func ConfGetDefault(key string, def interface{}) interface{} {
	if _instance_toml_conf != nil {
		return _instance_toml_conf.GetDefault(key, def)
	}

	return nil
}
