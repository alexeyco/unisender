package unisender

// Option is optional request value
type Option struct {
	name  string
	value string
}

// OptionBeforeSubscribeUrl returns before_subscribe_url_option
func OptionBeforeSubscribeUrl(u string) Option {
	return Option{
		name:  "before_subscribe_url",
		value: u,
	}
}

// OptionAfterSubscribeUrl returns after_subscribe_url option
func OptionAfterSubscribeUrl(u string) Option {
	return Option{
		name:  "after_subscribe_url",
		value: u,
	}
}
