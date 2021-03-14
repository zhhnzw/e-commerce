module goods

go 1.12

replace (
	cloud.google.com/go => github.com/googleapis/google-cloud-go v0.37.4
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20181203042331-505ab145d0a9
	golang.org/x/exp => github.com/golang/exp v0.0.0-20190402192236-7fd597ecf556
	golang.org/x/image => github.com/golang/image v0.0.0-20190321063152-3fc05d484e9f
	golang.org/x/lint => github.com/golang/lint v0.0.0-20190313153728-d0100b6bd8b3
	golang.org/x/mobile => github.com/golang/mobile v0.0.0-20190327163128-167ebed0ec6d
	golang.org/x/net => github.com/golang/net v0.0.0-20190311031020-56fb01167e7d
	golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20190523182746-aaccbc9213b0
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190227155943-e225da77a7e6
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190225065934-cc5685c2db12
	golang.org/x/text => github.com/golang/text v0.3.0
	golang.org/x/time => github.com/golang/time v0.0.0-20190308202827-9d24e82272b4
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190402200628-202502a5a924
	google.golang.org/api => github.com/googleapis/google-api-go-client v0.5.0
	google.golang.org/appengine => github.com/golang/appengine v1.6.0
	google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20180831171423-11092d34479b
	google.golang.org/grpc => github.com/grpc/grpc-go v1.21.0
)

require (
	github.com/fsnotify/fsnotify v1.4.9
	github.com/go-redis/redis v6.15.2+incompatible
	github.com/go-sql-driver/mysql v1.4.1
	github.com/golang/protobuf v1.4.3
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.0
	github.com/jinzhu/gorm v1.9.10
	github.com/kr/pretty v0.2.1 // indirect
	github.com/natefinch/lumberjack v2.0.0+incompatible
	github.com/onsi/ginkgo v1.15.1 // indirect
	github.com/onsi/gomega v1.11.0 // indirect
	github.com/pkg/errors v0.8.1
	github.com/satori/go.uuid v1.2.0
	github.com/smartystreets/goconvey v1.6.4
	github.com/spf13/viper v1.7.1
	go.uber.org/zap v1.16.0
	google.golang.org/grpc v1.21.1
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
)
