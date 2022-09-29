package label

import rscliuitkit "github.com/Red-Sock/rscli-uikit"

type Attribute func(l *Label)

func AttributeNextScreen(next func() rscliuitkit.UIElement) Attribute {
	return func(l *Label) {
		l.next = next
	}
}
