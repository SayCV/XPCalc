// This file was created by ui2walk and may be regenerated.
// DO NOT EDIT OR YOUR MODIFICATIONS WILL BE LOST!

package main

import (
	"github.com/lxn/walk"
	//. "github.com/lxn/walk/declarative"
)

type dialogUI struct {
	windowIcon *walk.Icon
	
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
	
	gBoxFnOps								*walk.GroupBox
	cbInverseFunction				*walk.CheckBox
	cbHyperbolicFunction		*walk.CheckBox
	
	gBoxActOps							*walk.GroupBox
	tbBackspace							*walk.ToolButton
	tbClearError						*walk.ToolButton
	tbClear									*walk.ToolButton
	
	gBoxFnsOps									*walk.GroupBox
	tbStatistical								*walk.ToolButton
	tbAverage										*walk.ToolButton
	tbSum												*walk.ToolButton
	tbs													*walk.ToolButton
	tbDat												*walk.ToolButton
	
	tbFnFsubE										*walk.ToolButton
	tbFnDms											*walk.ToolButton
	tbFnSin											*walk.ToolButton
	tbFnCos											*walk.ToolButton
	tbFnTan											*walk.ToolButton
	tbFnLeftBracket							*walk.ToolButton
	tbFnRightBracket						*walk.ToolButton
	tbFnExp											*walk.ToolButton
	tbFnLn											*walk.ToolButton
	tbFnXpowerY									*walk.ToolButton
	tbFnLog											*walk.ToolButton
	tbFnXpower3									*walk.ToolButton
	tbFnXpower2									*walk.ToolButton
	tbFnNFactorial							*walk.ToolButton
	tbFnXReciprocal							*walk.ToolButton
	
	tbFnMC											*walk.ToolButton
	tbFnMR											*walk.ToolButton
	tbFnMS											*walk.ToolButton
	tbFnMA											*walk.ToolButton
	tbFnPI											*walk.ToolButton
	
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
	row_pos := 0
	
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
	if err := w.SetClientSize(walk.Size{560, 300}); err != nil {
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
	row_pos = pos_x
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
	
	pos_x = gap_x; pos_y = line_pos; add_w = 270/2; add_h = 32
	// gBoxFnOps
	if w.ui.gBoxFnOps, err = walk.NewGroupBox(w); err != nil {
		return err
	}
	w.ui.gBoxFnOps.SetName("gBoxFnOps")
	if err := w.ui.gBoxFnOps.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	//if err := w.ui.gBoxFnOps.SetTitle(nil); err != nil {
	//	return err
	//}
	
	pos_x = gap_x; pos_y = gap_y * 6; add_w = (add_w - (gap_y * 2))/2; add_h = 18
	// cbInverseFunction
	if w.ui.cbInverseFunction, err = walk.NewCheckBox(w.ui.gBoxFnOps); err != nil {
		return err
	}
	w.ui.cbInverseFunction.SetName("cbInverseFunction")
	if err := w.ui.cbInverseFunction.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.cbInverseFunction.SetText(`Inv`); err != nil {
		return err
	}
	
	pos_x += gap_x + add_w; pos_y += 0; add_w += 0; add_h += 0
	// cbHyperbolicFunction
	if w.ui.cbHyperbolicFunction, err = walk.NewCheckBox(w.ui.gBoxFnOps); err != nil {
		return err
	}
	w.ui.cbHyperbolicFunction.SetName("cbHyperbolicFunction")
	if err := w.ui.cbHyperbolicFunction.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.cbHyperbolicFunction.SetText(`Hyp`); err != nil {
		return err
	}
	
	pos_x = row_pos; pos_y = line_pos; add_w = 270; add_h = 32
	// gBoxActOps
	//if w.ui.gBoxActOps, err = walk.NewGroupBox(w); err != nil {
	//	return err
	//}
	//w.ui.gBoxActOps.SetName("gBoxActOps")
	//if err := w.ui.gBoxActOps.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
	//	return err
	//}
	//if err := w.ui.gBoxActOps.SetTitle(nil); err != nil {
	//	return err
	//}
	//w.ui.gBoxActOps.SetVisible(false)
	
	pos_x += gap_x; pos_y += gap_y; add_w = (add_w - (gap_y * 6))/3; add_h = 32
	// tbBackspace
	if w.ui.tbBackspace, err = walk.NewToolButton(w); err != nil {
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
	if w.ui.tbClearError, err = walk.NewToolButton(w); err != nil {
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
	if w.ui.tbClear, err = walk.NewToolButton(w); err != nil {
		return err
	}
	w.ui.tbClear.SetName("tbClear")
	if err := w.ui.tbClear.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbClear.SetText(`C`); err != nil {
		return err
	}
	
	line_pos += 32 + 12
	line_pos_4 := line_pos
	add_w = 0
	pos_x = row_pos; pos_y = line_pos; add_w = 270; add_h = 32
	// gBoxMainOps

	pos_x += gap_x; pos_y += gap_y; add_w = (add_w - (gap_y * 6))/6; add_h = 32
	// tbNumber7
	if w.ui.tbNumber7, err = walk.NewToolButton(w); err != nil {
		return err
	}
	w.ui.tbNumber7.SetName("tbNumber7")
	if err := w.ui.tbNumber7.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbNumber7.SetText(`7`); err != nil {
		return err
	}
	
	pos_x += gap_x + add_w; pos_y += 0; add_w += 0; add_h += 0
	// tbNumber8
	if w.ui.tbNumber8, err = walk.NewToolButton(w); err != nil {
		return err
	}
	w.ui.tbNumber8.SetName("tbNumber8")
	if err := w.ui.tbNumber8.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbNumber8.SetText(`8`); err != nil {
		return err
	}
	
	pos_x += gap_x + add_w; pos_y += 0; add_w += 0; add_h += 0
	// tbNumber9
	if w.ui.tbNumber9, err = walk.NewToolButton(w); err != nil {
		return err
	}
	w.ui.tbNumber9.SetName("tbNumber9")
	if err := w.ui.tbNumber9.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbNumber9.SetText(`9`); err != nil {
		return err
	}
	
	pos_x += gap_x + add_w; pos_y += 0; add_w += 0; add_h += 0
	// tbOpsDiv
	if w.ui.tbOpsDiv, err = walk.NewToolButton(w); err != nil {
		return err
	}
	w.ui.tbOpsDiv.SetName("tbOpsDiv")
	if err := w.ui.tbOpsDiv.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbOpsDiv.SetText(`/`); err != nil {
		return err
	}
	
	pos_x += gap_x + add_w; pos_y += 0; add_w += 0; add_h += 0
	// tbOpsMod
	if w.ui.tbOpsMod, err = walk.NewToolButton(w); err != nil {
		return err
	}
	w.ui.tbOpsMod.SetName("tbOpsMod")
	if err := w.ui.tbOpsMod.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbOpsMod.SetText(`Mod`); err != nil {
		return err
	}
	
	pos_x += gap_x + add_w; pos_y += 0; add_w += 0; add_h += 0
	// tbOpsAnd
	if w.ui.tbOpsAnd, err = walk.NewToolButton(w); err != nil {
		return err
	}
	w.ui.tbOpsAnd.SetName("tbOpsAnd")
	if err := w.ui.tbOpsAnd.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbOpsAnd.SetText(`And`); err != nil {
		return err
	}
	
	line_pos += 32 + gap_y
	pos_x = row_pos; pos_y = line_pos; add_w = 270; add_h = 32
	pos_x += gap_x; pos_y += gap_y; add_w = (add_w - (gap_y * 6))/6; add_h = 32
	
	// tbNumber4
	if w.ui.tbNumber4, err = walk.NewToolButton(w); err != nil {
		return err
	}
	w.ui.tbNumber4.SetName("tbNumber4")
	if err := w.ui.tbNumber4.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbNumber4.SetText(`4`); err != nil {
		return err
	}
	
	pos_x += gap_x + add_w; pos_y += 0; add_w += 0; add_h += 0
	// tbNumber5
	if w.ui.tbNumber5, err = walk.NewToolButton(w); err != nil {
		return err
	}
	w.ui.tbNumber5.SetName("tbNumber5")
	if err := w.ui.tbNumber5.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbNumber5.SetText(`5`); err != nil {
		return err
	}
	
	pos_x += gap_x + add_w; pos_y += 0; add_w += 0; add_h += 0
	// tbNumber6
	if w.ui.tbNumber6, err = walk.NewToolButton(w); err != nil {
		return err
	}
	w.ui.tbNumber6.SetName("tbNumber6")
	if err := w.ui.tbNumber6.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbNumber6.SetText(`6`); err != nil {
		return err
	}
	
	pos_x += gap_x + add_w; pos_y += 0; add_w += 0; add_h += 0
	// tbOpsMul
	if w.ui.tbOpsMul, err = walk.NewToolButton(w); err != nil {
		return err
	}
	w.ui.tbOpsMul.SetName("tbOpsMul")
	if err := w.ui.tbOpsMul.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbOpsMul.SetText(`*`); err != nil {
		return err
	}
	
	pos_x += gap_x + add_w; pos_y += 0; add_w += 0; add_h += 0
	// tbOpsOr
	if w.ui.tbOpsOr, err = walk.NewToolButton(w); err != nil {
		return err
	}
	w.ui.tbOpsOr.SetName("tbOpsOr")
	if err := w.ui.tbOpsOr.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbOpsOr.SetText(`Or`); err != nil {
		return err
	}
	
	pos_x += gap_x + add_w; pos_y += 0; add_w += 0; add_h += 0
	// tbOpsXor
	if w.ui.tbOpsXor, err = walk.NewToolButton(w); err != nil {
		return err
	}
	w.ui.tbOpsXor.SetName("tbOpsXor")
	if err := w.ui.tbOpsXor.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbOpsXor.SetText(`Xor`); err != nil {
		return err
	}
	
	line_pos += 32 + gap_y
	pos_x = row_pos; pos_y = line_pos; add_w = 270; add_h = 32
	pos_x += gap_x; pos_y += gap_y; add_w = (add_w - (gap_y * 6))/6; add_h = 32
	
	// tbNumber1
	if w.ui.tbNumber1, err = walk.NewToolButton(w); err != nil {
		return err
	}
	w.ui.tbNumber1.SetName("tbNumber1")
	if err := w.ui.tbNumber1.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbNumber1.SetText(`1`); err != nil {
		return err
	}
	
	pos_x += gap_x + add_w; pos_y += 0; add_w += 0; add_h += 0
	// tbNumber2
	if w.ui.tbNumber2, err = walk.NewToolButton(w); err != nil {
		return err
	}
	w.ui.tbNumber2.SetName("tbNumber2")
	if err := w.ui.tbNumber2.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbNumber2.SetText(`2`); err != nil {
		return err
	}
	
	pos_x += gap_x + add_w; pos_y += 0; add_w += 0; add_h += 0
	// tbNumber3
	if w.ui.tbNumber3, err = walk.NewToolButton(w); err != nil {
		return err
	}
	w.ui.tbNumber3.SetName("tbNumber3")
	if err := w.ui.tbNumber3.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbNumber3.SetText(`3`); err != nil {
		return err
	}
	
	pos_x += gap_x + add_w; pos_y += 0; add_w += 0; add_h += 0
	// tbOpsSub
	if w.ui.tbOpsSub, err = walk.NewToolButton(w); err != nil {
		return err
	}
	w.ui.tbOpsSub.SetName("tbOpsSub")
	if err := w.ui.tbOpsSub.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbOpsSub.SetText(`-`); err != nil {
		return err
	}
	
	pos_x += gap_x + add_w; pos_y += 0; add_w += 0; add_h += 0
	// tbOpsLsh
	if w.ui.tbOpsLsh, err = walk.NewToolButton(w); err != nil {
		return err
	}
	w.ui.tbOpsLsh.SetName("tbOpsLsh")
	if err := w.ui.tbOpsLsh.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbOpsLsh.SetText(`Lsh`); err != nil {
		return err
	}
	
	pos_x += gap_x + add_w; pos_y += 0; add_w += 0; add_h += 0
	// tbOpsNot
	if w.ui.tbOpsNot, err = walk.NewToolButton(w); err != nil {
		return err
	}
	w.ui.tbOpsNot.SetName("tbOpsNot")
	if err := w.ui.tbOpsNot.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbOpsNot.SetText(`Not`); err != nil {
		return err
	}
	
	line_pos += 32 + gap_y
	pos_x = row_pos; pos_y = line_pos; add_w = 270; add_h = 32
	pos_x += gap_x; pos_y += gap_y; add_w = (add_w - (gap_y * 6))/6; add_h = 32
	
	// tbNumber0
	if w.ui.tbNumber0, err = walk.NewToolButton(w); err != nil {
		return err
	}
	w.ui.tbNumber0.SetName("tbNumber0")
	if err := w.ui.tbNumber0.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbNumber0.SetText(`0`); err != nil {
		return err
	}
	
	pos_x += gap_x + add_w; pos_y += 0; add_w += 0; add_h += 0
	// tbOpsPolar
	if w.ui.tbOpsPolar, err = walk.NewToolButton(w); err != nil {
		return err
	}
	w.ui.tbOpsPolar.SetName("tbOpsPolar")
	if err := w.ui.tbOpsPolar.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbOpsPolar.SetText(`+/-`); err != nil {
		return err
	}
	
	pos_x += gap_x + add_w; pos_y += 0; add_w += 0; add_h += 0
	// tbOpsPoint
	if w.ui.tbOpsPoint, err = walk.NewToolButton(w); err != nil {
		return err
	}
	w.ui.tbOpsPoint.SetName("tbOpsPoint")
	if err := w.ui.tbOpsPoint.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbOpsPoint.SetText(`.`); err != nil {
		return err
	}
	
	pos_x += gap_x + add_w; pos_y += 0; add_w += 0; add_h += 0
	// tbOpsAdd
	if w.ui.tbOpsAdd, err = walk.NewToolButton(w); err != nil {
		return err
	}
	w.ui.tbOpsAdd.SetName("tbOpsAdd")
	if err := w.ui.tbOpsAdd.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbOpsAdd.SetText(`+`); err != nil {
		return err
	}
	
	pos_x += gap_x + add_w; pos_y += 0; add_w += 0; add_h += 0
	// tbOpsEqual
	if w.ui.tbOpsEqual, err = walk.NewToolButton(w); err != nil {
		return err
	}
	w.ui.tbOpsEqual.SetName("tbOpsEqual")
	if err := w.ui.tbOpsEqual.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbOpsEqual.SetText(`=`); err != nil {
		return err
	}
	
	pos_x += gap_x + add_w; pos_y += 0; add_w += 0; add_h += 0
	// tbOpsInt
	if w.ui.tbOpsInt, err = walk.NewToolButton(w); err != nil {
		return err
	}
	w.ui.tbOpsInt.SetName("tbOpsInt")
	if err := w.ui.tbOpsInt.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbOpsInt.SetText(`Int`); err != nil {
		return err
	}
	
	line_pos += 32 + gap_y
	pos_x = row_pos; pos_y = line_pos; add_w = 270; add_h = 32
	pos_x += gap_x; pos_y += gap_y; add_w = (add_w - (gap_y * 6))/6; add_h = 32
	
	// tbCharA
	if w.ui.tbCharA, err = walk.NewToolButton(w); err != nil {
		return err
	}
	w.ui.tbCharA.SetName("tbCharA")
	if err := w.ui.tbCharA.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbCharA.SetText(`A`); err != nil {
		return err
	}
	w.ui.tbCharA.SetEnabled(false)
	
	pos_x += gap_x + add_w; pos_y += 0; add_w += 0; add_h += 0
	// tbCharB
	if w.ui.tbCharB, err = walk.NewToolButton(w); err != nil {
		return err
	}
	w.ui.tbCharB.SetName("tbCharB")
	if err := w.ui.tbCharB.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbCharB.SetText(`B`); err != nil {
		return err
	}
	w.ui.tbCharB.SetEnabled(false)
	
	pos_x += gap_x + add_w; pos_y += 0; add_w += 0; add_h += 0
	// tbCharC
	if w.ui.tbCharC, err = walk.NewToolButton(w); err != nil {
		return err
	}
	w.ui.tbCharC.SetName("tbCharC")
	if err := w.ui.tbCharC.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbCharC.SetText(`C`); err != nil {
		return err
	}
	w.ui.tbCharC.SetEnabled(false)
	
	pos_x += gap_x + add_w; pos_y += 0; add_w += 0; add_h += 0
	// tbCharD
	if w.ui.tbCharD, err = walk.NewToolButton(w); err != nil {
		return err
	}
	w.ui.tbCharD.SetName("tbCharD")
	if err := w.ui.tbCharD.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbCharD.SetText(`D`); err != nil {
		return err
	}
	w.ui.tbCharD.SetEnabled(false)
	
	pos_x += gap_x + add_w; pos_y += 0; add_w += 0; add_h += 0
	// tbCharE
	if w.ui.tbCharE, err = walk.NewToolButton(w); err != nil {
		return err
	}
	w.ui.tbCharE.SetName("tbCharE")
	if err := w.ui.tbCharE.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbCharE.SetText(`E`); err != nil {
		return err
	}
	w.ui.tbCharE.SetEnabled(false)
	
	pos_x += gap_x + add_w; pos_y += 0; add_w += 0; add_h += 0
	// tbCharF
	if w.ui.tbCharF, err = walk.NewToolButton(w); err != nil {
		return err
	}
	w.ui.tbCharF.SetName("tbCharF")
	if err := w.ui.tbCharF.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbCharF.SetText(`F`); err != nil {
		return err
	}
	w.ui.tbCharF.SetEnabled(false)
	
	add_w = 0
	line_pos = line_pos_4
	pos_x = gap_x; pos_y = line_pos; add_w = 250; add_h = 32
	pos_x += gap_x; pos_y += gap_y; add_w = (add_w - (gap_y * 6))/6; add_h = 32
	
	// tbStatistical
	if w.ui.tbStatistical, err = walk.NewToolButton(w); err != nil {
		return err
	}
	w.ui.tbStatistical.SetName("tbStatistical")
	if err := w.ui.tbStatistical.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbStatistical.SetText(`Sta`); err != nil {
		return err
	}
	
	pos_x += gap_x + add_w * 3/2; pos_y += 0; add_w += 0; add_h += 0
	// tbFnFsubE
	if w.ui.tbFnFsubE, err = walk.NewToolButton(w); err != nil {
		return err
	}
	w.ui.tbFnFsubE.SetName("tbFnFsubE")
	if err := w.ui.tbFnFsubE.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbFnFsubE.SetText(`F-E`); err != nil {
		return err
	}
	
	pos_x += gap_x + add_w; pos_y += 0; add_w += 0; add_h += 0
	// tbFnLeftBracket
	if w.ui.tbFnLeftBracket, err = walk.NewToolButton(w); err != nil {
		return err
	}
	w.ui.tbFnLeftBracket.SetName("tbFnLeftBracket")
	if err := w.ui.tbFnLeftBracket.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbFnLeftBracket.SetText(`(`); err != nil {
		return err
	}
	
	pos_x += gap_x + add_w; pos_y += 0; add_w += 0; add_h += 0
	// tbFnRightBracket
	if w.ui.tbFnRightBracket, err = walk.NewToolButton(w); err != nil {
		return err
	}
	w.ui.tbFnRightBracket.SetName("tbFnRightBracket")
	if err := w.ui.tbFnRightBracket.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbFnRightBracket.SetText(`)`); err != nil {
		return err
	}
	
	pos_x += gap_x + add_w * 3/2; pos_y += 0; add_w += 0; add_h += 0
	// tbFnMC
	if w.ui.tbFnMC, err = walk.NewToolButton(w); err != nil {
		return err
	}
	w.ui.tbFnMC.SetName("tbFnMC")
	if err := w.ui.tbFnMC.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbFnMC.SetText(`MC`); err != nil {
		return err
	}
	
	line_pos += 32 + gap_y
	pos_x = gap_x; pos_y = line_pos; add_w = 250; add_h = 32
	pos_x += gap_x; pos_y += gap_y; add_w = (add_w - (gap_y * 6))/6; add_h = 32
	
	// tbAverage
	if w.ui.tbAverage, err = walk.NewToolButton(w); err != nil {
		return err
	}
	w.ui.tbAverage.SetName("tbAverage")
	if err := w.ui.tbAverage.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbAverage.SetText(`Ave`); err != nil {
		return err
	}
	w.ui.tbAverage.SetEnabled(false)
	
	pos_x += gap_x + add_w * 3/2; pos_y += 0; add_w += 0; add_h += 0
	// tbFnDms
	if w.ui.tbFnDms, err = walk.NewToolButton(w); err != nil {
		return err
	}
	w.ui.tbFnDms.SetName("tbFnDms")
	if err := w.ui.tbFnDms.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbFnDms.SetText(`dms`); err != nil {
		return err
	}
	
	pos_x += gap_x + add_w; pos_y += 0; add_w += 0; add_h += 0
	// tbFnExp
	if w.ui.tbFnExp, err = walk.NewToolButton(w); err != nil {
		return err
	}
	w.ui.tbFnExp.SetName("tbFnExp")
	if err := w.ui.tbFnExp.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbFnExp.SetText(`Exp`); err != nil {
		return err
	}
	
	pos_x += gap_x + add_w; pos_y += 0; add_w += 0; add_h += 0
	// tbFnLn
	if w.ui.tbFnLn, err = walk.NewToolButton(w); err != nil {
		return err
	}
	w.ui.tbFnLn.SetName("tbFnLn")
	if err := w.ui.tbFnLn.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbFnLn.SetText(`Ln`); err != nil {
		return err
	}
	
	pos_x += gap_x + add_w * 3/2; pos_y += 0; add_w += 0; add_h += 0
	// tbFnMR
	if w.ui.tbFnMR, err = walk.NewToolButton(w); err != nil {
		return err
	}
	w.ui.tbFnMR.SetName("tbFnMR")
	if err := w.ui.tbFnMR.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbFnMR.SetText(`MR`); err != nil {
		return err
	}
	
	line_pos += 32 + gap_y
	pos_x = gap_x; pos_y = line_pos; add_w = 250; add_h = 32
	pos_x += gap_x; pos_y += gap_y; add_w = (add_w - (gap_y * 6))/6; add_h = 32
	
	// tbSum
	if w.ui.tbSum, err = walk.NewToolButton(w); err != nil {
		return err
	}
	w.ui.tbSum.SetName("tbSum")
	if err := w.ui.tbSum.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbSum.SetText(`Sum`); err != nil {
		return err
	}
	w.ui.tbSum.SetEnabled(false)
	
	pos_x += gap_x + add_w * 3/2; pos_y += 0; add_w += 0; add_h += 0
	// tbFnSin
	if w.ui.tbFnSin, err = walk.NewToolButton(w); err != nil {
		return err
	}
	w.ui.tbFnSin.SetName("tbFnSin")
	if err := w.ui.tbFnSin.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbFnSin.SetText(`sin`); err != nil {
		return err
	}
	
	pos_x += gap_x + add_w; pos_y += 0; add_w += 0; add_h += 0
	// tbFnXpowerY
	if w.ui.tbFnXpowerY, err = walk.NewToolButton(w); err != nil {
		return err
	}
	w.ui.tbFnXpowerY.SetName("tbFnXpowerY")
	if err := w.ui.tbFnXpowerY.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbFnXpowerY.SetText(`x^y`); err != nil {
		return err
	}
	
	pos_x += gap_x + add_w; pos_y += 0; add_w += 0; add_h += 0
	// tbFnLog
	if w.ui.tbFnLog, err = walk.NewToolButton(w); err != nil {
		return err
	}
	w.ui.tbFnLog.SetName("tbFnLog")
	if err := w.ui.tbFnLog.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbFnLog.SetText(`Log`); err != nil {
		return err
	}
	
	pos_x += gap_x + add_w * 3/2; pos_y += 0; add_w += 0; add_h += 0
	// tbFnLog
	if w.ui.tbFnLog, err = walk.NewToolButton(w); err != nil {
		return err
	}
	w.ui.tbFnLog.SetName("tbFnLog")
	if err := w.ui.tbFnLog.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbFnLog.SetText(`Log`); err != nil {
		return err
	}
	
	line_pos += 32 + gap_y
	pos_x = gap_x; pos_y = line_pos; add_w = 250; add_h = 32
	pos_x += gap_x; pos_y += gap_y; add_w = (add_w - (gap_y * 6))/6; add_h = 32
	
	// tbs
	if w.ui.tbs, err = walk.NewToolButton(w); err != nil {
		return err
	}
	w.ui.tbs.SetName("tbs")
	if err := w.ui.tbs.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbs.SetText(`s`); err != nil {
		return err
	}
	w.ui.tbs.SetEnabled(false)
	
	pos_x += gap_x + add_w * 3/2; pos_y += 0; add_w += 0; add_h += 0
	// tbFnCos
	if w.ui.tbFnCos, err = walk.NewToolButton(w); err != nil {
		return err
	}
	w.ui.tbFnCos.SetName("tbFnSin")
	if err := w.ui.tbFnCos.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbFnCos.SetText(`cos`); err != nil {
		return err
	}
	
	pos_x += gap_x + add_w; pos_y += 0; add_w += 0; add_h += 0
	// tbFnXpower3
	if w.ui.tbFnXpower3, err = walk.NewToolButton(w); err != nil {
		return err
	}
	w.ui.tbFnXpower3.SetName("tbFnXpower3")
	if err := w.ui.tbFnXpower3.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbFnXpower3.SetText(`x^3`); err != nil {
		return err
	}
	
	pos_x += gap_x + add_w; pos_y += 0; add_w += 0; add_h += 0
	// tbFnNFactorial
	if w.ui.tbFnNFactorial, err = walk.NewToolButton(w); err != nil {
		return err
	}
	w.ui.tbFnNFactorial.SetName("tbFnNFactorial")
	if err := w.ui.tbFnNFactorial.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbFnNFactorial.SetText(`n!`); err != nil {
		return err
	}
	
	pos_x += gap_x + add_w * 3/2; pos_y += 0; add_w += 0; add_h += 0
	// tbFnMA
	if w.ui.tbFnMA, err = walk.NewToolButton(w); err != nil {
		return err
	}
	w.ui.tbFnMA.SetName("tbFnMA")
	if err := w.ui.tbFnMA.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbFnMA.SetText(`M+`); err != nil {
		return err
	}
	
	line_pos += 32 + gap_y
	pos_x = gap_x; pos_y = line_pos; add_w = 250; add_h = 32
	pos_x += gap_x; pos_y += gap_y; add_w = (add_w - (gap_y * 6))/6; add_h = 32
	
	// tbDat
	if w.ui.tbDat, err = walk.NewToolButton(w); err != nil {
		return err
	}
	w.ui.tbDat.SetName("tbDat")
	if err := w.ui.tbDat.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbDat.SetText(`Dat`); err != nil {
		return err
	}
	w.ui.tbDat.SetEnabled(false)
	
	pos_x += gap_x + add_w * 3/2; pos_y += 0; add_w += 0; add_h += 0
	// tbFnTan
	if w.ui.tbFnTan, err = walk.NewToolButton(w); err != nil {
		return err
	}
	w.ui.tbFnTan.SetName("tbFnTan")
	if err := w.ui.tbFnTan.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbFnTan.SetText(`tan`); err != nil {
		return err
	}
	
	pos_x += gap_x + add_w; pos_y += 0; add_w += 0; add_h += 0
	// tbFnXpower2
	if w.ui.tbFnXpower2, err = walk.NewToolButton(w); err != nil {
		return err
	}
	w.ui.tbFnXpower2.SetName("tbFnXpower2")
	if err := w.ui.tbFnXpower2.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbFnXpower2.SetText(`x^2`); err != nil {
		return err
	}
	
	pos_x += gap_x + add_w; pos_y += 0; add_w += 0; add_h += 0
	// tbFnXReciprocal
	if w.ui.tbFnXReciprocal, err = walk.NewToolButton(w); err != nil {
		return err
	}
	w.ui.tbFnXReciprocal.SetName("tbFnXReciprocal")
	if err := w.ui.tbFnXReciprocal.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbFnXReciprocal.SetText(`1/x`); err != nil {
		return err
	}
	
	pos_x += gap_x + add_w * 3/2; pos_y += 0; add_w += 0; add_h += 0
	// tbFnPI
	if w.ui.tbFnPI, err = walk.NewToolButton(w); err != nil {
		return err
	}
	w.ui.tbFnPI.SetName("tbFnPI")
	if err := w.ui.tbFnPI.SetBounds(walk.Rectangle{pos_x, pos_y, add_w, add_h}); err != nil {
		return err
	}
	if err := w.ui.tbFnPI.SetText(`PI`); err != nil {
		return err
	}
	
	
	//ic, err := walk.NewIconFromImage(makeDigitImage(counter))
	//icon, err := walk.NewIconFromFile("./icon/4.ico")
	icon, err := walk.NewIconFromResource("ICON_CALC")
	if err != nil {
		return
	}
	w.SetIcon(icon)
	if w.ui.windowIcon != nil {
		w.ui.windowIcon.Dispose()
	}
	w.ui.windowIcon = icon
	
	succeeded = true

	return nil
}

