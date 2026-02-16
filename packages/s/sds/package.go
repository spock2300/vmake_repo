package main

import "gitee.com/spock2300/vmake/pkg/api"

func Package(p *api.Package) {
	p.SetHomepage("https://github.com/redis/sds").
		SetDescription("Simple Dynamic Strings for C").
		SetLicense("BSD-2-Clause").
		SetGit("https://github.com/redis/sds.git")

	p.AddVersion("2.0.0", "2.0.0")
	p.AddVersion("1.0.5", "1.0.5")

	p.Build(func(ctx *api.PackageContext) {
		t := ctx.Target("sds").
			SetKind(api.TargetStatic).
			AddFiles(ctx.SourceDir() + "/sds.c").
			AddIncludes(ctx.SourceDir())
		ctx.Build(t)

		ctx.MkdirAll(ctx.InstallDir() + "/lib")
		ctx.MkdirAll(ctx.InstallDir() + "/include")
		ctx.CopyFile(ctx.BuildDir()+"/libsds.a", ctx.InstallDir()+"/lib/")
		ctx.CopyFile(ctx.SourceDir()+"/sds.h", ctx.InstallDir()+"/include/")
	})
}
