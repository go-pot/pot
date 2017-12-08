package logo

import "gopkg.in/pot.v1"

var logo = `
 ______
 | ___ \    _
 | |_/ /__ | |_
 |  __/ _ \| __|
 | | | (_) | |_
 \_|  \___/ \__|
                 %v
`

func PrintLogo(ver interface{}) {
	pot.Printf(logo, ver)
}
