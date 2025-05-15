# Details

Date : 2025-05-14 11:16:25

Directory /Users/neilluo/workspaces/storage/godis

Total : 145 files,  21127 codes, 1456 comments, 2613 blanks, all 25196 lines

[Summary](results.md) / Details / [Diff Summary](diff.md) / [Diff Details](diff-details.md)

## Files
| filename | language | code | comment | blank | total |
| :--- | :--- | ---: | ---: | ---: | ---: |
| [.github/workflows/coverall.yml](/.github/workflows/coverall.yml) | YAML | 23 | 2 | 9 | 34 |
| [.travis.yml](/.travis.yml) | YAML | 8 | 0 | 2 | 10 |
| [README.md](/README.md) | Markdown | 136 | 0 | 41 | 177 |
| [README\_CN.md](/README_CN.md) | Markdown | 125 | 0 | 35 | 160 |
| [aof/aof.go](/aof/aof.go) | Go | 255 | 40 | 26 | 321 |
| [aof/marshal.go](/aof/marshal.go) | Go | 101 | 2 | 16 | 119 |
| [aof/rdb.go](/aof/rdb.go) | Go | 170 | 11 | 15 | 196 |
| [aof/rewrite.go](/aof/rewrite.go) | Go | 114 | 18 | 17 | 149 |
| [build-all.sh](/build-all.sh) | Shell Script | 5 | 1 | 2 | 8 |
| [build-darwin.sh](/build-darwin.sh) | Shell Script | 1 | 1 | 2 | 4 |
| [build-linux.sh](/build-linux.sh) | Shell Script | 1 | 1 | 1 | 3 |
| [cluster/cluster.go](/cluster/cluster.go) | Go | 34 | 2 | 5 | 41 |
| [cluster/commands/default.go](/cluster/commands/default.go) | Go | 92 | 0 | 4 | 96 |
| [cluster/commands/del.go](/cluster/commands/del.go) | Go | 60 | 2 | 6 | 68 |
| [cluster/commands/del\_test.go](/cluster/commands/del_test.go) | Go | 30 | 2 | 7 | 39 |
| [cluster/commands/mset.go](/cluster/commands/mset.go) | Go | 181 | 9 | 16 | 206 |
| [cluster/commands/mset\_test.go](/cluster/commands/mset_test.go) | Go | 35 | 2 | 5 | 42 |
| [cluster/commands/rename.go](/cluster/commands/rename.go) | Go | 143 | 9 | 15 | 167 |
| [cluster/commands/rename\_test.go](/cluster/commands/rename_test.go) | Go | 39 | 2 | 9 | 50 |
| [cluster/commands/tcc\_utils.go](/cluster/commands/tcc_utils.go) | Go | 52 | 5 | 10 | 67 |
| [cluster/core/command.go](/cluster/core/command.go) | Go | 56 | 7 | 11 | 74 |
| [cluster/core/connection.go](/cluster/core/connection.go) | Go | 153 | 7 | 20 | 180 |
| [cluster/core/connection\_inmem.go](/cluster/core/connection_inmem.go) | Go | 67 | 2 | 16 | 85 |
| [cluster/core/core.go](/cluster/core/core.go) | Go | 157 | 6 | 20 | 183 |
| [cluster/core/core\_test.go](/cluster/core/core_test.go) | Go | 212 | 12 | 15 | 239 |
| [cluster/core/cron.go](/cluster/core/cron.go) | Go | 32 | 1 | 4 | 37 |
| [cluster/core/migration.go](/cluster/core/migration.go) | Go | 260 | 19 | 23 | 302 |
| [cluster/core/node\_manager.go](/cluster/core/node_manager.go) | Go | 221 | 18 | 21 | 260 |
| [cluster/core/replica\_manager.go](/cluster/core/replica_manager.go) | Go | 126 | 11 | 19 | 156 |
| [cluster/core/tcc.go](/cluster/core/tcc.go) | Go | 119 | 17 | 26 | 162 |
| [cluster/core/tests.go](/cluster/core/tests.go) | Go | 37 | 3 | 4 | 44 |
| [cluster/core/utils.go](/cluster/core/utils.go) | Go | 72 | 11 | 16 | 99 |
| [cluster/raft/fsm.go](/cluster/raft/fsm.go) | Go | 171 | 20 | 20 | 211 |
| [cluster/raft/fsm\_utils.go](/cluster/raft/fsm_utils.go) | Go | 92 | 12 | 10 | 114 |
| [cluster/raft/raft.go](/cluster/raft/raft.go) | Go | 183 | 8 | 21 | 212 |
| [cluster/raft/utils.go](/cluster/raft/utils.go) | Go | 33 | 1 | 9 | 43 |
| [commands.md](/commands.md) | Markdown | 114 | 0 | 1 | 115 |
| [config/config.go](/config/config.go) | Go | 137 | 10 | 17 | 164 |
| [config/config\_test.go](/config/config_test.go) | Go | 25 | 0 | 3 | 28 |
| [database/aof\_test.go](/database/aof_test.go) | Go | 263 | 9 | 16 | 288 |
| [database/cluster\_helper.go](/database/cluster_helper.go) | Go | 117 | 13 | 10 | 140 |
| [database/cluster\_helper\_test.go](/database/cluster_helper_test.go) | Go | 51 | 1 | 6 | 58 |
| [database/commandinfo.go](/database/commandinfo.go) | Go | 115 | 2 | 10 | 127 |
| [database/commandinfo\_test.go](/database/commandinfo_test.go) | Go | 18 | 0 | 3 | 21 |
| [database/database.go](/database/database.go) | Go | 236 | 44 | 40 | 320 |
| [database/doc.go](/database/doc.go) | Go | 1 | 17 | 2 | 20 |
| [database/geo.go](/database/geo.go) | Go | 244 | 12 | 23 | 279 |
| [database/geo\_test.go](/database/geo_test.go) | Go | 95 | 0 | 8 | 103 |
| [database/hash.go](/database/hash.go) | Go | 488 | 35 | 63 | 586 |
| [database/hash\_test.go](/database/hash_test.go) | Go | 339 | 17 | 38 | 394 |
| [database/keys.go](/database/keys.go) | Go | 429 | 25 | 57 | 511 |
| [database/keys\_test.go](/database/keys_test.go) | Go | 347 | 11 | 40 | 398 |
| [database/list.go](/database/list.go) | Go | 599 | 41 | 76 | 716 |
| [database/list\_test.go](/database/list_test.go) | Go | 535 | 25 | 62 | 622 |
| [database/persistence.go](/database/persistence.go) | Go | 118 | 5 | 9 | 132 |
| [database/persistence\_test.go](/database/persistence_test.go) | Go | 79 | 1 | 7 | 87 |
| [database/replication\_master.go](/database/replication_master.go) | Go | 396 | 26 | 41 | 463 |
| [database/replication\_master\_test.go](/database/replication_master_test.go) | Go | 283 | 6 | 27 | 316 |
| [database/replication\_slave.go](/database/replication_slave.go) | Go | 395 | 33 | 39 | 467 |
| [database/replication\_slave\_test.go](/database/replication_slave_test.go) | Go | 225 | 11 | 15 | 251 |
| [database/router.go](/database/router.go) | Go | 86 | 6 | 12 | 104 |
| [database/server.go](/database/server.go) | Go | 365 | 42 | 42 | 449 |
| [database/set.go](/database/set.go) | Go | 378 | 20 | 37 | 435 |
| [database/set\_test.go](/database/set_test.go) | Go | 247 | 10 | 31 | 288 |
| [database/sortedset.go](/database/sortedset.go) | Go | 761 | 40 | 86 | 887 |
| [database/sortedset\_test.go](/database/sortedset_test.go) | Go | 554 | 107 | 150 | 811 |
| [database/string.go](/database/string.go) | Go | 796 | 28 | 64 | 888 |
| [database/string\_test.go](/database/string_test.go) | Go | 701 | 11 | 104 | 816 |
| [database/systemcmd.go](/database/systemcmd.go) | Go | 150 | 23 | 11 | 184 |
| [database/systemcmd\_test.go](/database/systemcmd_test.go) | Go | 69 | 0 | 8 | 77 |
| [database/transaction.go](/database/transaction.go) | Go | 152 | 14 | 14 | 180 |
| [database/transaction\_test.go](/database/transaction_test.go) | Go | 111 | 0 | 7 | 118 |
| [database/tx\_utils.go](/database/tx_utils.go) | Go | 158 | 2 | 15 | 175 |
| [database/tx\_utils\_test.go](/database/tx_utils_test.go) | Go | 166 | 5 | 16 | 187 |
| [database/util\_test.go](/database/util_test.go) | Go | 12 | 0 | 3 | 15 |
| [datastruct/bitmap/bitmap.go](/datastruct/bitmap/bitmap.go) | Go | 86 | 2 | 13 | 101 |
| [datastruct/bitmap/bitmap\_test.go](/datastruct/bitmap/bitmap_test.go) | Go | 159 | 3 | 7 | 169 |
| [datastruct/dict/concurrent.go](/datastruct/dict/concurrent.go) | Go | 420 | 17 | 52 | 489 |
| [datastruct/dict/concurrent\_test.go](/datastruct/dict/concurrent_test.go) | Go | 514 | 24 | 37 | 575 |
| [datastruct/dict/dict.go](/datastruct/dict/dict.go) | Go | 16 | 2 | 3 | 21 |
| [datastruct/dict/simple.go](/datastruct/dict/simple.go) | Go | 117 | 13 | 16 | 146 |
| [datastruct/dict/simple\_test.go](/datastruct/dict/simple_test.go) | Go | 78 | 0 | 5 | 83 |
| [datastruct/list/interface.go](/datastruct/list/interface.go) | Go | 18 | 3 | 4 | 25 |
| [datastruct/list/linked.go](/datastruct/list/linked.go) | Go | 246 | 22 | 24 | 292 |
| [datastruct/list/linked\_test.go](/datastruct/list/linked_test.go) | Go | 239 | 0 | 19 | 258 |
| [datastruct/list/quicklist.go](/datastruct/list/quicklist.go) | Go | 326 | 37 | 28 | 391 |
| [datastruct/list/quicklist\_test.go](/datastruct/list/quicklist_test.go) | Go | 244 | 1 | 17 | 262 |
| [datastruct/lock/lock\_map.go](/datastruct/lock/lock_map.go) | Go | 137 | 14 | 18 | 169 |
| [datastruct/set/set.go](/datastruct/set/set.go) | Go | 144 | 17 | 20 | 181 |
| [datastruct/set/set\_test.go](/datastruct/set/set_test.go) | Go | 57 | 0 | 4 | 61 |
| [datastruct/sortedset/border.go](/datastruct/sortedset/border.go) | Go | 161 | 15 | 25 | 201 |
| [datastruct/sortedset/skiplist.go](/datastruct/sortedset/skiplist.go) | Go | 270 | 44 | 33 | 347 |
| [datastruct/sortedset/skiplist\_test.go](/datastruct/sortedset/skiplist_test.go) | Go | 12 | 0 | 3 | 15 |
| [datastruct/sortedset/sortedset.go](/datastruct/sortedset/sortedset.go) | Go | 214 | 23 | 23 | 260 |
| [datastruct/sortedset/sortedset\_test.go](/datastruct/sortedset/sortedset_test.go) | Go | 42 | 0 | 6 | 48 |
| [go.mod](/go.mod) | Go Module File | 28 | 0 | 4 | 32 |
| [go.sum](/go.sum) | Go Checksum File | 184 | 0 | 1 | 185 |
| [interface/database/db.go](/interface/database/db.go) | Go | 31 | 6 | 8 | 45 |
| [interface/redis/conn.go](/interface/redis/conn.go) | Go | 27 | 2 | 9 | 38 |
| [interface/redis/reply.go](/interface/redis/reply.go) | Go | 4 | 1 | 2 | 7 |
| [interface/tcp/handler.go](/interface/tcp/handler.go) | Go | 10 | 2 | 4 | 16 |
| [lib/consistenthash/consistenthash.go](/lib/consistenthash/consistenthash.go) | Go | 64 | 9 | 13 | 86 |
| [lib/consistenthash/consistenthash\_test.go](/lib/consistenthash/consistenthash_test.go) | Go | 15 | 0 | 3 | 18 |
| [lib/geohash/geohash.go](/lib/geohash/geohash.go) | Go | 86 | 7 | 12 | 105 |
| [lib/geohash/geohash\_test.go](/lib/geohash/geohash_test.go) | Go | 35 | 0 | 5 | 40 |
| [lib/geohash/neighbor.go](/lib/geohash/neighbor.go) | Go | 101 | 7 | 15 | 123 |
| [lib/idgenerator/snowflake.go](/lib/idgenerator/snowflake.go) | Go | 56 | 6 | 9 | 71 |
| [lib/idgenerator/snowflake\_test.go](/lib/idgenerator/snowflake_test.go) | Go | 15 | 0 | 3 | 18 |
| [lib/logger/files.go](/lib/logger/files.go) | Go | 35 | 0 | 11 | 46 |
| [lib/logger/logger.go](/lib/logger/logger.go) | Go | 159 | 16 | 23 | 198 |
| [lib/pool/pool.go](/lib/pool/pool.go) | Go | 108 | 8 | 17 | 133 |
| [lib/pool/pool\_test.go](/lib/pool/pool_test.go) | Go | 162 | 1 | 7 | 170 |
| [lib/sync/atomic/bool.go](/lib/sync/atomic/bool.go) | Go | 13 | 3 | 5 | 21 |
| [lib/sync/wait/wait.go](/lib/sync/wait/wait.go) | Go | 31 | 6 | 7 | 44 |
| [lib/timewheel/delay.go](/lib/timewheel/delay.go) | Go | 15 | 3 | 7 | 25 |
| [lib/timewheel/delay\_test.go](/lib/timewheel/delay_test.go) | Go | 17 | 1 | 3 | 21 |
| [lib/timewheel/timewheel.go](/lib/timewheel/timewheel.go) | Go | 166 | 8 | 27 | 201 |
| [lib/timewheel/timewheel\_test.go](/lib/timewheel/timewheel_test.go) | Go | 93 | 13 | 17 | 123 |
| [lib/utils/rand\_string.go](/lib/utils/rand_string.go) | Go | 32 | 2 | 7 | 41 |
| [lib/utils/utils.go](/lib/utils/utils.go) | Go | 83 | 11 | 10 | 104 |
| [lib/wildcard/wildcard.go](/lib/wildcard/wildcard.go) | Go | 67 | 4 | 7 | 78 |
| [lib/wildcard/wildcard\_test.go](/lib/wildcard/wildcard_test.go) | Go | 170 | 6 | 15 | 191 |
| [main.go](/main.go) | Go | 54 | 0 | 7 | 61 |
| [node1.conf](/node1.conf) | Properties | 9 | 0 | 3 | 12 |
| [node2.conf](/node2.conf) | Properties | 9 | 0 | 2 | 11 |
| [node3.conf](/node3.conf) | Properties | 9 | 0 | 2 | 11 |
| [pubsub/hub.go](/pubsub/hub.go) | Go | 15 | 4 | 4 | 23 |
| [pubsub/pubsub.go](/pubsub/pubsub.go) | Go | 121 | 15 | 23 | 159 |
| [redis.conf](/redis.conf) | Properties | 8 | 0 | 3 | 11 |
| [redis/client/client.go](/redis/client/client.go) | Go | 196 | 10 | 23 | 229 |
| [redis/client/client\_test.go](/redis/client/client_test.go) | Go | 129 | 0 | 14 | 143 |
| [redis/connection/conn.go](/redis/connection/conn.go) | Go | 162 | 31 | 39 | 232 |
| [redis/connection/fake.go](/redis/connection/fake.go) | Go | 93 | 6 | 12 | 111 |
| [redis/parser/parser.go](/redis/parser/parser.go) | Go | 210 | 7 | 12 | 229 |
| [redis/parser/parser\_test.go](/redis/parser/parser_test.go) | Go | 77 | 0 | 5 | 82 |
| [redis/protocol/asserts/assert.go](/redis/protocol/asserts/assert.go) | Go | 136 | 9 | 12 | 157 |
| [redis/protocol/consts.go](/redis/protocol/consts.go) | Go | 52 | 16 | 27 | 95 |
| [redis/protocol/errors.go](/redis/protocol/errors.go) | Go | 52 | 13 | 22 | 87 |
| [redis/protocol/reply.go](/redis/protocol/reply.go) | Go | 134 | 32 | 34 | 200 |
| [redis/server/pubsub\_test.go](/redis/server/pubsub_test.go) | Go | 44 | 2 | 5 | 51 |
| [redis/server/server.go](/redis/server/server.go) | Go | 95 | 11 | 12 | 118 |
| [redis/server/server\_test.go](/redis/server/server_test.go) | Go | 41 | 0 | 4 | 45 |
| [tcp/echo.go](/tcp/echo.go) | Go | 65 | 13 | 11 | 89 |
| [tcp/echo\_test.go](/tcp/echo_test.go) | Go | 79 | 1 | 9 | 89 |
| [tcp/server.go](/tcp/server.go) | Go | 79 | 11 | 9 | 99 |

[Summary](results.md) / Details / [Diff Summary](diff.md) / [Diff Details](diff-details.md)