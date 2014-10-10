package main

import (
	"github.com/lxn/walk"
)

type Dialog struct {
	*walk.Dialog
	ui dialogUI
}

func runDialog(owner walk.Form) (int, error) {
	dlg := new(Dialog)
	if err := dlg.init(owner); err != nil {
		return 0, err
	}

	// TODO: Do further required setup, e.g. for event handling, here.

	return dlg.Run(), nil
}
