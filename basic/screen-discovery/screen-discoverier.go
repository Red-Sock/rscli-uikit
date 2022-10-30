package screen_discovery

import rscliuitkit "github.com/Red-Sock/rscli-uikit"

type ScreenDiscovery struct {
	PreviousScreen rscliuitkit.UIElement
}

func (sd *ScreenDiscovery) SetPreviousScreen(element rscliuitkit.UIElement) {
	sd.PreviousScreen = element
}
