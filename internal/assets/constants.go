package assets

const ENV_LOGLEVEL_KEY string = "LOG_LEVEL"
const ENV_DEBUG_KEY string = "DEBUG"
const ENV_FORCE_HTTPS_KEY string = "FORCE_HTTPS"
const ENV_ADMIN_USERNAME_KEY string = "ADMIN_USERNAME"
const ENV_ADMIN_PASSWORD_KEY string = "ADMIN_PASSWORD"

const PGPOutputDir string = "./enc/pgp_out"
const UsesOutputDir string = "./enc/uses"
const ContactOutputDir string = "./enc/contact"

const SQLitePath string = "./conf/main.sqlite"

const DBKEY_BOX_PRICE string = "BoxSetPrice"
const DBKEY_STICKER_PRICE string = "StickerPrice"
const DBKEY_PUBLIC_KEY string = "PublicKey"

const LOG_OUTPUT_DIR string = "./logs"

const BoxSetPrice float32 = 43.75
const StickerCost float32 = 2.0

const USERNAME_MAX_LEN int = 16
const USERNAME_MIN_LEN int = 3

const PASSWORD_MAX_LEN int = 64 // limited by bcrypt
const PASSWORD_MIN_LEN int = 16

var AllowOriginDomains []string = []string{
	"https://astoryofand.com",
	"https://*.astoryofand.com",
}
