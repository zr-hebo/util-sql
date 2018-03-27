package sqlutil

var testSQLs = []string{
	"",
	"",

	"",
	"",
	"",
	"",
	"",
	"",
	"",
	`set @sharding = 't_news singleShard'
CREATE TABLE t_news (
  ID bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  TITLE varchar(1000) DEFAULT NULL COMMENT '新闻标题',
  SUB_TITLE varchar(2000) DEFAULT NULL COMMENT '新闻副标题',
  TYPE int(4) DEFAULT 1 comment '新闻类型, 1:文字新闻; 2:图片新闻; 3:重点资讯; 4:平台资讯',
  PUBLISH_STATUS int(4) DEFAULT 0 comment '发布状态, 1:发布, 0:未发布',
  VALIDITY_START timestamp NULL DEFAULT NULL COMMENT '有效期开始时间',
  VALIDITY_END timestamp NULL DEFAULT NULL COMMENT '有效期结束时间',
  IMPORT_LEVEL int(4) DEFAULT 1 comment '重要等级, 1:一般, 2:中等, 3：重要',
  IS_COMMENTS int(4) DEFAULT 0 comment '是否可以被评论, 1:可以, 0:不可以;',
  PUBLISH_TIME timestamp NULL DEFAULT NULL COMMENT '新闻发布时间',
  HITS bigint(20) DEFAULT NULL COMMENT '点击量',
  DELFLAG tinyint(4) DEFAULT NULL COMMENT '是否删除，1：删除，0：使用',
  REMOTEURL varchar(512) DEFAULT NULL COMMENT '云端地址',
  CREATOR varchar(256) DEFAULT NULL COMMENT '创建者',
  CREATETIME timestamp NULL DEFAULT NULL COMMENT '创建时间',
  UPDATER varchar(256) DEFAULT NULL COMMENT '最后修改者',
  UPDATETIME timestamp NULL DEFAULT NULL COMMENT '最后修改时间',
  PRIMARY KEY (ID)
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8 COMMENT='新闻表';`,
	"",
}
