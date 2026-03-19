package main

import "gitee.com/spock2300/vmake/pkg/api"

func Main(b *api.Builder) {
	b.OnPackage(func(ctx *api.PackageContext) {
		ctx.SetHomepage("https://github.com/redis/sds").
			SetDescription("Simple Dynamic Strings for C").
			SetLicense("BSD-2-Clause").
			SetGit("https://github.com/redis/sds.git").
			AddVersion("2.0.0", "2.0.0").
			AddVersion("1.0.5", "1.0.5")
	})

	b.OnBuild(func(ctx *api.BuildContext) {
		ctx.Target("sds").
			SetKind(api.TargetStatic).
			AddFiles("sds.c").
			AddPublicIncludes(".")
	})
}
