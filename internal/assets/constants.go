package assets

const ENV_LOGLEVEL_KEY string = "LOG_LEVEL"
const ENV_DEBUG_KEY string = "DEBUG"
const ENV_FORCE_HTTPS_KEY string = "FORCE_HTTPS"
const ENV_ADMIN_USERNAME_KEY string = "ADMIN_USERNAME"
const ENV_ADMIN_PASSWORD_KEY string = "ADMIN_PASSWORD"

const PGPOutputDir string = "./enc/pgp_out"
const UsesOutputDir string = "./enc/uses"

const SQLitePath string = "./conf/main.sqlite"

const DBKEY_BOX_PRICE string = "BoxSetPrice"
const DBKEY_STICKER_PRICE string = "StickerPrice"
const DBKEY_PUBLIC_KEY string = "PublicKey"

const BoxSetPrice float32 = 35.0
const StickerCost float32 = 2.0

var AllowOriginDomains []string = []string{
	"https://astoryofand.com",
	"https://*.astoryofand.com",
}
