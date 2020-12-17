# simple-jwt-auth-in-go

GoでJWTを使ったシンプルな認証を実装してみる。  
※DBは利用していないため、起動するたびに`signin`する。  
※簡単のために、モジュールの分割はある程度にとどめています。  


## /signin
サービスへの登録ロジック。

### リクエスト
```
{
    "name": "Kade",
    "password": "qwerty"
}
```

### レスポンス
```
{
    "name": "Kade",
    "Token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiS2FkZSIsImV4cCI6MTYwODI4NDcxNH0.my6PvbNhkg_i1w_cX0UmrK3AJZ_1e3WtMBiL-urtA0s"
}
```

## /login
サービスにログインしてトークンを取得する。 

### リクエスト
```
{
    "name": "Kade",
    "password": "qwerty"
}
```

### レスポンス
```
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiS2FkZSIsImV4cCI6MTYwODI4NTI3N30._4H0KlwgRQODMol7-Kgy4bcAu7FSTVeOSzRQvriDJCI"
}
```

## GET /hello
ログインしている（`Authorization`ヘッダに有効なJWTトークンが存在する）場合は、挨拶が表示される。  
ログインしていない場合は、エラーが返ってくる。  

### ログインしている時
```
{"message":"Hello, Kade"}
```

### ログインしていない時
```
401
```

### トークンの有効期限が切れている時
```
token is expired by 1s
```
