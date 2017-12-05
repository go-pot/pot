package logo

import "github.com/wzshiming/pot"

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
