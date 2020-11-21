module github.com/finalone/easy-blog

go 1.14

require (
	github.com/finalone/easy-blog/app/controller v0.0.0
	github.com/finalone/easy-blog/app/dao v0.0.0
	github.com/finalone/easy-blog/app/model v0.0.0
	github.com/finalone/easy-blog/app/service v0.0.0
	github.com/gin-gonic/gin v1.6.3
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/go-sql-driver/mysql v1.4.0
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/jmoiron/sqlx v1.2.0
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/michelloworld/ez-gin-template v0.0.0-20171028145326-7d77b0197797
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/ugorji/go v1.1.13 // indirect
	golang.org/x/crypto v0.0.0-20201016220609-9e8e0b390897 // indirect
	golang.org/x/sys v0.0.0-20201018230417-eeed37f84f13 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/yaml.v2 v2.3.0
)

replace (
	github.com/finalone/easy-blog/app/controller => ./app/controller
	github.com/finalone/easy-blog/app/dao => ./app/dao
	github.com/finalone/easy-blog/app/model => ./app/model
	github.com/finalone/easy-blog/app/service => ./app/service
)
