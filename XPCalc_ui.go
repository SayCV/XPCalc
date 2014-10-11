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
	
	// Left Top Width Height
	gap_x := 2
	gap_y := 2
	pos_x := gap_x
	pos_y := gap_y
	add_w := 0
	add_h := 0
	line_pos := 0
	
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
	
	pos_x = gap_x; pos_y = gap_y; add_w = 480; add_h = 32
	// textEdit
	if w.ui.textEdit, err = walk.NewTextEdit(w); err != nil {
		return err
	}
	w.ui.textEdit.SetName("textEdit")
	if err := w.ui.textEdit.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	
	pos_x = gap_x; pos_y += gap_y + add_h; add_w = 270; add_h = 32
	line_pos = pos_y
	// gBoxNumberNotation
	if w.ui.gBoxNumberNotation, err = walk.NewGroupBox(w); err != nil {
		return err
	}
	w.ui.gBoxNumberNotation.SetName("gBoxNumberNotation")
	if err := w.ui.gBoxNumberNotation.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	//if err := w.ui.gBoxNumberNotation.SetTitle(nil); err != nil {
	//	return err
	//}
	
	pos_x = gap_x; pos_y = gap_y * 6; add_w = (add_w - (gap_y * 6))/4; add_h = 18
	// rbNumberNotation_16
	if w.ui.rbNumberNotation_16, err = walk.NewRadioButton(w.ui.gBoxNumberNotation); err != nil {
		return err
	}
	w.ui.rbNumberNotation_16.SetName("rbNumberNotation_16")
	if err := w.ui.rbNumberNotation_16.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.rbNumberNotation_16.SetText(`16bin`); err != nil {
		return err
	}
	
	pos_x += gap_x + add_w; pos_y += 0; add_w += 0; add_h += 0
	// rbNumberNotation_10
	if w.ui.rbNumberNotation_10, err = walk.NewRadioButton(w.ui.gBoxNumberNotation); err != nil {
		return err
	}
	w.ui.rbNumberNotation_10.SetName("rbNumberNotation_10")
	if err := w.ui.rbNumberNotation_10.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.rbNumberNotation_10.SetText(`10bin`); err != nil {
		return err
	}
	w.ui.rbNumberNotation_10.SetChecked(true)
	
	pos_x += gap_x + add_w; pos_y += 0; add_w += 0; add_h += 0
	// rbNumberNotation_08
	if w.ui.rbNumberNotation_08, err = walk.NewRadioButton(w.ui.gBoxNumberNotation); err != nil {
		return err
	}
	w.ui.rbNumberNotation_08.SetName("rbNumberNotation_08")
	if err := w.ui.rbNumberNotation_08.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.rbNumberNotation_08.SetText(`08bin`); err != nil {
		return err
	}
	
	pos_x += gap_x + add_w; pos_y += 0; add_w += 0; add_h += 0
	// rbNumberNotation_02
	if w.ui.rbNumberNotation_02, err = walk.NewRadioButton(w.ui.gBoxNumberNotation); err != nil {
		return err
	}
	w.ui.rbNumberNotation_02.SetName("rbNumberNotation_02")
	if err := w.ui.rbNumberNotation_02.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.rbNumberNotation_02.SetText(`02bin`); err != nil {
		return err
	}
	
	pos_x += add_w + gap_x * 10; pos_y = line_pos; add_w = 270; add_h = 32
	// gBoxNumberUnitBytes
	if w.ui.gBoxNumberUnitBytes, err = walk.NewGroupBox(w); err != nil {
		return err
	}
	w.ui.gBoxNumberNotation.SetName("gBoxNumberUnitBytes")
	if err := w.ui.gBoxNumberUnitBytes.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	//if err := w.ui.gBoxNumberUnitBytes.SetTitle(nil); err != nil {
	//	return err
	//}
	
	pos_x = gap_x; pos_y = gap_y * 6; add_w = (add_w - (gap_y * 6))/4; add_h = 18
	// rbNumberUnitBytes_08
	if w.ui.rbNumberUnitBytes_08, err = walk.NewRadioButton(w.ui.gBoxNumberUnitBytes); err != nil {
		return err
	}
	w.ui.rbNumberUnitBytes_08.SetName("rbNumberUnitBytes_08")
	if err := w.ui.rbNumberUnitBytes_08.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.rbNumberUnitBytes_08.SetText(`ang`); err != nil {
		return err
	}
	w.ui.rbNumberUnitBytes_08.SetChecked(true)
	
	pos_x += gap_x + add_w; pos_y += 0; add_w += 0; add_h += 0
	// rbNumberUnitBytes_04
	if w.ui.rbNumberUnitBytes_04, err = walk.NewRadioButton(w.ui.gBoxNumberUnitBytes); err != nil {
		return err
	}
	w.ui.rbNumberUnitBytes_04.SetName("rbNumberUnitBytes_04")
	if err := w.ui.rbNumberUnitBytes_04.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.rbNumberUnitBytes_04.SetText(`rad`); err != nil {
		return err
	}
	
	pos_x += gap_x + add_w; pos_y += 0; add_w += 0; add_h += 0
	// rbNumberUnitBytes_02
	if w.ui.rbNumberUnitBytes_02, err = walk.NewRadioButton(w.ui.gBoxNumberUnitBytes); err != nil {
		return err
	}
	w.ui.rbNumberUnitBytes_02.SetName("rbNumberUnitBytes_02")
	if err := w.ui.rbNumberUnitBytes_02.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.rbNumberUnitBytes_02.SetText(`gra`); err != nil {
		return err
	}
	
	pos_x += gap_x + add_w; pos_y += 0; add_w += 0; add_h += 0
	// rbNumberUnitBytes_01
	if w.ui.rbNumberUnitBytes_01, err = walk.NewRadioButton(w.ui.gBoxNumberUnitBytes); err != nil {
		return err
	}
	w.ui.rbNumberUnitBytes_01.SetName("rbNumberUnitBytes_01")
	if err := w.ui.rbNumberUnitBytes_01.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.rbNumberUnitBytes_01.SetText(`Byte`); err != nil {
		return err
	}
	w.ui.rbNumberUnitBytes_01.SetVisible(false)
	
	line_pos += 32
	add_w = 0

	pos_x += add_w + gap_x * 10; pos_y = line_pos; add_w = 270; add_h = 32
	// gBoxFunctionOps
	if w.ui.gBoxFunctionOps, err = walk.NewGroupBox(w); err != nil {
		return err
	}
	w.ui.gBoxFunctionOps.SetName("gBoxFunctionOps")
	if err := w.ui.gBoxFunctionOps.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	//if err := w.ui.gBoxFunctionOps.SetTitle(nil); err != nil {
	//	return err
	//}
	
	pos_x += gap_x; pos_y = gap_y * 6; add_w = (add_w - (gap_y * 6))/4; add_h = 18
	// tbBackspace
	if w.ui.tbBackspace, err = walk.NewToolButton(w.ui.gBoxFunctionOps); err != nil {
		return err
	}
	w.ui.tbBackspace.SetName("tbBackspace")
	if err := w.ui.tbBackspace.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbBackspace.SetText(`Backspace`); err != nil {
		return err
	}
	
	pos_x += gap_x + add_w; pos_y += 0; add_w += 0; add_h += 0
	// tbClearError
	if w.ui.tbClearError, err = walk.NewToolButton(w.ui.gBoxFunctionOps); err != nil {
		return err
	}
	w.ui.tbClearError.SetName("tbClearError")
	if err := w.ui.tbClearError.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbClearError.SetText(`CE`); err != nil {
		return err
	}
	
	pos_x += gap_x + add_w; pos_y += 0; add_w += 0; add_h += 0
	// tbClear
	if w.ui.tbClear, err = walk.NewToolButton(w.ui.gBoxFunctionOps); err != nil {
		return err
	}
	w.ui.tbClear.SetName("tbClear")
	if err := w.ui.tbClear.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbClear.SetText(`C`); err != nil {
		return err
	}
	
	succeeded = true

	return nil
}
