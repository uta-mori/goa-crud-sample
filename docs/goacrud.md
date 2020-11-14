# GOA フレームワークを使って API を構築する

始める前に、GOA とは何かを理解しておきましょう。

- Goa は、Go で API やマイクロサービスを開発するためのフレームワークです。
- Goa はコード生成に完全に依存しており、反復的なコーディングの必要性を減らします。
- Goa はデザインベースのフレームワークで、「デザインは唯一の真理の源である」と言われており、デザインとドキュメントの両方がそこから派生しています。
- より詳細な情報はこちらをご覧ください。

ホームディレクトリの下に新しいプロジェクトを作成します

`go mod init book`

次に、Goa モジュールがインストールされ、最新であることを確認します

```
go get -u goa.design/goa/v3
go get -u goa.design/goa/v3/...
```

それでは、book の記録を作成し、すべての book を一覧表示するための設計から始めましょう。
ファイル`design/design.go`を作成して開き、次のコードを貼り付けます…

```go
package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("book", func() {
	Title("Book Store")
	Description("Service to perform CRUD operations using goa")
	Server("book", func() {
		Host("localhost", func() {
			URI("http://localhost:8000")
		})
	})
})
```

さらに、book 型を作成して追加し、`design/type.go`を開いて、次のコードを貼り付けます。

```go
package design

import . "goa.design/goa/v3/dsl"

var Book = ResultType("application/vnd.book", "Book", func() {
	Description("Details of a book")

	Attribute("id", UInt32, "ID of the book", func() {
		Example("id", 1)
	})
	Attribute("name", String, "Name of book", func() {
		Example("name", "book1")
		MaxLength(100)
	})
	Attribute("description", String, "Description of the book", func() {
		Example("name", "Books are human's best friend")
		MaxLength(100)
	})
	Attribute("price", UInt32, "Price of the book", func() {
		Example("price", 100)
	})

	Required("id", "name", "description", "price")
})
```

それでは、`design/design.go`ファイルを更新して、作成と一覧取得のメソッドを作成しましょう。

```go
var _ = Service("book", func() {
	Description("The book service gives details of the book.")

	Error("not-found", ErrorResult, "Book Not Found Error")

	//Method to add a new book
	Method("create", func() {
		Description("Adds a new book to the book store.")
		Payload(Book)
		Result(Book)
		HTTP(func() {
			POST("/")
			Response(StatusCreated)
		})
	})

	//Method to get all existing books
	Method("list", func() {
		Description("List all entries")
		Result(ArrayOf(Book))
		HTTP(func() {
			GET("/books")
			Response(StatusOK)
		})
	})
})
```

サービスの設計が完了したので、`goa gen`コマンドを実行してコードを生成しましょう。design パッケージは book モジュールの下に作成されているため、コマンドラインは次のようになります。

`goa gen book/design`

以下のようなファイルが生成されます

```
gen
├── book
│   ├── client.go
│   ├── endpoints.go
│   ├── service.go
│   └── views
│       └── view.go
└── http
    ├── book
    │   ├── client
    │   │   ├── client.go
    │   │   ├── cli.go
    │   │   ├── encode_decode.go
    │   │   ├── paths.go
    │   │   └── types.go
    │   └── server
    │       ├── encode_decode.go
    │       ├── paths.go
    │       ├── server.go
    │       └── types.go
    ├── cli
    │   └── book
    │       └── cli.go
    ├── openapi3.json
    ├── openapi3.yaml
    ├── openapi.json
    └── openapi.yaml
```

次に、`goa example`コマンドを実行して、サービスの基本的な実装と、ゴルーチンを起動して HTTP を開始するビルド可能なサーバーファイルを生成できます。

```
├── book.go
├── cmd
│   ├── book
│   │   ├── http.go
│   │   └── main.go
│   └── book-cli
│       ├── http.go
│       └── main.go
```

`book.go`ファイルは以下のようになります

```go
package bookapi

import (
	book "book/gen/book"
	"context"
	"log"
)

// book service example implementation.
// The example methods log the requests and return zero values.
type booksrvc struct {
	logger *log.Logger
}

// NewBook returns the book service implementation.
func NewBook(logger *log.Logger) book.Service {
	return &booksrvc{logger}
}

// Adds a new book to the book store.
func (s *booksrvc) Create(ctx context.Context, p *book.Book) (res *book.Book, err error) {
	res = &book.Book{}
	s.logger.Print("book.create")
	return
}

// List all entries
func (s *booksrvc) List(ctx context.Context) (res []*book.Book, err error) {
	s.logger.Print("book.list")
	return
}

```

ファイル book.go を編集し、create および list メソッドを実装します。 book.go の次のコードをコピーして置き換えます
