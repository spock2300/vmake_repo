package main

import "github.com/spock2300/vmake/pkg/api"

func Main(p *api.Package) {
	p.OnPackage(func(p *api.Package) {
		p.SetHomepage("https://github.com/cisco/libsrtp").
			SetDescription("Secure RTP (SRTP) and UST Reference Implementations").
			SetLicense("BSD-3-Clause").
			SetGit(
				"https://gitee.com/mirrors/libsrtp2.git",
				"https://github.com/cisco/libsrtp.git",
			).
			AddVersion("2.5.0", "v2.5.0").
			AddVersion("2.4.2", "v2.4.2")
	})

	p.OnConfig(func(ctx *api.ConfigContext) {
		ctx.Option("shared").
			SetDefault(false).
			SetDescription("Build shared library")
	})

	p.OnBuild(func(ctx *api.BuildContext) {
		ctx.Target("srtp2").
			SetKind(api.TargetVoid).
			AddProvidedLibs("srtp2").
			SetBuildFunc(func(p *api.Package) error {
				args := []string{
					"-DENABLE_TESTS=OFF",
					"-DBUILD_SHARED_LIBS=" + p.BoolStr("shared"),
				}
				p.CMakeConfigure(args...)
				p.CMakeBuild()
				p.CMakeInstall()
				return nil
			})
	})
}
