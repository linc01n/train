package train

type config struct {
	AssetsPath   string
	AssetsUrl    string
	BundleAssets bool
}

var Config config = config{
	AssetsPath: "assets",
	AssetsUrl:  "/assets",
}

func init() {
	Config.BundleAssets = HasPublicAssets()
}
