package sqlutil

import (
	"testing"
)

func TestRemoveComments(t *testing.T) {
	testCases := []struct {
		input  string
		output string
	}{
		{
			input:  "/* comments */delete from abc",
			output: " delete from abc",
		},
		{
			input: "-- ----------------------------\n" +
				"select 1" +
				"-- ----------------------------\n",
			output: "select 1",
		},
		{
			input: "-- ----------------------------\n" +
				"-- Table structure for osc\n" +
				"-- ----------------------------\n" +
				"DROP TABLE IF EXISTS `osc`; CREATE TABLE `osc` ( `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id', `appename` varchar(32) NOT NULL COMMENT '应用英文名', `keyspace` varchar(32) DEFAULT NULL COMMENT '逻辑库', `tablename` varchar(128) DEFAULT NULL COMMENT '表名', `shard` varchar(16) DEFAULT NULL COMMENT '分片', `host` varchar(64) DEFAULT NULL COMMENT 'host', `alter` varchar(256) DEFAULT NULL COMMENT 'alter语句', `status` int(11) DEFAULT NULL COMMENT '状态', `message` text COMMENT '信息', `createtime` timestamp NOT NULL DEFAULT '2017-01-01 00:00:00' COMMENT '创建时间', `updatetime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间', PRIMARY KEY (`id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8;\n-- ----------------------------\n-- Records of osc\n-- ----------------------------\n",
			output: "DROP TABLE IF EXISTS `osc`; CREATE TABLE `osc` ( `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id', `appename` varchar(32) NOT NULL COMMENT '应用英文名', `keyspace` varchar(32) DEFAULT NULL COMMENT '逻辑库', `tablename` varchar(128) DEFAULT NULL COMMENT '表名', `shard` varchar(16) DEFAULT NULL COMMENT '分片', `host` varchar(64) DEFAULT NULL COMMENT 'host', `alter` varchar(256) DEFAULT NULL COMMENT 'alter语句', `status` int(11) DEFAULT NULL COMMENT '状态', `message` text COMMENT '信息', `createtime` timestamp NOT NULL DEFAULT '2017-01-01 00:00:00' COMMENT '创建时间', `updatetime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间', PRIMARY KEY (`id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8;\n",
		},
		{
			input: "begin;delete from messages where 1 =1;INSERT into messages(page, time_created_ns, message) values('10', 1, 1), ('20', 2, 1), ('30', 3, 1), ('40', 4, 1);update messages set message='2' where message='1';select * from messages where 1 =1;commit;",
		},
		{
			input: "CREATE TABLE if not exists `student` ( `id` int(10) unsigned NOT NULL AUTO_INCREMENT, `name` varchar(64) NOT NULL DEFAULT '' comment 'haha', `age` int NOT NULL DEFAULT 0 comment 'age for student;', `class` char(64) NOT NULL DEFAULT '', PRIMARY KEY (`id`));select * from student;",
		},
		{
			input: "-- \n" +
				"-- Dumping data for table `dbsv3_users`\n" +
				"-- LOCK TABLES `dbsv3_users` WRITE;\n" +
				"/*!40000 ALTER TABLE `dbsv3_users` DISABLE KEYS */; \n",
			output: "/*!40000 ALTER TABLE `dbsv3_users` DISABLE KEYS */; \n",
		},
		{
			input: "/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;\n" +
				"/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;\n" +
				"/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;\n",
			output: "/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;\n" +
				"/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;\n" +
				"/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;\n",
		},
		{
			input: "--\n" +
				"-- Table structure for table `dbsv3_users`; is important\n" +
				"select 1",
			output: "select 1",
		},
		{
			input: `set @sharding = 't_news_comments_reply singleShard'
CREATE TABLE t_news_comments_reply (
  ID bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  PARTENT_ID bigint(20) NOT NULL AUTO_INCREMENT COMMENT '父级标识,如果该条记录是评论，此字段代表新闻id；如果该条记录是回复，此字段代表评论id或上级回复id',
  CONTENT varchar(2000) DEFAULT NULL COMMENT '评论或回复内容',
  IS_ANONYMOUS int(4) DEFAULT 0 comment '是否匿名(特指评论), 1:匿名, 0:非匿名',
  OPERATION_TYPE int(4) DEFAULT 1 comment '操作类型, 1:评论, 2:回复;',
  PUBLISH_TIME timestamp NULL DEFAULT NULL COMMENT '评论或回复时间;',
  DELFLAG tinyint(4) DEFAULT NULL COMMENT '是否删除，1：删除;0：使用',
  CREATOR varchar(256) DEFAULT NULL COMMENT '创建者',
  CREATETIME timestamp NULL DEFAULT NULL COMMENT '创建时间',
  UPDATER varchar(256) DEFAULT NULL COMMENT '最后修改者',
  UPDATETIME timestamp NULL DEFAULT NULL COMMENT '最后修改时间',
  PRIMARY KEY (ID)
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8 COMMENT='评论-回复表';`,
			output: `set @sharding = 't_news_comments_reply singleShard'
CREATE TABLE t_news_comments_reply (
  ID bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  PARTENT_ID bigint(20) NOT NULL AUTO_INCREMENT COMMENT '父级标识,如果该条记录是评论，此字段代表新闻id；如果该条记录是回复，此字段代表评论id或上级回复id',
  CONTENT varchar(2000) DEFAULT NULL COMMENT '评论或回复内容',
  IS_ANONYMOUS int(4) DEFAULT 0 comment '是否匿名(特指评论), 1:匿名, 0:非匿名',
  OPERATION_TYPE int(4) DEFAULT 1 comment '操作类型, 1:评论, 2:回复;',
  PUBLISH_TIME timestamp NULL DEFAULT NULL COMMENT '评论或回复时间;',
  DELFLAG tinyint(4) DEFAULT NULL COMMENT '是否删除，1：删除;0：使用',
  CREATOR varchar(256) DEFAULT NULL COMMENT '创建者',
  CREATETIME timestamp NULL DEFAULT NULL COMMENT '创建时间',
  UPDATER varchar(256) DEFAULT NULL COMMENT '最后修改者',
  UPDATETIME timestamp NULL DEFAULT NULL COMMENT '最后修改时间',
  PRIMARY KEY (ID)
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8 COMMENT='评论-回复表';`,
		},
		{
			input:  "/*!40101 ######SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;",
			output: "/*!40101 ######SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;",
		},
	}

	for idx, tc := range testCases {
		output := tc.output
		if len(output) < 1 {
			output = tc.input
		}

		result, err := RemoveComments(tc.input)
		if err != nil {
			t.Fatalf(err.Error())
		}

		if result != output {
			t.Errorf("remove comments from %dth SQL:%s\nexpect: %#v,\nget: %#v", idx+1, tc.input, output, result)
		}
	}
}
