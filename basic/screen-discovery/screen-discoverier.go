package screen_discovery

import rscliuitkit "github.com/Red-Sock/rscli-uikit"

type ScreenDiscovery struct {
	PreviousScreen rscliuitkit.UIElement
}

func (sd *ScreenDiscovery) SetPreviousScreen(element rscliuitkit.UIElement) {
	if sd.PreviousScreen != nil {
		// Preventing rewriting previous screen, so you can set previous screen
		// at ui-element logic, and it won't be rewritten by main handler
		return
	}
	sd.PreviousScreen = element
}
