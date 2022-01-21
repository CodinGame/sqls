package dialect

var mysql8Keyword = []string{
	"ACCOUNT",
	"ACTION",
	"ADD",
	"ADMIN",
	"AFTER",
	"AGGREGATE",
	"ALGORITHM",
	"ALL",
	"ALTER",
	"ANALYZE",
	"ARCHIVE",
	"AS",
	"ASC",
	"AT",
	"ATTRIBUTE",
	"AUTOCOMMIT",
	"AUTOEXTEND_SIZE",
	"AUTO_INCREMENT",
	"AVG_ROW_LENGTH",
	"BACKUP",
	"BEFORE",
	"BEGIN",
	"BINARY",
	"BINLOG",
	"BOOL",
	"BOOLEAN",
	"BTREE",
	"BY",
	"BYTE",
	"CACHE",
	"CALL",
	"CASCADE",
	"CASE",
	"CATALOG_NAME",
	"CHAIN",
	"CHANGE",
	"CHANNEL",
	"CHAR",
	"CHARACTER",
	"CHARSET",
	"CHECK",
	"CHECKSUM",
	"CIPHER",
	"CLASS_ORIGIN",
	"CLIENT",
	"CLONE",
	"CLOSE",
	"COALESCE",
	"CODE",
	"COLLATE",
	"COLLATION",
	"COLUMN",
	"COLUMNS",
	"COLUMN_NAME",
	"COMMENT",
	"COMMIT",
	"COMMITTED",
	"COMPACT",
	"COMPLETION",
	"COMPONENT",
	"COMPRESSED",
	"COMPRESSION",
	"CONCURRENT",
	"CONDITION",
	"CONNECTION",
	"CONSISTENT",
	"CONSTRAINT",
	"CONSTRAINT_CATALOG",
	"CONSTRAINT_NAME",
	"CONSTRAINT_SCHEMA",
	"CONTINUE",
	"COUNT",
	"CREATE",
	"CROSS",
	"CSV",
	"CURRENT",
	"CURRENT_USER",
	"CURSOR",
	"CURSOR_NAME",
	"DATA",
	"DATABASE",
	"DATABASES",
	"DATAFILE",
	"DATE",
	"DAY",
	"DEALLOCATE",
	"DEC",
	"DECIMAL",
	"DECLARE",
	"DEFAULT",
	"DEFAULT_AUTH",
	"DEFINER",
	"DEFINITION",
	"DELAYED",
	"DELAY_KEY_WRITE",
	"DELETE",
	"DESC",
	"DESCRIBE",
	"DESCRIPTION",
	"DIAGNOSTICS",
	"DIRECTORY",
	"DISABLE",
	"DISCARD",
	"DISTINCT",
	"DISTINCTROW",
	"DO",
	"DROP",
	"DUAL",
	"DUMPFILE",
	"DUPLICATE",
	"DYNAMIC",
	"ELSE",
	"ELSEIF",
	"ENABLE",
	"ENCLOSED",
	"ENCRYPTION",
	"END",
	"ENDS",
	"ENGINE",
	"ENGINES",
	"ERROR",
	"ERRORS",
	"ESCAPED",
	"EVENT",
	"EVENTS",
	"EVERY",
	"EXCEPT",
	"EXCHANGE",
	"EXECUTE",
	"EXISTS",
	"EXIT",
	"EXPIRE",
	"EXPLAIN",
	"EXPORT",
	"EXTENDED",
	"EXTENT_SIZE",
	"FAILED_LOGIN_ATTEMPTS",
	"FALSE",
	"FAST",
	"FEDERATED",
	"FETCH",
	"FIELDS",
	"FILE",
	"FILE_BLOCK_SIZE",
	"FILTER",
	"FIRST",
	"FIXED",
	"FLOAT4",
	"FLOAT8",
	"FLUSH",
	"FOR",
	"FORCE",
	"FOREIGN",
	"FORMAT",
	"FROM",
	"FULL",
	"FULLTEXT",
	"FUNCTION",
	"GENERAL",
	"GET",
	"GLOBAL",
	"GRANT",
	"GRANTS",
	"GROUP",
	"HANDLER",
	"HAVING",
	"HEAP",
	"HELP",
	"HELP_DATE",
	"HELP_VERSION",
	"HIGH_PRIORITY",
	"HISTORY",
	"HOST",
	"HOSTS",
	"IDENTIFIED",
	"IF",
	"IGNORE",
	"IGNORE_SERVER_IDS",
	"IMPORT",
	"IN",
	"INDEX",
	"INDEXES",
	"INFILE",
	"INITIAL_SIZE",
	"INNER",
	"INNODB",
	"INSERT",
	"INSERT_METHOD",
	"INSTALL",
	"INSTANCE",
	"INT1",
	"INT2",
	"INT3",
	"INT4",
	"INT8",
	"INTEGER",
	"INTERVAL",
	"INTO",
	"INVISIBLE",
	"IO_THREAD",
	"ISOLATION",
	"ISSUER",
	"ITERATE",
	"JOIN",
	"JSON",
	"KEY",
	"KEYS",
	"KEY_BLOCK_SIZE",
	"KILL",
	"LAST",
	"LEAVE",
	"LEAVES",
	"LEFT",
	"LEVEL",
	"LIKE",
	"LIMIT",
	"LINES",
	"LOAD",
	"LOCAL",
	"LOCK",
	"LOGFILE",
	"LOGS",
	"LONG",
	"LONGBINARY",
	"LOOP",
	"LOW_PRIORITY",
	"MASTER",
	"MASTER_AUTO_POSITION",
	"MASTER_BIND",
	"MASTER_COMPRESSION_ALGORITHMS",
	"MASTER_CONNECT_RETRY",
	"MASTER_HEARTBEAT_PERIOD",
	"MASTER_HOST",
	"MASTER_LOG_FILE",
	"MASTER_LOG_POS",
	"MASTER_PASSWORD",
	"MASTER_PORT",
	"MASTER_RETRY_COUNT",
	"MASTER_SSL",
	"MASTER_SSL_CA",
	"MASTER_SSL_CERT",
	"MASTER_SSL_CIPHER",
	"MASTER_SSL_CRL",
	"MASTER_SSL_CRLPATH",
	"MASTER_SSL_KEY",
	"MASTER_SSL_VERIFY_SERVER_CERT",
	"MASTER_TLS_CIPHERSUITES",
	"MASTER_TLS_VERSION",
	"MASTER_USER",
	"MASTER_ZSTD_COMPRESSION_LEVEL",
	"MAX_CONNECTIONS_PER_HOUR",
	"MAX_QUERIES_PER_HOUR",
	"MAX_ROWS",
	"MAX_SIZE",
	"MAX_UPDATES_PER_HOUR",
	"MAX_USER_CONNECTIONS",
	"MEDIUM",
	"MEMORY",
	"MERGE",
	"MESSAGE_TEXT",
	"MIDDLEINT",
	"MIN_ROWS",
	"MODE",
	"MODIFY",
	"MRG_MYISAM",
	"MUTEX",
	"MYISAM",
	"MYSQL_ERRNO",
	"NAME",
	"NAMES",
	"NATIONAL",
	"NATURAL",
	"NCHAR",
	"NDB",
	"NDBCLUSTER",
	"NETWORK_NAMESPACE",
	"NEVER",
	"NEXT",
	"NO",
	"NODEGROUP",
	"NONE",
	"NOT",
	"NO_WRITE_TO_BINLOG",
	"NULL",
	"NUMBER",
	"NUMERIC",
	"NVARCHAR",
	"OFFSET",
	"OLD",
	"ON",
	"ONLY",
	"OPEN",
	"OPTIMIZE",
	"OPTIMIZER_COSTS",
	"OPTION",
	"OPTIONAL",
	"OPTIONALLY",
	"OPTIONS",
	"OR",
	"ORDER",
	"ORGANIZATION",
	"OUTER",
	"OUTFILE",
	"OWNER",
	"PACK_KEYS",
	"PARSER",
	"PARTIAL",
	"PARTITION",
	"PARTITIONING",
	"PASSWORD",
	"PASSWORD_LOCK_TIME",
	"PERSIST",
	"PERSIST_ONLY",
	"PLUGIN",
	"PLUGINS",
	"PLUGIN_DIR",
	"PORT",
	"PRECISION",
	"PREPARE",
	"PRESERVE",
	"PREV",
	"PRIMARY",
	"PRIVILEGES",
	"PRIVILEGE_CHECKS_USER",
	"PROCEDURE",
	"PROCESS",
	"PROCESSLIST",
	"PROFILE",
	"PROFILES",
	"PROXY",
	"PURGE",
	"QUERY",
	"QUICK",
	"RANDOM",
	"READ",
	"REAL",
	"REBUILD",
	"RECOVER",
	"REDO_LOG",
	"REDUNDANT",
	"REFERENCE",
	"REFERENCES",
	"RELAY",
	"RELAYLOG",
	"RELAY_LOG_FILE",
	"RELAY_LOG_POS",
	"RELEASE",
	"RELOAD",
	"REMOVE",
	"RENAME",
	"REORGANIZE",
	"REPAIR",
	"REPEAT",
	"REPEATABLE",
	"REPLACE",
	"REPLICA",
	"REPLICAS",
	"REPLICATE_DO_DB",
	"REPLICATE_DO_TABLE",
	"REPLICATE_IGNORE_DB",
	"REPLICATE_IGNORE_TABLE",
	"REPLICATE_REWRITE_DB",
	"REPLICATE_WILD_DO_TABLE",
	"REPLICATE_WILD_IGNORE_TABLE",
	"REPLICATION",
	"REQUIRE",
	"RESET",
	"RESIGNAL",
	"RESOURCE",
	"RESTART",
	"RESTRICT",
	"RETAIN",
	"RETURN",
	"RETURNED_SQLSTATE",
	"RETURNS",
	"REUSE",
	"REVOKE",
	"RIGHT",
	"ROLE",
	"ROLLBACK",
	"ROW",
	"ROWS",
	"ROW_COUNT",
	"ROW_FORMAT",
	"SAVEPOINT",
	"SCHEDULE",
	"SCHEMA",
	"SCHEMAS",
	"SCHEMA_NAME",
	"SELECT",
	"SERIAL",
	"SERIALIZABLE",
	"SERVER",
	"SESSION",
	"SET",
	"SHARE",
	"SHOW",
	"SHUTDOWN",
	"SIGNAL",
	"SLAVE",
	"SLOW",
	"SNAPSHOT",
	"SOCKET",
	"SONAME",
	"SPATIAL",
	"SQLSTATE",
	"SQL_AFTER_GTIDS",
	"SQL_AFTER_MTS_GAPS",
	"SQL_BEFORE_GTIDS",
	"SQL_BIG_RESULT",
	"SQL_BUFFER_RESULT",
	"SQL_CALC_FOUND_ROWS",
	"SQL_LOG_BIN",
	"SQL_NO_CACHE",
	"SQL_SMALL_RESULT",
	"SQL_THREAD",
	"SSL",
	"START",
	"STARTING",
	"STARTS",
	"STATS_AUTO_RECALC",
	"STATS_PERSISTENT",
	"STATS_SAMPLE_PAGES",
	"STATUS",
	"STOP",
	"STORAGE",
	"STORED",
	"STRAIGHT_JOIN",
	"STRING",
	"SUBCLASS_ORIGIN",
	"SUBJECT",
	"SUPER",
	"SYSTEM",
	"TABLE",
	"TABLES",
	"TABLESPACE",
	"TABLE_NAME",
	"TEMPORARY",
	"TERMINATED",
	"THEN",
	"THREAD_PRIORITY",
	"TIME",
	"TIMESTAMP",
	"TLS",
	"TO",
	"TRADITIONAL",
	"TRANSACTION",
	"TREE",
	"TRIGGER",
	"TRIGGERS",
	"TRUE",
	"TRUNCATE",
	"TYPE",
	"UNBOUNDED",
	"UNCOMMITTED",
	"UNDO",
	"UNINSTALL",
	"UNION",
	"UNIQUE",
	"UNLOCK",
	"UNSIGNED",
	"UNTIL",
	"UPDATE",
	"UPGRADE",
	"USAGE",
	"USE",
	"USER",
	"USER_RESOURCES",
	"USE_FRM",
	"USING",
	"VALUE",
	"VALUES",
	"VARCHARACTER",
	"VARIABLE",
	"VARIABLES",
	"VARYING",
	"VCPU",
	"VIEW",
	"VIRTUAL",
	"VISIBLE",
	"WAIT",
	"WARNINGS",
	"WHEN",
	"WHERE",
	"WHILE",
	"WITH",
	"WORK",
	"WRAPPER",
	"WRITE",
	"X509",
	"XA",
	"ZEROFILL",
}

var mysql8Function = []string{
	"%",
	"&",
	"(JSON",
	"*",
	"+",
	"-",
	"->",
	"->>",
	"/",
	":=",
	"<",
	"<<",
	"<=",
	"<=>",
	"<>",
	"=",
	">",
	">=",
	">>",
	"ABS",
	"ACOS",
	"ADDDATE",
	"ADDTIME",
	"AES_DECRYPT",
	"AES_ENCRYPT",
	"AGAINST",
	"AND",
	"ANY_VALUE",
	"ARRAY",
	"ASC",
	"ASCII",
	"ASIN",
	"ASYMMETRIC_DECRYPT",
	"ASYMMETRIC_DERIVE",
	"ASYMMETRIC_ENCRYPT",
	"ASYMMETRIC_SIGN",
	"ASYMMETRIC_VERIFY",
	"ATAN",
	"ATAN2",
	"AVG",
	"BENCHMARK",
	"BETWEEN",
	"BIN",
	"BINARY",
	"BIN_TO_UUID",
	"BIT_AND",
	"BIT_COUNT",
	"BIT_LENGTH",
	"BIT_OR",
	"BIT_XOR",
	"BOOLEAN",
	"BOTH",
	"BY",
	"CAN_ACCESS_COLUMN",
	"CAN_ACCESS_DATABASE",
	"CAN_ACCESS_TABLE",
	"CAN_ACCESS_USER",
	"CAN_ACCESS_VIEW",
	"CASE",
	"CAST",
	"CEIL",
	"CEILING",
	"CHAR",
	"CHARACTER_LENGTH",
	"CHARSET",
	"CHAR_LENGTH",
	"COALESCE",
	"COERCIBILITY",
	"COLLATION",
	"COMPRESS",
	"CONCAT",
	"CONCAT_WS",
	"CONNECTION_ID",
	"CONV",
	"CONVERT",
	"CONVERT_TZ",
	"COS",
	"COT",
	"COUNT",
	"CRC32",
	"CREATE_ASYMMETRIC_PRIV_KEY",
	"CREATE_ASYMMETRIC_PUB_KEY",
	"CREATE_DH_PARAMETERS",
	"CREATE_DIGEST",
	"CUME_DIST",
	"CURDATE",
	"CURRENT_DATE",
	"CURRENT_ROLE",
	"CURRENT_TIME",
	"CURRENT_TIMESTAMP",
	"CURRENT_USER",
	"CURTIME",
	"DATABASE",
	"DATE",
	"DATEDIFF",
	"DATETIME",
	"DATE_ADD",
	"DATE_FORMAT",
	"DATE_SUB",
	"DAY",
	"DAYNAME",
	"DAYOFMONTH",
	"DAYOFWEEK",
	"DAYOFYEAR",
	"DAY_HOUR",
	"DAY_MINUTE",
	"DAY_SECOND",
	"DECIMAL",
	"DEFAULT",
	"DEGREES",
	"DENSE_RANK",
	"DESC",
	"DISTINCT",
	"DIV",
	"ELSE",
	"ELT",
	"END",
	"ESCAPE",
	"EXP",
	"EXPANSION",
	"EXPORT_SET",
	"EXTRACT",
	"EXTRACTION)",
	"EXTRACTVALUE",
	"FIELD",
	"FIND_IN_SET",
	"FIRST_VALUE",
	"FLOOR",
	"FORMAT",
	"FORMAT_BYTES",
	"FORMAT_PICO_TIME",
	"FOUND_ROWS",
	"FROM",
	"FROM_BASE64",
	"FROM_DAYS",
	"FROM_UNIXTIME",
	"FUNCTION",
	"GEOMCOLLECTION",
	"GEOMETRYCOLLECTION",
	"GET_DD_COLUMN_PRIVILEGES",
	"GET_DD_CREATE_OPTIONS",
	"GET_DD_INDEX_SUB_PART_LENGTH",
	"GET_FORMAT",
	"GET_LOCK",
	"GREATEST",
	"GROUPING",
	"GROUP_CONCAT",
	"GTID_SUBSET",
	"GTID_SUBTRACT",
	"HEX",
	"HOUR",
	"HOUR_MINUTE",
	"HOUR_SECOND",
	"ICU_VERSION",
	"IF",
	"IFNULL",
	"IN",
	"INET6_ATON",
	"INET6_NTOA",
	"INET_ATON",
	"INET_NTOA",
	"INLINE",
	"INSERT",
	"INSTR",
	"INTEGER",
	"INTERNAL_AUTO_INCREMENT",
	"INTERNAL_AVG_ROW_LENGTH",
	"INTERNAL_CHECKSUM",
	"INTERNAL_CHECK_TIME",
	"INTERNAL_DATA_FREE",
	"INTERNAL_DATA_LENGTH",
	"INTERNAL_DD_CHAR_LENGTH",
	"INTERNAL_GET_COMMENT_OR_ERROR",
	"INTERNAL_GET_ENABLED_ROLE_JSON",
	"INTERNAL_GET_HOSTNAME",
	"INTERNAL_GET_USERNAME",
	"INTERNAL_GET_VIEW_WARNING_OR_ERROR",
	"INTERNAL_INDEX_COLUMN_CARDINALITY",
	"INTERNAL_INDEX_LENGTH",
	"INTERNAL_IS_ENABLED_ROLE",
	"INTERNAL_IS_MANDATORY_ROLE",
	"INTERNAL_KEYS_DISABLED",
	"INTERNAL_MAX_DATA_LENGTH",
	"INTERNAL_TABLE_ROWS",
	"INTERNAL_UPDATE_TIME",
	"INTERVAL",
	"IS",
	"ISNULL",
	"IS_FREE_LOCK",
	"IS_IPV4",
	"IS_IPV4_COMPAT",
	"IS_IPV4_MAPPED",
	"IS_IPV6",
	"IS_USED_LOCK",
	"IS_UUID",
	"IS_VISIBLE_DD_OBJECT",
	"JSON",
	"JSON_ARRAY",
	"JSON_ARRAYAGG",
	"JSON_ARRAY_APPEND",
	"JSON_ARRAY_INSERT",
	"JSON_CONTAINS",
	"JSON_CONTAINS_PATH",
	"JSON_DEPTH",
	"JSON_EXTRACT",
	"JSON_INSERT",
	"JSON_KEYS",
	"JSON_LENGTH",
	"JSON_MERGE",
	"JSON_MERGE_PATCH",
	"JSON_MERGE_PRESERVE",
	"JSON_OBJECT",
	"JSON_OBJECTAGG",
	"JSON_OVERLAPS",
	"JSON_PRETTY",
	"JSON_QUOTE",
	"JSON_REMOVE",
	"JSON_REPLACE",
	"JSON_SCHEMA_VALID",
	"JSON_SCHEMA_VALIDATION_REPORT",
	"JSON_SEARCH",
	"JSON_SET",
	"JSON_STORAGE_FREE",
	"JSON_STORAGE_SIZE",
	"JSON_TABLE",
	"JSON_TYPE",
	"JSON_UNQUOTE",
	"JSON_VALID",
	"JSON_VALUE",
	"LAG",
	"LAST_DAY",
	"LAST_INSERT_ID",
	"LAST_VALUE",
	"LCASE",
	"LEAD",
	"LEADING",
	"LEAST",
	"LEFT",
	"LENGTH",
	"LIKE",
	"LINESTRING",
	"LN",
	"LOAD_FILE",
	"LOCALTIME",
	"LOCALTIMESTAMP",
	"LOCATE",
	"LOG",
	"LOG10",
	"LOG2",
	"LOWER",
	"LPAD",
	"LTRIM",
	"MAKEDATE",
	"MAKETIME",
	"MAKE_SET",
	"MASTER_POS_WAIT",
	"MATCH",
	"MAX",
	"MBRCONTAINS",
	"MBRCOVEREDBY",
	"MBRCOVERS",
	"MBRDISJOINT",
	"MBREQUALS",
	"MBRINTERSECTS",
	"MBROVERLAPS",
	"MBRTOUCHES",
	"MBRWITHIN",
	"MD5",
	"MEMBER",
	"MICROSECOND",
	"MID",
	"MIN",
	"MINUTE",
	"MINUTE_SECOND",
	"MOD",
	"MODE",
	"MONTH",
	"MONTHNAME",
	"MULTILINESTRING",
	"MULTIPOINT",
	"MULTIPOLYGON",
	"NAME_CONST",
	"NOT",
	"NOW",
	"NTH_VALUE",
	"NTILE",
	"NULL",
	"NULLIF",
	"OCT",
	"OCTET_LENGTH",
	"OF",
	"OR",
	"ORD",
	"ORDER",
	"PATH)",
	"PERCENT_RANK",
	"PERIOD_ADD",
	"PERIOD_DIFF",
	"PI",
	"POINT",
	"POLYGON",
	"POSITION",
	"POW",
	"POWER",
	"PS_CURRENT_THREAD_ID",
	"PS_THREAD_ID",
	"QUARTER",
	"QUERY",
	"QUOTE",
	"RADIANS",
	"RAND",
	"RANDOM_BYTES",
	"RANK",
	"REGEXP",
	"REGEXP_INSTR",
	"REGEXP_LIKE",
	"REGEXP_REPLACE",
	"REGEXP_SUBSTR",
	"RELEASE_ALL_LOCKS",
	"RELEASE_LOCK",
	"REPEAT",
	"REPLACE",
	"REVERSE",
	"RIGHT",
	"RLIKE",
	"ROLES_GRAPHML",
	"ROUND",
	"ROW_COUNT",
	"ROW_NUMBER",
	"RPAD",
	"RTRIM",
	"SCHEMA",
	"SECOND",
	"SEC_TO_TIME",
	"SEPARATOR",
	"SESSION_USER",
	"SHA",
	"SHA1",
	"SHA2",
	"SIGN",
	"SIGNED",
	"SIN",
	"SLEEP",
	"SOUNDEX",
	"SOUNDS",
	"SPACE",
	"SQRT",
	"STATEMENT_DIGEST",
	"STATEMENT_DIGEST_TEXT",
	"STD",
	"STDDEV",
	"STDDEV_POP",
	"STDDEV_SAMP",
	"STRCMP",
	"STR_TO_DATE",
	"ST_AREA",
	"ST_ASBINARY",
	"ST_ASGEOJSON",
	"ST_ASTEXT",
	"ST_ASWKB",
	"ST_ASWKT",
	"ST_BUFFER",
	"ST_BUFFER_STRATEGY",
	"ST_CENTROID",
	"ST_CONTAINS",
	"ST_CONVEXHULL",
	"ST_CROSSES",
	"ST_DIFFERENCE",
	"ST_DIMENSION",
	"ST_DISJOINT",
	"ST_DISTANCE",
	"ST_DISTANCE_SPHERE",
	"ST_ENDPOINT",
	"ST_ENVELOPE",
	"ST_EQUALS",
	"ST_EXTERIORRING",
	"ST_FRECHETDISTANCE",
	"ST_GEOHASH",
	"ST_GEOMCOLLFROMTEXT",
	"ST_GEOMCOLLFROMWKB",
	"ST_GEOMETRYCOLLECTIONFROMTEXT",
	"ST_GEOMETRYCOLLECTIONFROMWKB",
	"ST_GEOMETRYFROMTEXT",
	"ST_GEOMETRYFROMWKB",
	"ST_GEOMETRYN",
	"ST_GEOMETRYTYPE",
	"ST_GEOMFROMGEOJSON",
	"ST_GEOMFROMTEXT",
	"ST_GEOMFROMWKB",
	"ST_HAUSDORFFDISTANCE",
	"ST_INTERIORRINGN",
	"ST_INTERSECTION",
	"ST_INTERSECTS",
	"ST_ISCLOSED",
	"ST_ISEMPTY",
	"ST_ISSIMPLE",
	"ST_ISVALID",
	"ST_LATFROMGEOHASH",
	"ST_LATITUDE",
	"ST_LENGTH",
	"ST_LINEFROMTEXT",
	"ST_LINEFROMWKB",
	"ST_LINEINTERPOLATEPOINT",
	"ST_LINEINTERPOLATEPOINTS",
	"ST_LINESTRINGFROMTEXT",
	"ST_LINESTRINGFROMWKB",
	"ST_LONGFROMGEOHASH",
	"ST_LONGITUDE",
	"ST_MAKEENVELOPE",
	"ST_MLINEFROMTEXT",
	"ST_MLINEFROMWKB",
	"ST_MPOINTFROMTEXT",
	"ST_MPOINTFROMWKB",
	"ST_MPOLYFROMTEXT",
	"ST_MPOLYFROMWKB",
	"ST_MULTILINESTRINGFROMTEXT",
	"ST_MULTILINESTRINGFROMWKB",
	"ST_MULTIPOINTFROMTEXT",
	"ST_MULTIPOINTFROMWKB",
	"ST_MULTIPOLYGONFROMTEXT",
	"ST_MULTIPOLYGONFROMWKB",
	"ST_NUMGEOMETRIES",
	"ST_NUMINTERIORRING",
	"ST_NUMINTERIORRINGS",
	"ST_NUMPOINTS",
	"ST_OVERLAPS",
	"ST_POINTATDISTANCE",
	"ST_POINTFROMGEOHASH",
	"ST_POINTFROMTEXT",
	"ST_POINTFROMWKB",
	"ST_POINTN",
	"ST_POLYFROMTEXT",
	"ST_POLYFROMWKB",
	"ST_POLYGONFROMTEXT",
	"ST_POLYGONFROMWKB",
	"ST_SIMPLIFY",
	"ST_SRID",
	"ST_STARTPOINT",
	"ST_SWAPXY",
	"ST_SYMDIFFERENCE",
	"ST_TOUCHES",
	"ST_TRANSFORM",
	"ST_UNION",
	"ST_VALIDATE",
	"ST_WITHIN",
	"ST_X",
	"ST_Y",
	"SUBDATE",
	"SUBSTR",
	"SUBSTRING",
	"SUBSTRING_INDEX",
	"SUBTIME",
	"SUM",
	"SYSDATE",
	"SYSTEM_USER",
	"TAN",
	"THEN",
	"TIME",
	"TIMEDIFF",
	"TIMESTAMP",
	"TIMESTAMPADD",
	"TIMESTAMPDIFF",
	"TIMEZONE",
	"TIME_FORMAT",
	"TIME_TO_SEC",
	"TO_BASE64",
	"TO_DAYS",
	"TO_SECONDS",
	"TRAILING",
	"TRIM",
	"TRUNCATE",
	"UCASE",
	"UNCOMPRESS",
	"UNCOMPRESSED_LENGTH",
	"UNHEX",
	"UNIX_TIMESTAMP",
	"UNSIGNED",
	"UPDATEXML",
	"UPPER",
	"USER",
	"UTC_DATE",
	"UTC_TIME",
	"UTC_TIMESTAMP",
	"UUID",
	"UUID_SHORT",
	"UUID_TO_BIN",
	"VALIDATE_PASSWORD_STRENGTH",
	"VALUES",
	"VARIANCE",
	"VAR_POP",
	"VAR_SAMP",
	"VERSION",
	"WAIT_FOR_EXECUTED_GTID_SET",
	"WAIT_UNTIL_SQL_THREAD_AFTER_GTIDS",
	"WEEK",
	"WEEKDAY",
	"WEEKOFYEAR",
	"WEIGHT_STRING",
	"WHEN",
	"WITH",
	"XOR",
	"YEAR",
	"YEARWEEK",
	"YEAR_MONTH",
	"^",
	"|",
	"~",
}

var mysql57Keyword = []string{
	"ACCOUNT",
	"ACTION",
	"ADD",
	"AFTER",
	"AGGREGATE",
	"ALGORITHM",
	"ALL",
	"ALTER",
	"ANALYSE",
	"ANALYZE",
	"ARCHIVE",
	"AS",
	"ASC",
	"AT",
	"AUTOCOMMIT",
	"AUTOEXTEND_SIZE",
	"AUTO_INCREMENT",
	"AVG_ROW_LENGTH",
	"BEFORE",
	"BEGIN",
	"BINARY",
	"BINLOG",
	"BOOL",
	"BOOLEAN",
	"BTREE",
	"BY",
	"BYTE",
	"CACHE",
	"CALL",
	"CASCADE",
	"CASE",
	"CATALOG_NAME",
	"CHAIN",
	"CHANGE",
	"CHANNEL",
	"CHAR",
	"CHARACTER",
	"CHARSET",
	"CHECK",
	"CHECKSUM",
	"CIPHER",
	"CLASS_ORIGIN",
	"CLIENT",
	"CLOSE",
	"COALESCE",
	"CODE",
	"COLLATE",
	"COLLATION",
	"COLUMN",
	"COLUMNS",
	"COLUMN_NAME",
	"COMMENT",
	"COMMIT",
	"COMMITTED",
	"COMPACT",
	"COMPLETION",
	"COMPRESSED",
	"COMPRESSION",
	"CONCURRENT",
	"CONDITION",
	"CONNECTION",
	"CONSISTENT",
	"CONSTRAINT",
	"CONSTRAINT_CATALOG",
	"CONSTRAINT_NAME",
	"CONSTRAINT_SCHEMA",
	"CONTINUE",
	"COUNT",
	"CREATE",
	"CROSS",
	"CSV",
	"CURRENT_USER",
	"CURSOR",
	"CURSOR_NAME",
	"DATA",
	"DATABASE",
	"DATABASES",
	"DATAFILE",
	"DATE",
	"DEALLOCATE",
	"DEC",
	"DECIMAL",
	"DECLARE",
	"DEFAULT",
	"DEFAULT_AUTH",
	"DEFINER",
	"DELAYED",
	"DELAY_KEY_WRITE",
	"DELETE",
	"DESC",
	"DESCRIBE",
	"DES_KEY_FILE",
	"DIAGNOSTICS",
	"DIRECTORY",
	"DISABLE",
	"DISCARD",
	"DISTINCT",
	"DISTINCTROW",
	"DO",
	"DROP",
	"DUAL",
	"DUMPFILE",
	"DUPLICATE",
	"DYNAMIC",
	"ELSE",
	"ELSEIF",
	"ENABLE",
	"ENCLOSED",
	"ENCRYPTION",
	"END",
	"ENDS",
	"ENGINE",
	"ENGINES",
	"ERROR",
	"ERRORS",
	"ESCAPED",
	"EVENT",
	"EVENTS",
	"EVERY",
	"EXCHANGE",
	"EXECUTE",
	"EXISTS",
	"EXIT",
	"EXPIRE",
	"EXPLAIN",
	"EXPORT",
	"EXTENDED",
	"EXTENT_SIZE",
	"FALSE",
	"FAST",
	"FEDERATED",
	"FETCH",
	"FIELDS",
	"FILE",
	"FILE_BLOCK_SIZE",
	"FILTER",
	"FIRST",
	"FIXED",
	"FLOAT4",
	"FLOAT8",
	"FLUSH",
	"FOR",
	"FORCE",
	"FOREIGN",
	"FORMAT",
	"FROM",
	"FULL",
	"FULLTEXT",
	"FUNCTION",
	"GENERAL",
	"GET",
	"GLOBAL",
	"GRANT",
	"GRANTS",
	"GROUP",
	"HANDLER",
	"HAVING",
	"HEAP",
	"HELP",
	"HELP_DATE",
	"HELP_VERSION",
	"HIGH_PRIORITY",
	"HOST",
	"HOSTS",
	"IDENTIFIED",
	"IF",
	"IGNORE",
	"IGNORE_SERVER_IDS",
	"IMPORT",
	"IN",
	"INDEX",
	"INDEXES",
	"INFILE",
	"INITIAL_SIZE",
	"INNER",
	"INNODB",
	"INSERT",
	"INSERT_METHOD",
	"INSTALL",
	"INSTANCE",
	"INT1",
	"INT2",
	"INT3",
	"INT4",
	"INT8",
	"INTEGER",
	"INTERVAL",
	"INTO",
	"IO_THREAD",
	"ISOLATION",
	"ISSUER",
	"ITERATE",
	"JOIN",
	"JSON",
	"KEY",
	"KEYS",
	"KEY_BLOCK_SIZE",
	"KILL",
	"LAST",
	"LEAVE",
	"LEAVES",
	"LEFT",
	"LEVEL",
	"LIKE",
	"LIMIT",
	"LINES",
	"LOAD",
	"LOCAL",
	"LOCK",
	"LOGFILE",
	"LOGS",
	"LONG",
	"LONGBINARY",
	"LOOP",
	"LOW_PRIORITY",
	"MASTER",
	"MASTER_AUTO_POSITION",
	"MASTER_BIND",
	"MASTER_CONNECT_RETRY",
	"MASTER_HEARTBEAT_PERIOD",
	"MASTER_HOST",
	"MASTER_LOG_FILE",
	"MASTER_LOG_POS",
	"MASTER_PASSWORD",
	"MASTER_PORT",
	"MASTER_RETRY_COUNT",
	"MASTER_SSL",
	"MASTER_SSL_CA",
	"MASTER_SSL_CERT",
	"MASTER_SSL_CIPHER",
	"MASTER_SSL_CRL",
	"MASTER_SSL_CRLPATH",
	"MASTER_SSL_KEY",
	"MASTER_SSL_VERIFY_SERVER_CERT",
	"MASTER_TLS_VERSION",
	"MASTER_USER",
	"MAX_CONNECTIONS_PER_HOUR",
	"MAX_QUERIES_PER_HOUR",
	"MAX_ROWS",
	"MAX_SIZE",
	"MAX_UPDATES_PER_HOUR",
	"MAX_USER_CONNECTIONS",
	"MEDIUM",
	"MEMORY",
	"MERGE",
	"MESSAGE_TEXT",
	"MIDDLEINT",
	"MIN_ROWS",
	"MODE",
	"MODIFY",
	"MRG_MYISAM",
	"MUTEX",
	"MYISAM",
	"MYSQL_ERRNO",
	"NAME",
	"NAMES",
	"NATIONAL",
	"NATURAL",
	"NCHAR",
	"NDB",
	"NDBCLUSTER",
	"NEVER",
	"NEXT",
	"NO",
	"NODEGROUP",
	"NONE",
	"NOT",
	"NO_WRITE_TO_BINLOG",
	"NULL",
	"NUMBER",
	"NUMERIC",
	"NVARCHAR",
	"OFFSET",
	"ON",
	"ONLY",
	"OPEN",
	"OPTIMIZE",
	"OPTIMIZER_COSTS",
	"OPTION",
	"OPTIONALLY",
	"OPTIONS",
	"ORDER",
	"OUTER",
	"OUTFILE",
	"OWNER",
	"PACK_KEYS",
	"PARSER",
	"PARTIAL",
	"PARTITION",
	"PARTITIONING",
	"PARTITIONS",
	"PASSWORD",
	"PLUGIN",
	"PLUGINS",
	"PLUGIN_DIR",
	"PORT",
	"PRECISION",
	"PREPARE",
	"PRESERVE",
	"PREV",
	"PRIMARY",
	"PRIVILEGES",
	"PROCEDURE",
	"PROCESS",
	"PROCESSLIST",
	"PROFILE",
	"PROFILES",
	"PROXY",
	"PURGE",
	"QUERY",
	"QUICK",
	"READ",
	"REAL",
	"REBUILD",
	"RECOVER",
	"REDUNDANT",
	"REFERENCES",
	"RELAY",
	"RELAYLOG",
	"RELAY_LOG_FILE",
	"RELAY_LOG_POS",
	"RELEASE",
	"RELOAD",
	"REMOVE",
	"RENAME",
	"REORGANIZE",
	"REPAIR",
	"REPEAT",
	"REPEATABLE",
	"REPLACE",
	"REPLICATE_DO_DB",
	"REPLICATE_DO_TABLE",
	"REPLICATE_IGNORE_DB",
	"REPLICATE_IGNORE_TABLE",
	"REPLICATE_REWRITE_DB",
	"REPLICATE_WILD_DO_TABLE",
	"REPLICATE_WILD_IGNORE_TABLE",
	"REPLICATION",
	"REQUIRE",
	"RESET",
	"RESIGNAL",
	"RESTRICT",
	"RETURN",
	"RETURNED_SQLSTATE",
	"RETURNS",
	"REVOKE",
	"RIGHT",
	"ROLLBACK",
	"ROWS",
	"ROW_COUNT",
	"ROW_FORMAT",
	"SAVEPOINT",
	"SCHEDULE",
	"SCHEMA",
	"SCHEMAS",
	"SCHEMA_NAME",
	"SECURITY",
	"SELECT",
	"SERIAL",
	"SERIALIZABLE",
	"SERVER",
	"SESSION",
	"SET",
	"SHARE",
	"SHOW",
	"SHUTDOWN",
	"SIGNAL",
	"SLAVE",
	"SLOW",
	"SNAPSHOT",
	"SOCKET",
	"SONAME",
	"SPATIAL",
	"SQLSTATE",
	"SQL_AFTER_GTIDS",
	"SQL_AFTER_MTS_GAPS",
	"SQL_BEFORE_GTIDS",
	"SQL_BIG_RESULT",
	"SQL_BUFFER_RESULT",
	"SQL_CACHE",
	"SQL_CALC_FOUND_ROWS",
	"SQL_LOG_BIN",
	"SQL_NO_CACHE",
	"SQL_SLAVE_SKIP_COUNTER",
	"SQL_SMALL_RESULT",
	"SQL_THREAD",
	"SSL",
	"START",
	"STARTING",
	"STARTS",
	"STATS_AUTO_RECALC",
	"STATS_PERSISTENT",
	"STATS_SAMPLE_PAGES",
	"STATUS",
	"STOP",
	"STORAGE",
	"STORED",
	"STRAIGHT_JOIN",
	"STRING",
	"SUBCLASS_ORIGIN",
	"SUBJECT",
	"SUPER",
	"TABLE",
	"TABLES",
	"TABLESPACE",
	"TABLE_NAME",
	"TEMPORARY",
	"TERMINATED",
	"THEN",
	"TIME",
	"TIMESTAMP",
	"TO",
	"TRADITIONAL",
	"TRANSACTION",
	"TRIGGER",
	"TRIGGERS",
	"TRUE",
	"TRUNCATE",
	"TYPE",
	"UNCOMMITTED",
	"UNDO",
	"UNINSTALL",
	"UNION",
	"UNIQUE",
	"UNLOCK",
	"UNSIGNED",
	"UNTIL",
	"UPDATE",
	"UPGRADE",
	"USAGE",
	"USE",
	"USER",
	"USER_RESOURCES",
	"USE_FRM",
	"USING",
	"VALUE",
	"VALUES",
	"VARCHARACTER",
	"VARIABLE",
	"VARIABLES",
	"VARYING",
	"VIEW",
	"VIRTUAL",
	"WAIT",
	"WARNINGS",
	"WHEN",
	"WHERE",
	"WHILE",
	"WITH",
	"WORK",
	"WRAPPER",
	"WRITE",
	"X509",
	"XA",
	"ZEROFILL",
}

var mysql57function = []string{
	"AES_DECRYPT",
	"AES_ENCRYPT",
	"AGAINST",
	"AND",
	"ANY_VALUE",
	"AREA",
	"ASBINARY",
	"ASC",
	"ASTEXT",
	"ASWKB",
	"ASWKT",
	"BETWEEN",
	"BIGINT",
	"BINARY",
	"BOOLEAN",
	"BOTH",
	"BUFFER",
	"BY",
	"CASE",
	"CEIL",
	"CEILING",
	"CENTROID",
	"CHAR",
	"CONTAINS",
	"CONVERT",
	"CONVEXHULL",
	"COUNT",
	"CREATE_DH_PARAMETERS",
	"CROSSES",
	"DATE",
	"DATETIME",
	"DATE_ADD",
	"DATE_SUB",
	"DAY",
	"DAY_HOUR",
	"DAY_MINUTE",
	"DAY_SECOND",
	"DECIMAL",
	"DESC",
	"DIMENSION",
	"DISJOINT",
	"DISTANCE",
	"DISTINCT",
	"ELSE",
	"END",
	"ENDPOINT",
	"ENVELOPE",
	"EQUALS",
	"ESCAPE",
	"EXPANSION",
	"EXTERIORRING",
	"EXTRACTION)",
	"FLOOR",
	"FROM",
	"GEOMCOLLFROMTEXT",
	"GEOMCOLLFROMWKB",
	"GEOMETRYCOLLECTION",
	"GEOMETRYCOLLECTIONFROMTEXT",
	"GEOMETRYCOLLECTIONFROMWKB",
	"GEOMETRYFROMTEXT",
	"GEOMETRYFROMWKB",
	"GEOMETRYN",
	"GEOMETRYTYPE",
	"GEOMFROMTEXT",
	"GEOMFROMWKB",
	"GLENGTH",
	"HOUR",
	"HOUR_MINUTE",
	"HOUR_SECOND",
	"IF",
	"IN",
	"INLINE",
	"INSERT",
	"INTEGER",
	"INTERIORRINGN",
	"INTERSECTS",
	"INTERVAL",
	"IS",
	"ISCLOSED",
	"ISEMPTY",
	"ISSIMPLE",
	"JSON",
	"JSON_APPEND",
	"JSON_ARRAY",
	"JSON_ARRAYAGG",
	"JSON_ARRAY_APPEND",
	"JSON_ARRAY_INSERT",
	"JSON_CONTAINS",
	"JSON_CONTAINS_PATH",
	"JSON_DEPTH",
	"JSON_EXTRACT",
	"JSON_INSERT",
	"JSON_KEYS",
	"JSON_LENGTH",
	"JSON_MERGE",
	"JSON_MERGE_PATCH",
	"JSON_MERGE_PRESERVE",
	"JSON_OBJECT",
	"JSON_OBJECTAGG",
	"JSON_PRETTY",
	"JSON_QUOTE",
	"JSON_REMOVE",
	"JSON_REPLACE",
	"JSON_SEARCH",
	"JSON_SET",
	"JSON_STORAGE_SIZE",
	"JSON_TYPE",
	"JSON_UNQUOTE",
	"JSON_VALID",
	"LEADING",
	"LIKE",
	"LINEFROMTEXT",
	"LINEFROMWKB",
	"LINESTRING",
	"LINESTRINGFROMTEXT",
	"LINESTRINGFROMWKB",
	"MATCH",
	"MBRCONTAINS",
	"MBRDISJOINT",
	"MBREQUAL",
	"MBRINTERSECTS",
	"MBROVERLAPS",
	"MBRTOUCHES",
	"MBRWITHIN",
	"MINUTE",
	"MINUTE_SECOND",
	"MLINEFROMTEXT",
	"MLINEFROMWKB",
	"MOD",
	"MODE",
	"MONTH",
	"MPOINTFROMTEXT",
	"MPOINTFROMWKB",
	"MPOLYFROMTEXT",
	"MPOLYFROMWKB",
	"MULTILINESTRING",
	"MULTILINESTRINGFROMTEXT",
	"MULTILINESTRINGFROMWKB",
	"MULTIPOINT",
	"MULTIPOINTFROMTEXT",
	"MULTIPOINTFROMWKB",
	"MULTIPOLYGON",
	"MULTIPOLYGONFROMTEXT",
	"MULTIPOLYGONFROMWKB",
	"NOT",
	"NULL",
	"NUMGEOMETRIES",
	"NUMINTERIORRINGS",
	"NUMPOINTS",
	"OR",
	"ORDER",
	"OVERLAPS",
	"PATH)",
	"POINT",
	"POINTFROMTEXT",
	"POINTFROMWKB",
	"POINTN",
	"POLYFROMTEXT",
	"POLYFROMWKB",
	"POLYGON",
	"POLYGONFROMTEXT",
	"POLYGONFROMWKB",
	"POW",
	"POWER",
	"QUERY",
	"RANDOM_BYTES",
	"RLIKE",
	"SECOND",
	"SEPARATOR",
	"SHA",
	"SHA1",
	"SHA2",
	"SIGNED",
	"SOUNDS",
	"SRID",
	"STARTPOINT",
	"STD",
	"STDDEV",
	"ST_AREA",
	"ST_ASBINARY",
	"ST_ASGEOJSON",
	"ST_ASTEXT",
	"ST_ASWKB",
	"ST_ASWKT",
	"ST_BUFFER",
	"ST_BUFFER_STRATEGY",
	"ST_CENTROID",
	"ST_CONTAINS",
	"ST_CONVEXHULL",
	"ST_CROSSES",
	"ST_DIFFERENCE",
	"ST_DIMENSION",
	"ST_DISJOINT",
	"ST_DISTANCE",
	"ST_DISTANCE_SPHERE",
	"ST_ENDPOINT",
	"ST_ENVELOPE",
	"ST_EQUALS",
	"ST_EXTERIORRING",
	"ST_GEOHASH",
	"ST_GEOMCOLLFROMTEXT",
	"ST_GEOMCOLLFROMWKB",
	"ST_GEOMETRYCOLLECTIONFROMTEXT",
	"ST_GEOMETRYCOLLECTIONFROMWKB",
	"ST_GEOMETRYFROMTEXT",
	"ST_GEOMETRYFROMWKB",
	"ST_GEOMETRYN",
	"ST_GEOMETRYTYPE",
	"ST_GEOMFROMGEOJSON",
	"ST_GEOMFROMTEXT",
	"ST_GEOMFROMWKB",
	"ST_INTERIORRINGN",
	"ST_INTERSECTION",
	"ST_INTERSECTS",
	"ST_ISCLOSED",
	"ST_ISEMPTY",
	"ST_ISSIMPLE",
	"ST_ISVALID",
	"ST_LATFROMGEOHASH",
	"ST_LINEFROMTEXT",
	"ST_LINEFROMWKB",
	"ST_LINESTRINGFROMTEXT",
	"ST_LINESTRINGFROMWKB",
	"ST_LONGFROMGEOHASH",
	"ST_MAKEENVELOPE",
	"ST_MLINEFROMTEXT",
	"ST_MLINEFROMWKB",
	"ST_MPOINTFROMTEXT",
	"ST_MPOINTFROMWKB",
	"ST_MPOLYFROMTEXT",
	"ST_MPOLYFROMWKB",
	"ST_MULTILINESTRINGFROMTEXT",
	"ST_MULTILINESTRINGFROMWKB",
	"ST_MULTIPOINTFROMTEXT",
	"ST_MULTIPOINTFROMWKB",
	"ST_MULTIPOLYGONFROMTEXT",
	"ST_MULTIPOLYGONFROMWKB",
	"ST_NUMGEOMETRIES",
	"ST_NUMINTERIORRING",
	"ST_NUMINTERIORRINGS",
	"ST_NUMPOINTS",
	"ST_OVERLAPS",
	"ST_POINTFROMGEOHASH",
	"ST_POINTFROMTEXT",
	"ST_POINTFROMWKB",
	"ST_POINTN",
	"ST_POLYFROMTEXT",
	"ST_POLYFROMWKB",
	"ST_POLYGONFROMTEXT",
	"ST_POLYGONFROMWKB",
	"ST_SIMPLIFY",
	"ST_SRID",
	"ST_STARTPOINT",
	"ST_SYMDIFFERENCE",
	"ST_TOUCHES",
	"ST_UNION",
	"ST_VALIDATE",
	"ST_WITHIN",
	"ST_X",
	"ST_Y",
	"THEN",
	"TIME",
	"TIMESTAMP",
	"TOUCHES",
	"TRAILING",
	"UNSIGNED",
	"WHEN",
	"WITH",
	"WITHIN",
	"X",
	"Y",
	"YEAR",
	"YEAR_MONTH",
}

var mysql56Keyword = []string{
	"ACTION",
	"ADD",
	"AFTER",
	"AGGREGATE",
	"ALGORITHM",
	"ALL",
	"ALTER",
	"ANALYSE",
	"ANALYZE",
	"ARCHIVE",
	"AS",
	"ASC",
	"AT",
	"AUTHORS",
	"AUTOCOMMIT",
	"AUTOEXTEND_SIZE",
	"AUTO_INCREMENT",
	"AVG_ROW_LENGTH",
	"BEFORE",
	"BEGIN",
	"BINARY",
	"BINLOG",
	"BOOL",
	"BOOLEAN",
	"BTREE",
	"BY",
	"BYTE",
	"CACHE",
	"CALL",
	"CASCADE",
	"CASE",
	"CATALOG_NAME",
	"CHAIN",
	"CHANGE",
	"CHAR",
	"CHARACTER",
	"CHARSET",
	"CHECK",
	"CHECKSUM",
	"CIPHER",
	"CLASS_ORIGIN",
	"CLIENT",
	"CLOSE",
	"COALESCE",
	"CODE",
	"COLLATE",
	"COLLATION",
	"COLUMN",
	"COLUMNS",
	"COLUMN_NAME",
	"COMMENT",
	"COMMIT",
	"COMMITTED",
	"COMPACT",
	"COMPLETION",
	"COMPRESSED",
	"CONCURRENT",
	"CONDITION",
	"CONNECTION",
	"CONSISTENT",
	"CONSTRAINT",
	"CONSTRAINT_CATALOG",
	"CONSTRAINT_NAME",
	"CONSTRAINT_SCHEMA",
	"CONTINUE",
	"CONTRIBUTORS",
	"COUNT",
	"CREATE",
	"CROSS",
	"CSV",
	"CURRENT_USER",
	"CURSOR",
	"CURSOR_NAME",
	"DATA",
	"DATABASE",
	"DATABASES",
	"DATAFILE",
	"DATE",
	"DEALLOCATE",
	"DEC",
	"DECIMAL",
	"DECLARE",
	"DEFAULT",
	"DEFAULT_AUTH",
	"DEFINER",
	"DELAYED",
	"DELAY_KEY_WRITE",
	"DELETE",
	"DESC",
	"DESCRIBE",
	"DES_KEY_FILE",
	"DIAGNOSTICS",
	"DIRECTORY",
	"DISABLE",
	"DISCARD",
	"DISTINCT",
	"DISTINCTROW",
	"DO",
	"DROP",
	"DUAL",
	"DUMPFILE",
	"DUPLICATE",
	"DYNAMIC",
	"ELSE",
	"ELSEIF",
	"ENABLE",
	"ENCLOSED",
	"END",
	"ENDS",
	"ENGINE",
	"ENGINES",
	"ERROR",
	"ERRORS",
	"ESCAPED",
	"EVENT",
	"EVENTS",
	"EVERY",
	"EXCHANGE",
	"EXECUTE",
	"EXISTS",
	"EXIT",
	"EXPIRE",
	"EXPLAIN",
	"EXPORT",
	"EXTENDED",
	"EXTENT_SIZE",
	"FALSE",
	"FAST",
	"FEDERATED",
	"FETCH",
	"FIELDS",
	"FILE",
	"FIRST",
	"FIXED",
	"FLOAT4",
	"FLOAT8",
	"FLUSH",
	"FOR",
	"FORCE",
	"FOREIGN",
	"FORMAT",
	"FROM",
	"FULL",
	"FULLTEXT",
	"FUNCTION",
	"GENERAL",
	"GET",
	"GLOBAL",
	"GRANT",
	"GRANTS",
	"GROUP",
	"HANDLER",
	"HAVING",
	"HEAP",
	"HELP",
	"HELP_DATE",
	"HELP_VERSION",
	"HIGH_PRIORITY",
	"HOST",
	"HOSTS",
	"IDENTIFIED",
	"IF",
	"IGNORE",
	"IGNORE_SERVER_IDS",
	"IMPORT",
	"IN",
	"INDEX",
	"INDEXES",
	"INFILE",
	"INITIAL_SIZE",
	"INNER",
	"INNODB",
	"INSERT",
	"INSERT_METHOD",
	"INSTALL",
	"INT1",
	"INT2",
	"INT3",
	"INT4",
	"INT8",
	"INTEGER",
	"INTERVAL",
	"INTO",
	"IO_THREAD",
	"ISOLATION",
	"ISSUER",
	"ITERATE",
	"JOIN",
	"JSON",
	"KEY",
	"KEYS",
	"KEY_BLOCK_SIZE",
	"KILL",
	"LAST",
	"LEAVE",
	"LEAVES",
	"LEFT",
	"LEVEL",
	"LIKE",
	"LIMIT",
	"LINES",
	"LOAD",
	"LOCAL",
	"LOCK",
	"LOGFILE",
	"LOGS",
	"LONG",
	"LONGBINARY",
	"LOOP",
	"LOW_PRIORITY",
	"MASTER",
	"MASTER_AUTO_POSITION",
	"MASTER_BIND",
	"MASTER_CONNECT_RETRY",
	"MASTER_HEARTBEAT_PERIOD",
	"MASTER_HOST",
	"MASTER_LOG_FILE",
	"MASTER_LOG_POS",
	"MASTER_PASSWORD",
	"MASTER_PORT",
	"MASTER_RETRY_COUNT",
	"MASTER_SSL",
	"MASTER_SSL_CA",
	"MASTER_SSL_CERT",
	"MASTER_SSL_CIPHER",
	"MASTER_SSL_CRL",
	"MASTER_SSL_CRLPATH",
	"MASTER_SSL_KEY",
	"MASTER_SSL_VERIFY_SERVER_CERT",
	"MASTER_USER",
	"MAX_CONNECTIONS_PER_HOUR",
	"MAX_QUERIES_PER_HOUR",
	"MAX_ROWS",
	"MAX_SIZE",
	"MAX_UPDATES_PER_HOUR",
	"MAX_USER_CONNECTIONS",
	"MEDIUM",
	"MEMORY",
	"MERGE",
	"MESSAGE_TEXT",
	"MIDDLEINT",
	"MIN_ROWS",
	"MODE",
	"MODIFY",
	"MRG_MYISAM",
	"MUTEX",
	"MYISAM",
	"MYSQL_ERRNO",
	"NAME",
	"NAMES",
	"NATIONAL",
	"NATURAL",
	"NCHAR",
	"NDB",
	"NDBCLUSTER",
	"NEXT",
	"NO",
	"NODEGROUP",
	"NONE",
	"NOT",
	"NO_WRITE_TO_BINLOG",
	"NULL",
	"NUMBER",
	"NUMERIC",
	"NVARCHAR",
	"OFFSET",
	"ON",
	"ONLY",
	"OPEN",
	"OPTIMIZE",
	"OPTION",
	"OPTIONALLY",
	"OPTIONS",
	"ORDER",
	"OUTER",
	"OUTFILE",
	"OWNER",
	"PACK_KEYS",
	"PARSER",
	"PARTIAL",
	"PARTITION",
	"PARTITIONS",
	"PASSWORD",
	"PLUGIN",
	"PLUGINS",
	"PLUGIN_DIR",
	"PORT",
	"PRECISION",
	"PREPARE",
	"PRESERVE",
	"PREV",
	"PRIMARY",
	"PRIVILEGES",
	"PROCEDURE",
	"PROCESS",
	"PROCESSLIST",
	"PROFILE",
	"PROFILES",
	"PROXY",
	"PURGE",
	"QUERY",
	"QUICK",
	"READ",
	"REAL",
	"REBUILD",
	"RECOVER",
	"REDUNDANT",
	"REFERENCES",
	"RELAY",
	"RELAYLOG",
	"RELAY_LOG_FILE",
	"RELAY_LOG_POS",
	"RELEASE",
	"RELOAD",
	"REMOVE",
	"RENAME",
	"REORGANIZE",
	"REPAIR",
	"REPEAT",
	"REPEATABLE",
	"REPLACE",
	"REPLICATION",
	"REQUIRE",
	"RESET",
	"RESIGNAL",
	"RESTRICT",
	"RETURN",
	"RETURNED_SQLSTATE",
	"RETURNS",
	"REVOKE",
	"RIGHT",
	"ROLLBACK",
	"ROWS",
	"ROW_COUNT",
	"ROW_FORMAT",
	"SAVEPOINT",
	"SCHEDULE",
	"SCHEMA",
	"SCHEMAS",
	"SCHEMA_NAME",
	"SECURITY",
	"SELECT",
	"SERIAL",
	"SERIALIZABLE",
	"SERVER",
	"SESSION",
	"SET",
	"SHARE",
	"SHOW",
	"SHUTDOWN",
	"SIGNAL",
	"SLAVE",
	"SLOW",
	"SNAPSHOT",
	"SOCKET",
	"SONAME",
	"SPATIAL",
	"SQLSTATE",
	"SQL_AFTER_GTIDS",
	"SQL_AFTER_MTS_GAPS",
	"SQL_BEFORE_GTIDS",
	"SQL_BIG_RESULT",
	"SQL_BUFFER_RESULT",
	"SQL_CACHE",
	"SQL_CALC_FOUND_ROWS",
	"SQL_LOG_BIN",
	"SQL_NO_CACHE",
	"SQL_SLAVE_SKIP_COUNTER",
	"SQL_SMALL_RESULT",
	"SQL_THREAD",
	"SSL",
	"START",
	"STARTING",
	"STARTS",
	"STATS_AUTO_RECALC",
	"STATS_PERSISTENT",
	"STATS_SAMPLE_PAGES",
	"STATUS",
	"STOP",
	"STORAGE",
	"STRAIGHT_JOIN",
	"STRING",
	"SUBCLASS_ORIGIN",
	"SUBJECT",
	"SUPER",
	"TABLE",
	"TABLES",
	"TABLESPACE",
	"TABLE_NAME",
	"TEMPORARY",
	"TERMINATED",
	"THEN",
	"TIME",
	"TIMESTAMP",
	"TO",
	"TRADITIONAL",
	"TRANSACTION",
	"TRIGGER",
	"TRIGGERS",
	"TRUE",
	"TRUNCATE",
	"TYPE",
	"UNCOMMITTED",
	"UNDO",
	"UNINSTALL",
	"UNION",
	"UNIQUE",
	"UNLOCK",
	"UNSIGNED",
	"UNTIL",
	"UPDATE",
	"UPGRADE",
	"USAGE",
	"USE",
	"USER",
	"USER_RESOURCES",
	"USE_FRM",
	"USING",
	"VALUE",
	"VALUES",
	"VARCHARACTER",
	"VARIABLE",
	"VARIABLES",
	"VARYING",
	"VIEW",
	"WAIT",
	"WARNINGS",
	"WHEN",
	"WHERE",
	"WHILE",
	"WITH",
	"WORK",
	"WRAPPER",
	"WRITE",
	"X509",
	"XA",
	"ZEROFILL",
}

var mysql56Function = []string{
	"%",
	"&",
	"(JSON",
	"*",
	"+",
	"-",
	"->",
	"->>",
	"/",
	":=",
	"<",
	"<<",
	"<=",
	"<=>",
	"<>",
	"=",
	">",
	">=",
	">>",
	"ABS",
	"ACOS",
	"ADDDATE",
	"ADDTIME",
	"AES_DECRYPT",
	"AES_ENCRYPT",
	"AGAINST",
	"AND",
	"ANY_VALUE",
	"ARRAY",
	"ASC",
	"ASCII",
	"ASIN",
	"ASYMMETRIC_DECRYPT",
	"ASYMMETRIC_DERIVE",
	"ASYMMETRIC_ENCRYPT",
	"ASYMMETRIC_SIGN",
	"ASYMMETRIC_VERIFY",
	"ATAN",
	"ATAN2",
	"AVG",
	"BENCHMARK",
	"BETWEEN",
	"BIN",
	"BINARY",
	"BIN_TO_UUID",
	"BIT_AND",
	"BIT_COUNT",
	"BIT_LENGTH",
	"BIT_OR",
	"BIT_XOR",
	"BOOLEAN",
	"BOTH",
	"BY",
	"CAN_ACCESS_COLUMN",
	"CAN_ACCESS_DATABASE",
	"CAN_ACCESS_TABLE",
	"CAN_ACCESS_USER",
	"CAN_ACCESS_VIEW",
	"CASE",
	"CAST",
	"CEIL",
	"CEILING",
	"CHAR",
	"CHARACTER_LENGTH",
	"CHARSET",
	"CHAR_LENGTH",
	"COALESCE",
	"COERCIBILITY",
	"COLLATION",
	"COMPRESS",
	"CONCAT",
	"CONCAT_WS",
	"CONNECTION_ID",
	"CONV",
	"CONVERT",
	"CONVERT_TZ",
	"COS",
	"COT",
	"COUNT",
	"CRC32",
	"CREATE_ASYMMETRIC_PRIV_KEY",
	"CREATE_ASYMMETRIC_PUB_KEY",
	"CREATE_DH_PARAMETERS",
	"CREATE_DIGEST",
	"CUME_DIST",
	"CURDATE",
	"CURRENT_DATE",
	"CURRENT_ROLE",
	"CURRENT_TIME",
	"CURRENT_TIMESTAMP",
	"CURRENT_USER",
	"CURTIME",
	"DATABASE",
	"DATE",
	"DATEDIFF",
	"DATETIME",
	"DATE_ADD",
	"DATE_FORMAT",
	"DATE_SUB",
	"DAY",
	"DAYNAME",
	"DAYOFMONTH",
	"DAYOFWEEK",
	"DAYOFYEAR",
	"DAY_HOUR",
	"DAY_MINUTE",
	"DAY_SECOND",
	"DECIMAL",
	"DEFAULT",
	"DEGREES",
	"DENSE_RANK",
	"DESC",
	"DISTINCT",
	"DIV",
	"ELSE",
	"ELT",
	"END",
	"ESCAPE",
	"EXP",
	"EXPANSION",
	"EXPORT_SET",
	"EXTRACT",
	"EXTRACTION)",
	"EXTRACTVALUE",
	"FIELD",
	"FIND_IN_SET",
	"FIRST_VALUE",
	"FLOOR",
	"FORMAT",
	"FORMAT_BYTES",
	"FORMAT_PICO_TIME",
	"FOUND_ROWS",
	"FROM",
	"FROM_BASE64",
	"FROM_DAYS",
	"FROM_UNIXTIME",
	"FUNCTION",
	"GEOMCOLLECTION",
	"GEOMETRYCOLLECTION",
	"GET_DD_COLUMN_PRIVILEGES",
	"GET_DD_CREATE_OPTIONS",
	"GET_DD_INDEX_SUB_PART_LENGTH",
	"GET_FORMAT",
	"GET_LOCK",
	"GREATEST",
	"GROUPING",
	"GROUP_CONCAT",
	"GTID_SUBSET",
	"GTID_SUBTRACT",
	"HEX",
	"HOUR",
	"HOUR_MINUTE",
	"HOUR_SECOND",
	"ICU_VERSION",
	"IF",
	"IFNULL",
	"IN",
	"INET6_ATON",
	"INET6_NTOA",
	"INET_ATON",
	"INET_NTOA",
	"INLINE",
	"INSERT",
	"INSTR",
	"INTEGER",
	"INTERNAL_AUTO_INCREMENT",
	"INTERNAL_AVG_ROW_LENGTH",
	"INTERNAL_CHECKSUM",
	"INTERNAL_CHECK_TIME",
	"INTERNAL_DATA_FREE",
	"INTERNAL_DATA_LENGTH",
	"INTERNAL_DD_CHAR_LENGTH",
	"INTERNAL_GET_COMMENT_OR_ERROR",
	"INTERNAL_GET_ENABLED_ROLE_JSON",
	"INTERNAL_GET_HOSTNAME",
	"INTERNAL_GET_USERNAME",
	"INTERNAL_GET_VIEW_WARNING_OR_ERROR",
	"INTERNAL_INDEX_COLUMN_CARDINALITY",
	"INTERNAL_INDEX_LENGTH",
	"INTERNAL_IS_ENABLED_ROLE",
	"INTERNAL_IS_MANDATORY_ROLE",
	"INTERNAL_KEYS_DISABLED",
	"INTERNAL_MAX_DATA_LENGTH",
	"INTERNAL_TABLE_ROWS",
	"INTERNAL_UPDATE_TIME",
	"INTERVAL",
	"IS",
	"ISNULL",
	"IS_FREE_LOCK",
	"IS_IPV4",
	"IS_IPV4_COMPAT",
	"IS_IPV4_MAPPED",
	"IS_IPV6",
	"IS_USED_LOCK",
	"IS_UUID",
	"IS_VISIBLE_DD_OBJECT",
	"JSON",
	"JSON_ARRAY",
	"JSON_ARRAYAGG",
	"JSON_ARRAY_APPEND",
	"JSON_ARRAY_INSERT",
	"JSON_CONTAINS",
	"JSON_CONTAINS_PATH",
	"JSON_DEPTH",
	"JSON_EXTRACT",
	"JSON_INSERT",
	"JSON_KEYS",
	"JSON_LENGTH",
	"JSON_MERGE",
	"JSON_MERGE_PATCH",
	"JSON_MERGE_PRESERVE",
	"JSON_OBJECT",
	"JSON_OBJECTAGG",
	"JSON_OVERLAPS",
	"JSON_PRETTY",
	"JSON_QUOTE",
	"JSON_REMOVE",
	"JSON_REPLACE",
	"JSON_SCHEMA_VALID",
	"JSON_SCHEMA_VALIDATION_REPORT",
	"JSON_SEARCH",
	"JSON_SET",
	"JSON_STORAGE_FREE",
	"JSON_STORAGE_SIZE",
	"JSON_TABLE",
	"JSON_TYPE",
	"JSON_UNQUOTE",
	"JSON_VALID",
	"JSON_VALUE",
	"LAG",
	"LAST_DAY",
	"LAST_INSERT_ID",
	"LAST_VALUE",
	"LCASE",
	"LEAD",
	"LEADING",
	"LEAST",
	"LEFT",
	"LENGTH",
	"LIKE",
	"LINESTRING",
	"LN",
	"LOAD_FILE",
	"LOCALTIME",
	"LOCALTIMESTAMP",
	"LOCATE",
	"LOG",
	"LOG10",
	"LOG2",
	"LOWER",
	"LPAD",
	"LTRIM",
	"MAKEDATE",
	"MAKETIME",
	"MAKE_SET",
	"MASTER_POS_WAIT",
	"MATCH",
	"MAX",
	"MBRCONTAINS",
	"MBRCOVEREDBY",
	"MBRCOVERS",
	"MBRDISJOINT",
	"MBREQUALS",
	"MBRINTERSECTS",
	"MBROVERLAPS",
	"MBRTOUCHES",
	"MBRWITHIN",
	"MD5",
	"MEMBER",
	"MICROSECOND",
	"MID",
	"MIN",
	"MINUTE",
	"MINUTE_SECOND",
	"MOD",
	"MODE",
	"MONTH",
	"MONTHNAME",
	"MULTILINESTRING",
	"MULTIPOINT",
	"MULTIPOLYGON",
	"NAME_CONST",
	"NOT",
	"NOW",
	"NTH_VALUE",
	"NTILE",
	"NULL",
	"NULLIF",
	"OCT",
	"OCTET_LENGTH",
	"OF",
	"OR",
	"ORD",
	"ORDER",
	"PATH)",
	"PERCENT_RANK",
	"PERIOD_ADD",
	"PERIOD_DIFF",
	"PI",
	"POINT",
	"POLYGON",
	"POSITION",
	"POW",
	"POWER",
	"PS_CURRENT_THREAD_ID",
	"PS_THREAD_ID",
	"QUARTER",
	"QUERY",
	"QUOTE",
	"RADIANS",
	"RAND",
	"RANDOM_BYTES",
	"RANK",
	"REGEXP",
	"REGEXP_INSTR",
	"REGEXP_LIKE",
	"REGEXP_REPLACE",
	"REGEXP_SUBSTR",
	"RELEASE_ALL_LOCKS",
	"RELEASE_LOCK",
	"REPEAT",
	"REPLACE",
	"REVERSE",
	"RIGHT",
	"RLIKE",
	"ROLES_GRAPHML",
	"ROUND",
	"ROW_COUNT",
	"ROW_NUMBER",
	"RPAD",
	"RTRIM",
	"SCHEMA",
	"SECOND",
	"SEC_TO_TIME",
	"SEPARATOR",
	"SESSION_USER",
	"SHA",
	"SHA1",
	"SHA2",
	"SIGN",
	"SIGNED",
	"SIN",
	"SLEEP",
	"SOUNDEX",
	"SOUNDS",
	"SPACE",
	"SQRT",
	"STATEMENT_DIGEST",
	"STATEMENT_DIGEST_TEXT",
	"STD",
	"STDDEV",
	"STDDEV_POP",
	"STDDEV_SAMP",
	"STRCMP",
	"STR_TO_DATE",
	"ST_AREA",
	"ST_ASBINARY",
	"ST_ASGEOJSON",
	"ST_ASTEXT",
	"ST_ASWKB",
	"ST_ASWKT",
	"ST_BUFFER",
	"ST_BUFFER_STRATEGY",
	"ST_CENTROID",
	"ST_CONTAINS",
	"ST_CONVEXHULL",
	"ST_CROSSES",
	"ST_DIFFERENCE",
	"ST_DIMENSION",
	"ST_DISJOINT",
	"ST_DISTANCE",
	"ST_DISTANCE_SPHERE",
	"ST_ENDPOINT",
	"ST_ENVELOPE",
	"ST_EQUALS",
	"ST_EXTERIORRING",
	"ST_FRECHETDISTANCE",
	"ST_GEOHASH",
	"ST_GEOMCOLLFROMTEXT",
	"ST_GEOMCOLLFROMWKB",
	"ST_GEOMETRYCOLLECTIONFROMTEXT",
	"ST_GEOMETRYCOLLECTIONFROMWKB",
	"ST_GEOMETRYFROMTEXT",
	"ST_GEOMETRYFROMWKB",
	"ST_GEOMETRYN",
	"ST_GEOMETRYTYPE",
	"ST_GEOMFROMGEOJSON",
	"ST_GEOMFROMTEXT",
	"ST_GEOMFROMWKB",
	"ST_HAUSDORFFDISTANCE",
	"ST_INTERIORRINGN",
	"ST_INTERSECTION",
	"ST_INTERSECTS",
	"ST_ISCLOSED",
	"ST_ISEMPTY",
	"ST_ISSIMPLE",
	"ST_ISVALID",
	"ST_LATFROMGEOHASH",
	"ST_LATITUDE",
	"ST_LENGTH",
	"ST_LINEFROMTEXT",
	"ST_LINEFROMWKB",
	"ST_LINEINTERPOLATEPOINT",
	"ST_LINEINTERPOLATEPOINTS",
	"ST_LINESTRINGFROMTEXT",
	"ST_LINESTRINGFROMWKB",
	"ST_LONGFROMGEOHASH",
	"ST_LONGITUDE",
	"ST_MAKEENVELOPE",
	"ST_MLINEFROMTEXT",
	"ST_MLINEFROMWKB",
	"ST_MPOINTFROMTEXT",
	"ST_MPOINTFROMWKB",
	"ST_MPOLYFROMTEXT",
	"ST_MPOLYFROMWKB",
	"ST_MULTILINESTRINGFROMTEXT",
	"ST_MULTILINESTRINGFROMWKB",
	"ST_MULTIPOINTFROMTEXT",
	"ST_MULTIPOINTFROMWKB",
	"ST_MULTIPOLYGONFROMTEXT",
	"ST_MULTIPOLYGONFROMWKB",
	"ST_NUMGEOMETRIES",
	"ST_NUMINTERIORRING",
	"ST_NUMINTERIORRINGS",
	"ST_NUMPOINTS",
	"ST_OVERLAPS",
	"ST_POINTATDISTANCE",
	"ST_POINTFROMGEOHASH",
	"ST_POINTFROMTEXT",
	"ST_POINTFROMWKB",
	"ST_POINTN",
	"ST_POLYFROMTEXT",
	"ST_POLYFROMWKB",
	"ST_POLYGONFROMTEXT",
	"ST_POLYGONFROMWKB",
	"ST_SIMPLIFY",
	"ST_SRID",
	"ST_STARTPOINT",
	"ST_SWAPXY",
	"ST_SYMDIFFERENCE",
	"ST_TOUCHES",
	"ST_TRANSFORM",
	"ST_UNION",
	"ST_VALIDATE",
	"ST_WITHIN",
	"ST_X",
	"ST_Y",
	"SUBDATE",
	"SUBSTR",
	"SUBSTRING",
	"SUBSTRING_INDEX",
	"SUBTIME",
	"SUM",
	"SYSDATE",
	"SYSTEM_USER",
	"TAN",
	"THEN",
	"TIME",
	"TIMEDIFF",
	"TIMESTAMP",
	"TIMESTAMPADD",
	"TIMESTAMPDIFF",
	"TIMEZONE",
	"TIME_FORMAT",
	"TIME_TO_SEC",
	"TO_BASE64",
	"TO_DAYS",
	"TO_SECONDS",
	"TRAILING",
	"TRIM",
	"TRUNCATE",
	"UCASE",
	"UNCOMPRESS",
	"UNCOMPRESSED_LENGTH",
	"UNHEX",
	"UNIX_TIMESTAMP",
	"UNSIGNED",
	"UPDATEXML",
	"UPPER",
	"USER",
	"UTC_DATE",
	"UTC_TIME",
	"UTC_TIMESTAMP",
	"UUID",
	"UUID_SHORT",
	"UUID_TO_BIN",
	"VALIDATE_PASSWORD_STRENGTH",
	"VALUES",
	"VARIANCE",
	"VAR_POP",
	"VAR_SAMP",
	"VERSION",
	"WAIT_FOR_EXECUTED_GTID_SET",
	"WAIT_UNTIL_SQL_THREAD_AFTER_GTIDS",
	"WEEK",
	"WEEKDAY",
	"WEEKOFYEAR",
	"WEIGHT_STRING",
	"WHEN",
	"WITH",
	"XOR",
	"YEAR",
	"YEARWEEK",
	"YEAR_MONTH",
	"^",
	"|",
	"~",
}