package js

import "github.com/slinkydeveloper/kfn/pkg/util"

func getTestFrameworkPackageJsonConfig(tf string) []packageJsonConfig {
	switch tf {
	case "mocha":
		return []packageJsonConfig{addLatestDevDependency("mocha"), addScript("test", "mocha")}
	case "tape":
		return []packageJsonConfig{addLatestDevDependency("tape"), addScript("test", "tape test/**/*.js")}
	default:
		return []packageJsonConfig{addScript("test", "node index_test.js")}
	}
}
