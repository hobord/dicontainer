package dal2

type Option func(*dal2)

func WithText(text string) Option {
	return func(d *dal2) {
		d.text = text
	}
}
