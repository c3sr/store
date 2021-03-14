module github.com/c3sr/store

go 1.15

replace google.golang.org/grpc => google.golang.org/grpc v1.29.1

require (
	cloud.google.com/go/storage v1.10.0 // indirect
	github.com/GeertJohan/go-sourcepath v0.0.0-20150925135350-83e8b8723a9b
	github.com/Microsoft/go-winio v0.4.15 // indirect
	github.com/Microsoft/hcsshim v0.8.10 // indirect
	github.com/aws/aws-sdk-go v1.37.30
	github.com/bmizerany/pat v0.0.0-20170815010413-6226ea591a40 // indirect
	github.com/boltdb/bolt v1.3.1
	github.com/c3sr/archive v1.0.0
	github.com/c3sr/aws v1.0.0
	github.com/c3sr/config v1.0.1
	github.com/c3sr/logger v1.0.1
	github.com/c3sr/uuid v1.0.1
	github.com/c3sr/vipertags v1.0.0
	github.com/containerd/continuity v0.0.0-20200928162600-f2cc35102c2a // indirect
	github.com/docker/docker v0.0.0-20200916142827-bd33bbf0497b // indirect
	github.com/dustin/go-humanize v1.0.0
	github.com/fatih/structs v1.1.0
	github.com/golang/mock v1.5.0
	github.com/google/martian/v3 v3.1.0 // indirect
	github.com/google/pprof v0.0.0-20201023163331-3e6fc7fc9c4c // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.1-0.20190118093823-f849b5445de4 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.9.5 // indirect
	github.com/k0kubun/pp/v3 v3.0.7
	github.com/labstack/echo/v4 v4.2.1
	github.com/mattn/go-runewidth v0.0.10 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/opencontainers/runc v0.1.1 // indirect
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/afero v1.5.1
	github.com/spf13/cast v1.3.1
	github.com/spf13/cobra v1.1.3
	github.com/stretchr/testify v1.7.0
	github.com/tus/tusd v0.0.0-20190905174543-852b6fa01dbd
	github.com/unknwon/com v1.0.1
	go.etcd.io/bbolt v1.3.3 // indirect
	golang.org/x/sys v0.0.0-20210313202042-bd2e13477e9c // indirect
	gopkg.in/Acconut/lockfile.v1 v1.1.0
	gopkg.in/cheggaaa/pb.v1 v1.0.28
)
