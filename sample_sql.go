package sqlutil

var testSQLs = []string{
	"/* asdfasfdsdf */delete from abc",
	"-- ----------------------------\n" +
		"-- Table structure for osc" +
		"-- ----------------------------\n",
	"-- ----------------------------\n" +
		"-- Table structure for osc\n" +
		"-- ----------------------------\n" +
		"DROP TABLE IF EXISTS `osc`; " +
		"CREATE TABLE `osc` ( `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id', `appename` varchar(32) NOT NULL COMMENT '应用英文名', `keyspace` varchar(32) DEFAULT NULL COMMENT '逻辑库', `tablename` varchar(128) DEFAULT NULL COMMENT '表名', `shard` varchar(16) DEFAULT NULL COMMENT '分片', `host` varchar(64) DEFAULT NULL COMMENT 'host', `alter` varchar(256) DEFAULT NULL COMMENT 'alter语句', `status` int(11) DEFAULT NULL COMMENT '状态', `message` text COMMENT '信息', `createtime` timestamp NOT NULL DEFAULT '2017-01-01 00:00:00' COMMENT '创建时间', `updatetime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间', PRIMARY KEY (`id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8;\n-- ----------------------------\n-- Records of osc\n-- ----------------------------\n",
	"begin;delete from messages where 1 =1;INSERT into messages(page, time_created_ns, message) values('10', 1, 1), ('20', 2, 1), ('30', 3, 1), ('40', 4, 1);update messages set message='2' where message='1';select * from messages where 1 =1;commit;",
	"CREATE TABLE if not exists `student` ( `id` int(10) unsigned NOT NULL AUTO_INCREMENT, `name` varchar(64) NOT NULL DEFAULT '' comment 'haha', `age` int NOT NULL DEFAULT 0 comment 'age for student;', `class` char(64) NOT NULL DEFAULT '', PRIMARY KEY (`id`));select * from student;",
	"-- \n" +
		"-- Dumping data for table `dbsv3_users`\n" +
		"-- LOCK TABLES `dbsv3_users` WRITE;\n" +
		"/*!40000 ALTER TABLE `dbsv3_users` DISABLE KEYS */; \n",
	"/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;\n" +
		"/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;\n" +
		"/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;\n",
	"--\n" +
		"-- Table structure for table `dbsv3_users`; is important\n" +
		"--",
	"show create table student(id int, name varchar(64) comment 'a student name; maybe empty; eg:zhangsan', age int)",
}
