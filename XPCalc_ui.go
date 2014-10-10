// This file was created by ui2walk and may be regenerated.
// DO NOT EDIT OR YOUR MODIFICATIONS WILL BE LOST!

package main

import (
	"github.com/lxn/walk"
)

type dialogUI struct {
	textEdit               *walk.TextEdit
	
	gBoxNumberNotation			*walk.GroupBox
	rbNumberNotation_16			*walk.RadioButton
	rbNumberNotation_10			*walk.RadioButton
	rbNumberNotation_08			*walk.RadioButton
	rbNumberNotation_02			*walk.RadioButton
	
	gBoxNumberUnitBytes			*walk.GroupBox
	rbNumberUnitBytes_08		*walk.RadioButton
	rbNumberUnitBytes_04		*walk.RadioButton
	rbNumberUnitBytes_02		*walk.RadioButton
	rbNumberUnitBytes_01		*walk.RadioButton
	
	cbInverseFunction				*walk.CheckBox
	cbHyperbolicFunction		*walk.CheckBox
	
	gBoxFunctionOps					*walk.GroupBox
	tbBackspace							*walk.ToolButton
	tbClearError						*walk.ToolButton
	tbClear									*walk.ToolButton
	
	gBoxStatisticalCalculation	*walk.GroupBox
	tbStatistical								*walk.ToolButton
	tbAverage										*walk.ToolButton
	tbSum												*walk.ToolButton
	tbs													*walk.ToolButton
	tbDat												*walk.ToolButton
	
	gBoxMainOps							*walk.GroupBox
	tbNumber0								*walk.ToolButton
	tbNumber1								*walk.ToolButton
	tbNumber2								*walk.ToolButton
	tbNumber3								*walk.ToolButton
	tbNumber4								*walk.ToolButton
	tbNumber5								*walk.ToolButton
	tbNumber6								*walk.ToolButton
	tbNumber7								*walk.ToolButton
	tbNumber8								*walk.ToolButton
	tbNumber9								*walk.ToolButton
	tbCharA									*walk.ToolButton
	tbCharB									*walk.ToolButton
	tbCharC									*walk.ToolButton
	tbCharD									*walk.ToolButton
	tbCharE									*walk.ToolButton
	tbCharF									*walk.ToolButton
	tbOpsAdd								*walk.ToolButton
	tbOpsSub								*walk.ToolButton
	tbOpsMul								*walk.ToolButton
	tbOpsDiv								*walk.ToolButton
	tbOpsMod								*walk.ToolButton
	tbOpsPolar							*walk.ToolButton
	tbOpsPoint							*walk.ToolButton
	tbOpsEqual							*walk.ToolButton
	tbOpsAnd								*walk.ToolButton
	tbOpsOr									*walk.ToolButton
	tbOpsXor								*walk.ToolButton
	tbOpsNot								*walk.ToolButton
	tbOpsLsh								*walk.ToolButton
	tbOpsInt								*walk.ToolButton
}

func (w *Dialog) init(owner walk.Form) (err error) {
	if w.Dialog, err = walk.NewDialog(owner); err != nil {
		return err
	}

	succeeded := false
	defer func() {
		if !succeeded {
			w.Dispose()
		}
	}()

	var font *walk.Font
	if font == nil {
		font = nil
	}

	w.SetName("XPCalc")
	if err := w.SetClientSize(walk.Size{600, 300}); err != nil {
		return err
	}
	if err := w.SetTitle(`XPCalc`); err != nil {
		return err
	}
	
	// textEdit
	if w.ui.textEdit, err = walk.NewTextEdit(w); err != nil {
		return err
	}
	w.ui.textEdit.SetName("textEdit")
	if err := w.ui.textEdit.SetBounds(walk.Rectangle{0, 0, 568, 32}); err != nil {
		return err
	}

	// gBoxNumberNotation
	if w.ui.gBoxNumberNotation, err = walk.NewGroupBox(w); err != nil {
		return err
	}
	w.ui.gBoxNumberNotation.SetName("gBoxNumberNotation")
	if err := w.ui.gBoxNumberNotation.SetBounds(walk.Rectangle{0, 64, 0, 32}); err != nil {
		return err
	}
	if err := w.ui.gBoxNumberNotation.SetTitle(`gBoxNumberNotation`); err != nil {
		return err
	}
	
	// rbNumberNotation_16
	if w.ui.rbNumberNotation_16, err = walk.NewToolButton(w.ui.gBoxNumberNotation); err != nil {
		return err
	}
	w.ui.rbNumberNotation_16.SetName("rbNumberNotation_16")
	if err := w.ui.rbNumberNotation_16.SetBounds(walk.Rectangle{10, 12, 59, 18}); err != nil {
		return err
	}
	if err := w.ui.rbNumberNotation_16.SetText(`Connected`); err != nil {
		return err
	}

	succeeded = true

	return nil
}
