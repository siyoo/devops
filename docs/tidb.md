1. 基本思想
Newsql 引导者 Google F1 数据库
<img src="https://github.com/siyoo/devops/blob/master/images/google spanner.png" height="400px" width="400px" />

F1 Query:
Spanner:
强一致的CP系统


2. tidb架构
1. sql protocol
2. parser + plan
3. tablet
所有table 数据划分为tablet存储在KVStore 层.  每个tablet 会有一个独立的raft group 来保证 各个replica 的数据一致.  每个raft goup 会有一个leader lease 负责这个tablet 数据的读写
4. 分布式思想（F1 Query）
5. 存储 badger (rocksdb golang实现)
6. 同步数据 google spanner 分布式存储


1.在mysql client输入sql执行之后发生了什么

词法分析(类Lex:https://github.com/mysql/mysql-server/blob/5.6/sql/sql_lex.cc)

词法语法解析的处理过程根编译原理上的东西基本类似；MySQL 并没有使用 lex 来实现词法分析，但是语法分析却用了 yacc
Lex-Yacc (词法扫描器-语法分析器) 在 Linux 上就是 flex-bison, 可以通过 yum install flex flex-devel bison 进行安装

语法分析(Lex)
语义分析(Yacc:https://github.com/mysql/mysql-server/blob/5.6/sql/sql_yacc.yy)
构造执行树
生成执行计划
计划的执行

https://github.com/mysql/mysql-server/blob/5.6/sql/sql_lex.cc

https://github.com/mysql/mysql-server/blob/5.6/sql/sql_yacc.yy


2.
handshake收到的信息

user password auth_type os client_version pid os_user

auth_type: caching_sha2_password(8.0) mysql_native_password(<8.0)

tidb 会切换 caching_sha2_password 的 data 为 mysql_native_password 的 data
```
	enclen := 1 + len("mysql_native_password") + 1 + len(cc.salt) + 1
	data := cc.alloc.AllocWithLen(4, enclen)
	data = append(data, 0xfe) // switch request
	data = append(data, []byte("mysql_native_password")...)
	data = append(data, byte(0x00)) // requires null
	data = append(data, cc.salt...)
	data = append(data, 0)
```

handshake 

校验 host/password, host && pwd 都在user表，user表在tidb启动&&每五分钟&&user表变更的时候，加载到内存中 

handshake之后建立连接，阻塞，等待接受数据包

收到数据包

dispatch 函数对收到的数据包进行解析
第一个byte是cmd type
```
	ComSleep byte = iota
	ComQuit
	ComInitDB
	ComQuery
	ComFieldList
	ComCreateDB
	ComDropDB
	ComRefresh
	ComShutdown
	ComStatistics
	ComProcessInfo
	ComConnect
	ComProcessKill
	ComDebug
	ComPing
	ComTime
	ComDelayedInsert
	ComChangeUser
	ComBinlogDump
	ComTableDump
	ComConnectOut
	ComRegisterSlave
	ComStmtPrepare
	ComStmtExecute
	ComStmtSendLongData
	ComStmtClose
	ComStmtReset
	ComSetOption
	ComStmtFetch
	ComDaemon
	ComBinlogDumpGtid
	ComResetConnection
	ComEnd
```


后面的为sql

根据不同的cmd，走到不同的handler
1. parse sql  cc.ctx.Parse(ctx, sql)
yyLexer + yyParse

