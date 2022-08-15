module trpc.app.Greeter

go 1.12

replace git.code.oa.com/trpcprotocol/test/helloworld => ./stub/git.code.oa.com/trpcprotocol/test/helloworld

require (
	git.code.oa.com/tpstelemetry/tps-sdk-go/instrumentation/trpctelemetry v0.4.19
	git.code.oa.com/trpc-go/trpc-config-rainbow v0.1.24
	git.code.oa.com/trpc-go/trpc-filter/debuglog v0.1.5
	git.code.oa.com/trpc-go/trpc-filter/recovery v0.1.3
	git.code.oa.com/trpc-go/trpc-go v0.9.4
	git.code.oa.com/trpc-go/trpc-log-atta v0.1.14
	git.code.oa.com/trpc-go/trpc-metrics-m007 v0.5.1
	git.code.oa.com/trpc-go/trpc-metrics-runtime v0.3.3
	git.code.oa.com/trpc-go/trpc-naming-polaris v0.3.4
	git.code.oa.com/trpcprotocol/test/helloworld v0.0.0-00010101000000-000000000000
	github.com/golang/mock v1.6.0
)
