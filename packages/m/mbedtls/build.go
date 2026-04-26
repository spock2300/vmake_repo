package main

import (
	"os"
	"path/filepath"
	"strings"

	"gitee.com/spock2300/vmake/pkg/api"
)

func Main(p *api.Package) {
	p.OnPackage(func(p *api.Package) {
		p.SetHomepage("https://www.trustedfirmware.org/projects/mbed-tls/").
			SetDescription("Lightweight, open source SSL/TLS library").
			SetLicense("Apache-2.0").
			SetGit(
				"https://gitee.com/mirrors/mbedtls.git",
				"https://github.com/Mbed-TLS/mbedtls.git",
			).
			SetSubmodules(true).
			AddVersion("3.6.0", "v3.6.0").
			AddVersion("2.28.8", "v2.28.8")
	})

	p.OnConfig(func(ctx *api.ConfigContext) {
		ctx.Option("shared").
			SetType(api.OptionBool).
			SetDefault(false).
			SetDescription("Build shared library")
	})

	p.OnBuild(func(ctx *api.BuildContext) {
		ctx.Target("mbedtls").
			SetKind(api.TargetVoid).
			AddProvidedLibs("mbedtls", "mbedx509", "mbedcrypto").
			SetBuildFunc(func(p *api.Package) error {
				args := []string{
					"-DUSE_SHARED_MBEDTLS_LIBRARY=" + p.BoolStr("shared"),
					"-DENABLE_PROGRAMS=OFF",
					"-DENABLE_TESTING=OFF",
				}
				configPath := filepath.Join(p.SrcDir(), "include", "mbedtls", "config.h")
				raw, _ := os.ReadFile(configPath)
				if raw != nil {
					s := strings.Replace(string(raw),
						"//#define MBEDTLS_SSL_DTLS_SRTP\n",
						"#define MBEDTLS_SSL_DTLS_SRTP\n", 1)
					os.WriteFile(configPath, []byte(s), 0644)
				}
				p.CMakeConfigure(args...)
				p.CMakeBuild("-j4")
				p.CMakeInstall()
				return nil
			})
	})
}
