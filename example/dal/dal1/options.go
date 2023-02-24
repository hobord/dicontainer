package dal1

type Option func(*dal1)

func WithText(text string) Option {
	return func(d *dal1) {
		d.text = text
	}
}
