# friend-code

好友邀请码：字符集去混淆、带校验位，玩家与码双向唯一，命令行用位置参数子命令，只用 Go 标准库。

```
go run . generate alice
go run . verify ABCDEF
go run . sample charset
go run . sample checksum
go run . sample dup
go run . sample case
go run . sample work
go test ./...
go vet ./...
```

## 口径（README 为准）

- **字符集**：邀请码只用 32 个不易混淆的字符（大写字母去掉 `I`、`O`，数字只用 `2` 到 `9`）。
- **长度与校验位**：码长固定 6 位，最后一位是校验位：前 5 位的字符序号和乘 31 再对 32 取模，得到的下标取字符。
- **大小写归一**：校验前先把输入统一成大写，再判字符与校验位。
- **唯一与幂等**：一个玩家只对应一个码，同一个玩家重复生成返回同一个码；一个码只能属于一个玩家。
- **错误分类**：长度不对返回 `ErrLength`，含非法字符返回 `ErrChar`，校验位不符返回 `ErrChecksum`。
- **代价**：`Owner` 按码查归属不许全扫台账，`Scanned` 不随台账条目数乘查询次数增长；10 万玩家每秒 10 万次查询。

## 输出契约（不改格式）

```
player=.. code=..
ok=..
same=..
lowercase_ok=..
confusable=..
work=..
```

## 目录

```
main.go               位置参数子命令入口
codec/codec.go        字符集、编码与校验
registry/registry.go  玩家与码的双向台账
codec/codec_test.go   编码与校验用例
registry/registry_test.go 台账用例
```
