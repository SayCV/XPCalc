package main

import (
	"fmt"
	//"log"
	//"time"
)

import (
	"github.com/lxn/walk"
)

type Dialog struct {
	*walk.Dialog
	ui dialogUI
}

func rbNumberNotation_10_onCliced(dlg *Dialog) {
	dlg.ui.rbNumberUnitBytes_08.SetText(`ang`)
	dlg.ui.rbNumberUnitBytes_04.SetText(`rad`)
	dlg.ui.rbNumberUnitBytes_02.SetText(`gar`)
	dlg.ui.rbNumberUnitBytes_01.SetText(``)
	dlg.ui.rbNumberUnitBytes_01.SetVisible(false)
}

func rbNumberNotation_xx_onCliced(dlg *Dialog) {
	dlg.ui.rbNumberUnitBytes_08.SetText(`8Bytes`)
	dlg.ui.rbNumberUnitBytes_04.SetText(`4Bytes`)
	dlg.ui.rbNumberUnitBytes_02.SetText(`2Bytes`)
	dlg.ui.rbNumberUnitBytes_01.SetText(`1Bytes`)
	dlg.ui.rbNumberUnitBytes_01.SetVisible(true)
}

func runDialog(owner walk.Form) (int, error) {
	dlg := new(Dialog)
	if err := dlg.init(owner); err != nil {
		return 0, err
	}

	// TODO: Do further required setup, e.g. for event handling, here.
	dlg.ui.rbNumberNotation_16.Clicked().Attach(func() {
		fmt.Println("Clicked rbNumberNotation_16")
		rbNumberNotation_xx_onCliced(dlg)
	})
	dlg.ui.rbNumberNotation_10.Clicked().Attach(func() {
		fmt.Println("Clicked rbNumberNotation_10")
		rbNumberNotation_10_onCliced(dlg)
	})
	dlg.ui.rbNumberNotation_08.Clicked().Attach(func() {
		fmt.Println("Clicked rbNumberNotation_08")
		rbNumberNotation_xx_onCliced(dlg)
	})
	dlg.ui.rbNumberNotation_02.Clicked().Attach(func() {
		fmt.Println("Clicked rbNumberNotation_02")
		rbNumberNotation_xx_onCliced(dlg)
	})
	
	dlg.ui.tbBackspace.Clicked().Attach(func() {
		fmt.Println("Clicked tbBackspace")
		//dlg.ui.rbNumberUnitBytes_01.SetVisible(false)
	})
	
	return dlg.Run(), nil
}
