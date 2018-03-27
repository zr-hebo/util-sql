# cmtrm
sql statement comment remover

### 使用示例如下：
将下面代码保存在main.go文件

执行命令`>go run main.go`
```
package main

import (
	"fmt"

	"github.com/zr-hebo/sqlutil"
)

func main() {
	stmts := []string{
		"/* asdfasfdsdf */delete from abc",
		`-- ----------------------------
-- Table structure for osc
-- ----------------------------`,
		"-- ----------------------------\n -- Table structure for osc\n -- ----------------------------\n DROP TABLE IF EXISTS `osc`; CREATE TABLE `osc` ( `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id', `appename` varchar(32) NOT NULL COMMENT '应用英文名', `keyspace` varchar(32) DEFAULT NULL COMMENT '逻辑库', `tablename` varchar(128) DEFAULT NULL COMMENT '表名', `shard` varchar(16) DEFAULT NULL COMMENT '分片', `host` varchar(64) DEFAULT NULL COMMENT 'host', `alter` varchar(256) DEFAULT NULL COMMENT 'alter语句', `status` int(11) DEFAULT NULL COMMENT '状态', `message` text COMMENT '信息', `createtime` timestamp NOT NULL DEFAULT '2017-01-01 00:00:00' COMMENT '创建时间', `updatetime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间', PRIMARY KEY (`id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8;\n-- ----------------------------\n-- Records of osc\n-- ----------------------------",
	}
	for _, stmt := range stmts {
		fmt.Printf("sql:%v\n", stmt)
		newStmt, err := sqlutil.RemoveComments(stmt)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("new_sql:%v\n", newStmt)
		fmt.Println()
	}
}
```
