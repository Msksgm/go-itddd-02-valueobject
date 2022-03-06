# GO-ITDDD-02-VALUEOBJECT

zenn の記事「[Go で値オブジェクトを実装（「入門ドメイン駆動設計」Chapter2）](https://zenn.dev/msksgm/articles/20220304-go-itddd-02-valueobject)」のサンプルコードです。

# 実行環境

- Go
  - 1.17

# 実行方法

fullname

```bash
> go run cmd/fullname/fullname.go
nameA: FullName [firstName=firstName, lastName=lastName] is equal to nameB: FullName [firstName=firstName, lastName=lastName]
```

money

```bash
> go run cmd/money/money.go
Money[amoun=4000, currency=JPY]
```

# テスト

```bash
?   	github.com/Msksgm/itddd-go-02-valueobject/cmd/fullname	[no test files]
?   	github.com/Msksgm/itddd-go-02-valueobject/cmd/money	[no test files]
=== RUN   TestNewFullName
=== RUN   TestNewFullName/success
=== RUN   TestNewFullName/fail_firstName_is_empty
=== RUN   TestNewFullName/fail_firstName_is_invalid
=== RUN   TestNewFullName/fail_lastName_is_empty
=== RUN   TestNewFullName/fail_lastName_is_invalid
--- PASS: TestNewFullName (0.00s)
    --- PASS: TestNewFullName/success (0.00s)
    --- PASS: TestNewFullName/fail_firstName_is_empty (0.00s)
    --- PASS: TestNewFullName/fail_firstName_is_invalid (0.00s)
    --- PASS: TestNewFullName/fail_lastName_is_empty (0.00s)
    --- PASS: TestNewFullName/fail_lastName_is_invalid (0.00s)
=== RUN   TestWithChangeFirstName
--- PASS: TestWithChangeFirstName (0.00s)
=== RUN   TestWithChangeLastName
--- PASS: TestWithChangeLastName (0.00s)
=== RUN   TestEquals
=== RUN   TestEquals/equal
=== RUN   TestEquals/not_equal
--- PASS: TestEquals (0.00s)
    --- PASS: TestEquals/equal (0.00s)
    --- PASS: TestEquals/not_equal (0.00s)
PASS
ok  	github.com/Msksgm/itddd-go-02-valueobject/domain/model/fullname	0.110s
=== RUN   TestNewMoney
=== RUN   TestNewMoney/success
--- PASS: TestNewMoney (0.00s)
    --- PASS: TestNewMoney/success (0.00s)
=== RUN   TestAdd
=== RUN   TestAdd/success
=== RUN   TestAdd/fail_currency_is_different
--- PASS: TestAdd (0.00s)
    --- PASS: TestAdd/success (0.00s)
    --- PASS: TestAdd/fail_currency_is_different (0.00s)
PASS
ok  	github.com/Msksgm/itddd-go-02-valueobject/domain/model/money	0.182s
=== RUN   TestWrap
--- PASS: TestWrap (0.00s)
PASS
ok  	github.com/Msksgm/itddd-go-02-valueobject/iterrors	0.071s

```
