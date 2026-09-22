package main

import "github.com/spock2300/vmake/pkg/api"

func Main(p *api.Package) {
	p.OnPackage(func(p *api.Package) {
		p.SetHomepage("https://github.com/sctplab/usrsctp").
			SetDescription("Userland SCTP stack supporting FreeBSD, Linux, Mac OS X, and Windows").
			SetLicense("BSD-3-Clause").
			SetGit(
				"https://github.com/sctplab/usrsctp.git",
			).
			SetSubmodules(false).
			AddVersion("0.9.5", "0.9.5.0")
	})

	p.OnConfig(func(ctx *api.ConfigContext) {
		ctx.Option("shared").
			SetType(api.OptionBool).
			SetDefault(false).
			SetDescription("Build shared library")
	})

	p.OnBuild(func(ctx *api.BuildContext) {
		ctx.Target("usrsctp").
			SetKind(api.TargetVoid).
			AddProvidedLibs("usrsctp").
			SetBuildFunc(func(p *api.Package) error {
				args := []string{
					"-DENABLE_PROGRAMS=OFF",
					"-DENABLE_TESTS=OFF",
					"-DBUILD_SHARED_LIBS=" + p.BoolStr("shared"),
					"-DCMAKE_C_FLAGS=" + p.MergedCFlags("-Wno-maybe-uninitialized"),
				}
				p.CMakeConfigure(args...)
				p.CMakeBuild()
				p.CMakeInstall()
				return nil
			})
	})
}
