- [参考にしたページ](https://www.webprofessional.jp/mock-rest-apis-using-json-server/)

- JSON Serverのインストール

```
$ npm install -g json-server

```

- スタブを作成

```
$ cat db.json

{
    "movies": [
        {
            "id": 1,
            "name": "The Godfather",
            "director": "Francis Ford Coppola",
            "rating": 9.1
        },
        {
            "id": 2,
            "name": "Casablanca",
            "director": "Michael Curtiz",
            "rating": 8.8
        },
        {
            "id": 4,
            "name": "Inception",
            "director": "Christopher Nolan",
            "rating": 9
        },
        {
            "id": 6,
            "name": "The Godfather",
            "director": "Francis Ford Coppola",
            "rating": 9.1
        },
        {
            "id": 7,
            "name": "Casablanca",
            "director": "Michael Curtiz",
            "rating": 8.8
        },
        {
            "id": 9,
            "name": "The Godfather",
            "director": "Francis Ford Coppola",
            "rating": 9.1
        },
        {
            "id": 10,
            "name": "Casablanca",
            "director": "Michael Curtiz",
            "rating": 8.8
        }
    ]
}

```

- サーバーを起動

```
$ json-server --watch db.json

```
$ json-server --watch db.json

  \{^_^}/ hi!

  Loading db.json
  Done

  Resources
  http://localhost:3000/movies

  Home
  http://localhost:3000

```

- curlでHTTPリクエスト
    - postmanでも良い

```
$  curl -X GET "http://localhost:3000/movies"
[
  {
    "id": 1,
    "name": "The Godfather",
    "director": "Francis Ford Coppola",
    "rating": 9.1
  },

...

```

```
$  curl -X GET "http://localhost:3000/movies/1"
{
  "id": 1,
  "name": "The Godfather",
  "director": "Francis Ford Coppola",
  "rating": 9.1
}

```

- POSTしてみる

```
 $  curl -X POST -H "Content-Type: application/json" -d '{
  "id": 3,
  "name": "Inception",
  "director": "Christopher Nolan",
  "rating": 9.0
}' "http://localhost:3000/movies"

$  curl -X GET "http://localhost:3000/movies/3"
{
  "id": 3,
  "name": "Inception",
  "director": "Christopher Nolan",
  "rating": 9
}

```

- APIサーバーを起動したまま、HTMLをブラウザで読み込む
    - h1要素をクリックするとJSONの値を取得してブラウザに表示