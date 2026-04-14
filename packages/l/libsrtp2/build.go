package main

import (
	"runtime"
	"strconv"

	"gitee.com/spock2300/vmake/pkg/api"
)

func Main(p *api.Package) {
	p.OnPackage(func(p *api.Package) {
		p.SetHomepage("https://github.com/cisco/libsrtp").
			SetDescription("Secure RTP (SRTP) and UST Reference Implementations").
			SetLicense("BSD-3-Clause").
			SetGit(
				"https://gitee.com/mirrors/libsrtp2.git",
				"https://github.com/cisco/libsrtp.git",
			).
			SetLibs("srtp2").
			AddVersion("2.5.0", "v2.5.0").
			AddVersion("2.4.2", "v2.4.2")
	})

	p.OnConfig(func(ctx *api.ConfigContext) {
		ctx.Option("shared").
			SetType(api.OptionBool).
			SetDefault(false).
			SetDescription("Build shared library")
	})

	p.OnBuild(func(ctx *api.BuildContext) {
		ctx.Target("srtp2").
			SetKind(api.TargetVoid).
			SetBuildFunc(func(p *api.Package) error {
				args := []string{
					"-DENABLE_TESTS=OFF",
					"-DBUILD_SHARED_LIBS=" + p.BoolStr("shared"),
				}
				p.CMakeConfigure(args...)
				p.CMakeBuild("-j" + strconv.Itoa(runtime.NumCPU()))
				p.CMakeInstall()
				return nil
			})
	})
}
