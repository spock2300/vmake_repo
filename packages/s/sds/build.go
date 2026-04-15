package main

import "gitee.com/spock2300/vmake/pkg/api"

func Main(p *api.Package) {
	p.OnPackage(func(p *api.Package) {
		p.SetHomepage("https://github.com/redis/sds").
			SetDescription("Simple Dynamic Strings for C").
			SetLicense("BSD-2-Clause").
			SetGit("https://github.com/redis/sds.git").
			AddVersion("2.0.0", "2.0.0").
			AddVersion("1.0.5", "1.0.5")
	})

	p.OnBuild(func(ctx *api.BuildContext) {
		ctx.Target("sds").
			SetKind(api.TargetStatic).
			AddFiles("sds.c").
			AddPublicIncludes(".", "@*.h").
			RemoveCFlags("-Werror", "-Wmissing-prototypes", "-Wmissing-declarations", "-Wold-style-definition", "-Wshadow")
	})
}
