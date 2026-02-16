package main

import "gitee.com/spock2300/vmake/pkg/api"

func Package(p *api.Package) {
	p.SetHomepage("https://codeplea.com/tinyexpr").
		SetDescription("Tiny expression evaluator for C").
		SetLicense("zlib").
		SetGit("https://github.com/codeplea/tinyexpr.git")

	p.AddVersion("1.0.0", "v1.0")

	p.Build(func(ctx *api.PackageContext) {
		t := ctx.Target("tinyexpr").
			SetKind(api.TargetStatic).
			AddFiles(ctx.SourceDir() + "/tinyexpr.c").
			AddIncludes(ctx.SourceDir())
		ctx.Build(t)

		ctx.MkdirAll(ctx.InstallDir() + "/lib")
		ctx.MkdirAll(ctx.InstallDir() + "/include")
		ctx.CopyFile(ctx.BuildDir()+"/libtinyexpr.a", ctx.InstallDir()+"/lib/")
		ctx.CopyFile(ctx.SourceDir()+"/tinyexpr.h", ctx.InstallDir()+"/include/")
	})
}
