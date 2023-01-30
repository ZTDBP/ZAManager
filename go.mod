module github.com/ZTDBP/ZAManager

go 1.17

require (
	github.com/DeanThompson/ginpprof v0.0.0-20201112072838-007b1e56b2e1
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/contrib v0.0.0-20201101042839-6a891bf89f19
	github.com/gin-gonic/gin v1.7.7
	github.com/go-redis/redis/v8 v8.11.5
	github.com/go-resty/resty/v2 v2.7.0
	github.com/google/go-github v17.0.0+incompatible
	github.com/google/uuid v1.3.0
	github.com/rs/xid v1.2.1
	github.com/rubenv/sql-migrate v1.1.1
	github.com/swaggo/files v0.0.0-20190704085106-630677cd5c14
	github.com/swaggo/gin-swagger v1.3.0
	github.com/swaggo/swag v1.7.0
	github.com/urfave/cli v1.22.7
	github.com/ztdbp/cfssl v0.0.5
	github.com/ztdbp/zaca-sdk v0.0.3
	go.uber.org/automaxprocs v1.5.1
	go.uber.org/zap v1.24.0
	golang.org/x/oauth2 v0.0.0-20220622183110-fd043fe589d2
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gopkg.in/yaml.v3 v3.0.1
	gorm.io/driver/mysql v1.3.3
	gorm.io/gorm v1.23.4
)

require (
	cloud.google.com/go/compute v1.7.0 // indirect
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/PuerkitoBio/purell v1.1.1 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/boj/redistore v0.0.0-20180917114910-cd5dcc76aeff // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/cloudflare/backoff v0.0.0-20161212185259-647f3cdfc87a // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/fsnotify/fsnotify v1.5.4 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-gorp/gorp/v3 v3.0.2 // indirect
	github.com/go-openapi/jsonpointer v0.19.3 // indirect
	github.com/go-openapi/jsonreference v0.19.4 // indirect
	github.com/go-openapi/spec v0.19.14 // indirect
	github.com/go-openapi/swag v0.19.11 // indirect
	github.com/go-playground/locales v0.13.0 // indirect
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/gomodule/redigo v2.0.0+incompatible // indirect
	github.com/google/certificate-transparency-go v1.1.4 // indirect
	github.com/google/go-querystring v1.0.0 // indirect
	github.com/gorilla/context v1.1.1 // indirect
	github.com/gorilla/securecookie v1.1.1 // indirect
	github.com/gorilla/sessions v1.2.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.4 // indirect
	github.com/jmoiron/sqlx v1.3.4 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/lib/pq v1.10.3 // indirect
	github.com/mailru/easyjson v0.7.0 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mattn/go-sqlite3 v2.0.3+incompatible // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/spiffe/go-spiffe/v2 v2.0.0-beta.4 // indirect
	github.com/ugorji/go/codec v1.1.13 // indirect
	github.com/weppos/publicsuffix-go v0.13.0 // indirect
	github.com/zeebo/errs v1.2.2 // indirect
	github.com/zmap/zcrypto v0.0.0-20200911161511-43ff0ea04f21 // indirect
	github.com/zmap/zlint/v2 v2.2.1 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	golang.org/x/crypto v0.0.0-20220411220226-7b82a4e95df4 // indirect
	golang.org/x/net v0.0.0-20220722155237-a158d28d115b // indirect
	golang.org/x/sync v0.0.0-20220722155255-886fb9371eb4 // indirect
	golang.org/x/sys v0.0.0-20220722155257-8c9f86f7a55f // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/tools v0.1.12 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace (
	github.com/Azure/azure-sdk-for-go/sdk/azcore v0.19.0 => github.com/Azure/azure-sdk-for-go/sdk/azcore v1.3.0
	github.com/Azure/azure-sdk-for-go/sdk/azcore v0.7.0 => github.com/Azure/azure-sdk-for-go/sdk/azcore v1.1.2
	github.com/Azure/azure-sdk-for-go/sdk/azidentity v0.11.0 => github.com/Azure/azure-sdk-for-go/sdk/azidentity v1.2.1
	github.com/ajstarks/svgo v0.0.0-20180226025133-644b8db467af => github.com/ajstarks/svgo v0.0.0-20211024235047-1546f124cd8b
	github.com/chzyer/logex v1.1.10 => github.com/chzyer/logex v1.2.1
	github.com/prometheus/prometheus v2.5.0+incompatible => github.com/prometheus/prometheus/v2 v2.29.2
	github.com/zmap/rc2 v0.0.0-20131011165748-24b9757f5521 => github.com/zmap/rc2 v0.0.0-20190804163417-abaa70531248
	github.com/zmap/zcrypto v0.0.0-20220803033029-557f3e4940be => github.com/zmap/zcrypto v0.0.0-20200911161511-43ff0ea04f21
	google.golang.org/genproto v0.0.0-20221111202108-142d8a6fa32e => google.golang.org/genproto v0.0.0-20220706185917-7780775163c4
)
