package main

import (
	"path/filepath"

	"github.com/spock2300/vmake/pkg/api"
)

func Main(p *api.Package) {
	p.OnPackage(func(p *api.Package) {
		p.SetHomepage("https://github.com/DaveGamble/cJSON").
			SetDescription("Ultralightweight JSON parser in ANSI C").
			SetLicense("MIT").
			SetGit(
				"https://gitee.com/mirrors/cJSON.git",
				"https://github.com/DaveGamble/cJSON.git",
			).
			AddVersion("1.7.18", "v1.7.18").
			AddVersion("1.7.17", "v1.7.17")
	})

	p.OnConfig(func(ctx *api.ConfigContext) {
		ctx.Option("shared").
			SetType(api.OptionBool).
			SetDefault(false).
			SetDescription("Build shared library")
	})

	p.OnBuild(func(ctx *api.BuildContext) {
		ctx.Target("cjson").
			SetKind(api.TargetVoid).
			AddProvidedLibs("cjson").
			SetBuildFunc(func(p *api.Package) error {
				args := []string{
					"-DENABLE_CJSON_TEST=OFF",
					"-DBUILD_SHARED_LIBS=" + p.BoolStr("shared"),
					"-DCMAKE_INSTALL_INCLUDEDIR=include",
				}
				p.CMakeConfigure(args...)
				p.CMakeBuild()
				p.CMakeInstall()
				src := filepath.Join(p.CMakeInstallDir(), "include", "cjson", "cJSON.h")
				dst := filepath.Join(p.CMakeInstallDir(), "include", "cJSON.h")
				api.CopyFile(src, dst)
				return nil
			})
	})
}
