package composit_label

import "github.com/Red-Sock/rscli-uikit/utils/common"

type Attribute func(cl *ComplexLabel)

func Position(positioner common.Positioner) Attribute {
	return func(cl *ComplexLabel) {
		cl.pos = positioner
	}
}
