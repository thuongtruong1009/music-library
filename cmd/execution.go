package cmd

import (
	"fmt"
	"music-management-system/pkg/constants"
)

func (d *Delivery) Execution() {
	loop := func() {
		d.helper.OutputNomal(constants.LABEL, constants.BANNER_TEXT)

		option := d.Run()
		d.HandleOption(option)
	}
	loop()

	for {
		fmt.Print(d.helper.OutputColor(constants.QUERY) + constants.AGAIN_TEXT)
		var input string
		fmt.Scanln(&input)
		if input == "y" {
			d.helper.ClearConsole()
			loop()
		} else {
			d.helper.ClearConsole()
			d.helper.OutputNomal(constants.LABEL, constants.THANKYOU_TEXT)
			break
		}
	}
}