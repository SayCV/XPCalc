package main

import (
	"fmt"
	//"log"
	. "math"
	//"math/rand"
	"bytes"
	//"strings"
	//"time"
	"strconv"
)

import (
	"github.com/lxn/walk"
)

// Constants for XPCalcOps.flag
const (
	OPS_NONE        = 1
	OPS_OBJ         = 2
	OPS_ACT         = 4
)

// Constants for XPCalcOps.notation
const (
	OPS_NOTATION_NONE       = 0
	OPS_NOTATION_16         = 1
	OPS_NOTATION_10         = 2
	OPS_NOTATION_08         = 3
	OPS_NOTATION_02         = 4
)

// Constants for XPCalcOps.unit_bytes
const (
	OPS_UNIT_BYTES_NONE       = 0
	OPS_UNIT_BYTES_08         = 1
	OPS_UNIT_BYTES_04         = 2
	OPS_UNIT_BYTES_02         = 3
	OPS_UNIT_BYTES_01         = 4
)

type XPCalcOps struct {
	val_1st    string
	val_2nd    string
	ops        string
	flag       uint32
	notation_old   int
	notation_new   int
	unit_bytes_old int
	unit_bytes_new int
}

type Dialog struct {
	*walk.Dialog
	ui dialogUI
	ops XPCalcOps
}

func opsRun(dlg *Dialog) {
	//dlg.ops.result = dlg.ui.textEdit.Text()
	
	val_1st, _ := strconv.ParseFloat(dlg.ops.val_1st, 64)
	val_2nd, _ := strconv.ParseFloat(dlg.ops.val_2nd, 64)
	var result float64
	
	result = val_2nd
	switch dlg.ops.ops {
		case "+":
			result = val_1st + val_2nd
		case "-":
			result = val_1st - val_2nd
		case "*":
			result = val_1st * val_2nd
		case "/":
			result = val_1st / val_2nd
		case "POWER":
			result = Pow(val_1st, val_2nd)
	}
	dlg.ops.flag = OPS_ACT
	val := strconv.FormatFloat(result, 'g', 'e', 64)
	dlg.ui.textEdit.SetText(val)
	dlg.ops.val_1st = val
	
	fmt.Printf("val_1st = %.6f\n", val_1st)
	fmt.Printf("ops = %s\n", dlg.ops.ops)
	fmt.Printf("val_2nd = %.6f\n", val_2nd)
	fmt.Printf("result = %s\n", val)
}

func rbNumberNotation_onCliced_btnInit(dlg *Dialog) {
	enable := dlg.ui.rbNumberNotation_10.Checked()
	
	if enable {
		dlg.ui.rbNumberUnitBytes_08.SetText(`ang`)
		dlg.ui.rbNumberUnitBytes_04.SetText(`rad`)
		dlg.ui.rbNumberUnitBytes_02.SetText(`gar`)
		dlg.ui.rbNumberUnitBytes_01.SetText(``)
		dlg.ui.rbNumberUnitBytes_01.SetVisible(false)
	} else {
		dlg.ui.rbNumberUnitBytes_08.SetText(`8Bytes`)
		dlg.ui.rbNumberUnitBytes_04.SetText(`4Bytes`)
		dlg.ui.rbNumberUnitBytes_02.SetText(`2Bytes`)
		dlg.ui.rbNumberUnitBytes_01.SetText(`1Bytes`)
		dlg.ui.rbNumberUnitBytes_01.SetVisible(true)
	}

	dlg.ui.tbFnFsubE.SetEnabled(enable)
	dlg.ui.tbFnDms.SetEnabled(enable)
	dlg.ui.tbFnSin.SetEnabled(enable)
	dlg.ui.tbFnCos.SetEnabled(enable)
	dlg.ui.tbFnTan.SetEnabled(enable)
	dlg.ui.tbFnExp.SetEnabled(enable)
	dlg.ui.tbFnPI.SetEnabled(enable)
	
	enable = dlg.ui.rbNumberNotation_16.Checked()
		
	dlg.ui.tbCharA.SetEnabled(enable)
	dlg.ui.tbCharB.SetEnabled(enable)
	dlg.ui.tbCharC.SetEnabled(enable)
	dlg.ui.tbCharD.SetEnabled(enable)
	dlg.ui.tbCharE.SetEnabled(enable)
	dlg.ui.tbCharF.SetEnabled(enable)

	enable = dlg.ui.rbNumberNotation_08.Checked()
	dlg.ui.tbNumber8.SetEnabled(!enable)
	dlg.ui.tbNumber9.SetEnabled(!enable)
	
	enable = dlg.ui.rbNumberNotation_02.Checked()
	dlg.ui.tbNumber2.SetEnabled(!enable)
	dlg.ui.tbNumber3.SetEnabled(!enable)
	dlg.ui.tbNumber4.SetEnabled(!enable)
	dlg.ui.tbNumber5.SetEnabled(!enable)
	dlg.ui.tbNumber6.SetEnabled(!enable)
	dlg.ui.tbNumber7.SetEnabled(!enable)
	dlg.ui.tbNumber8.SetEnabled(!enable)
	dlg.ui.tbNumber9.SetEnabled(!enable)
	
	if dlg.ui.rbNumberNotation_16.Checked() {
		dlg.ops.notation_new = OPS_NOTATION_16
	} else if dlg.ui.rbNumberNotation_10.Checked() {
		dlg.ops.notation_new = OPS_NOTATION_10
	} else if dlg.ui.rbNumberNotation_08.Checked() {
		dlg.ops.notation_new = OPS_NOTATION_08
	} else if dlg.ui.rbNumberNotation_02.Checked() {
		dlg.ops.notation_new = OPS_NOTATION_02
	} else {
		dlg.ops.notation_new = OPS_NOTATION_NONE
	}
	
	var result string
	
	result = dlg.ui.textEdit.Text()
	i, err := strconv.ParseInt(result, dlg.ops.notation_old, 64)
	if err != nil {
		panic(err)
	}
	result = strconv.FormatInt(i, dlg.ops.notation_new)
	dlg.ui.textEdit.SetText(result)
}

func appendByte(dlg *Dialog, c byte) {
	var buffer bytes.Buffer
	var result string
	
	if dlg.ops.flag == OPS_ACT {
		result = ""
	} else {
		result = dlg.ui.textEdit.Text()
	}
	
	dlg.ops.flag = OPS_OBJ
	
	if result == "0" {
		result = string(c)
	} else {
		//result += string(c)
		buffer.WriteString(result)
		buffer.WriteString(string(c))
		result = buffer.String()
	}
	dlg.ui.textEdit.SetText(result)
	dlg.ops.flag = OPS_OBJ
	dlg.ops.val_2nd = result
}

func runDialog(owner walk.Form) (int, error) {
	dlg := new(Dialog)
	if err := dlg.init(owner); err != nil {
		return 0, err
	}
	
	dlg.ops.notation_old   = OPS_NOTATION_10
	dlg.ops.notation_new   = dlg.ops.notation_old
	dlg.ops.unit_bytes_old = OPS_UNIT_BYTES_08
	dlg.ops.unit_bytes_new = dlg.ops.unit_bytes_old
	dlg.ops.flag       = OPS_NONE
	dlg.ops.val_1st    = "0"
	dlg.ops.val_2nd    = ""
	bStatisticalOpen  := false
	
	// TODO: Do further required setup, e.g. for event handling, here.
	
	dlg.ui.rbNumberNotation_16.Clicked().Attach(func() {
		fmt.Println("Clicked rbNumberNotation_16")
		rbNumberNotation_onCliced_btnInit(dlg)
	})
	dlg.ui.rbNumberNotation_10.Clicked().Attach(func() {
		fmt.Println("Clicked rbNumberNotation_10")
		rbNumberNotation_onCliced_btnInit(dlg)
	})
	dlg.ui.rbNumberNotation_08.Clicked().Attach(func() {
		fmt.Println("Clicked rbNumberNotation_08")
		rbNumberNotation_onCliced_btnInit(dlg)
	})
	dlg.ui.rbNumberNotation_02.Clicked().Attach(func() {
		fmt.Println("Clicked rbNumberNotation_02")
		rbNumberNotation_onCliced_btnInit(dlg)
	})
	
	dlg.ui.tbBackspace.Clicked().Attach(func() {
		fmt.Println("Clicked tbBackspace")
		
		// Removed last character of a string
		result := dlg.ui.textEdit.Text()
		sz := len(result)
		if sz > 0 {
    	result = result[:sz-1]
		}
		if len(result)==0 {
			result = "0"
		}
		dlg.ui.textEdit.SetText(result)
	})
	dlg.ui.tbClearError.Clicked().Attach(func() {
		fmt.Println("Clicked tbClearError")
		dlg.ui.textEdit.SetText(`0`)
		dlg.ops.val_1st = dlg.ui.textEdit.Text()
		dlg.ops.val_2nd = ""
		dlg.ops.ops     = ""
	})
	dlg.ui.tbClear.Clicked().Attach(func() {
		fmt.Println("Clicked tbClear")
		dlg.ui.textEdit.SetText(`0`)
		dlg.ops.val_1st = dlg.ui.textEdit.Text()
		dlg.ops.val_2nd = ""
		dlg.ops.ops     = ""
	})
	
	dlg.ui.tbNumber0.Clicked().Attach(func() {
		fmt.Println("Clicked tbNumber0")
		appendByte(dlg, byte('0'))
	})
	dlg.ui.tbNumber1.Clicked().Attach(func() {
		fmt.Println("Clicked tbNumber1")
		appendByte(dlg, byte('1'))
	})
	dlg.ui.tbNumber2.Clicked().Attach(func() {
		fmt.Println("Clicked tbNumber2")
		appendByte(dlg, byte('2'))
	})
	dlg.ui.tbNumber3.Clicked().Attach(func() {
		fmt.Println("Clicked tbNumber3")
		appendByte(dlg, byte('3'))
	})
	dlg.ui.tbNumber4.Clicked().Attach(func() {
		fmt.Println("Clicked tbNumber4")
		appendByte(dlg, byte('4'))
	})
	dlg.ui.tbNumber5.Clicked().Attach(func() {
		fmt.Println("Clicked tbNumber5")
		appendByte(dlg, byte('5'))
	})
	dlg.ui.tbNumber6.Clicked().Attach(func() {
		fmt.Println("Clicked tbNumber5")
		appendByte(dlg, byte('6'))
	})
	dlg.ui.tbNumber7.Clicked().Attach(func() {
		fmt.Println("Clicked tbNumber7")
		appendByte(dlg, byte('7'))
	})
	dlg.ui.tbNumber8.Clicked().Attach(func() {
		fmt.Println("Clicked tbNumber8")
		appendByte(dlg, byte('8'))
	})
	dlg.ui.tbNumber9.Clicked().Attach(func() {
		fmt.Println("Clicked tbNumber9")
		appendByte(dlg, byte('9'))
	})
	dlg.ui.tbCharA.Clicked().Attach(func() {
		fmt.Println("Clicked tbCharA")
		appendByte(dlg, byte('A'))
	})
	dlg.ui.tbCharB.Clicked().Attach(func() {
		fmt.Println("Clicked tbCharB")
		appendByte(dlg, byte('B'))
	})
	dlg.ui.tbCharC.Clicked().Attach(func() {
		fmt.Println("Clicked tbCharC")
		appendByte(dlg, byte('C'))
	})
	dlg.ui.tbCharD.Clicked().Attach(func() {
		fmt.Println("Clicked tbCharD")
		appendByte(dlg, byte('D'))
	})
	dlg.ui.tbCharE.Clicked().Attach(func() {
		fmt.Println("Clicked tbCharE")
		appendByte(dlg, byte('E'))
	})
	dlg.ui.tbCharF.Clicked().Attach(func() {
		fmt.Println("Clicked tbCharF")
		appendByte(dlg, byte('F'))
	})
	
	
	dlg.ui.tbStatistical.Clicked().Attach(func() {
		bStatisticalOpen = !bStatisticalOpen
		
		fmt.Println("Clicked tbStatistical")
		dlg.ui.tbAverage.SetEnabled(bStatisticalOpen)
		dlg.ui.tbSum.SetEnabled(bStatisticalOpen)
		dlg.ui.tbs.SetEnabled(bStatisticalOpen)
		dlg.ui.tbDat.SetEnabled(bStatisticalOpen)
	})
	dlg.ui.tbFnXpower3.Clicked().Attach(func() {
		fmt.Println("Clicked tbFnXpower3")
		dlg.ops.ops = "POWER"
		dlg.ops.val_1st = dlg.ui.textEdit.Text()
		dlg.ops.val_2nd = "3"
		opsRun(dlg)
	})
	dlg.ui.tbFnXpower2.Clicked().Attach(func() {
		fmt.Println("Clicked tbFnXpower2")
		dlg.ops.ops = "POWER"
		dlg.ops.val_1st = dlg.ui.textEdit.Text()
		dlg.ops.val_2nd = "2"
		opsRun(dlg)
	})
	dlg.ui.tbFnPI.Clicked().Attach(func() {
		fmt.Println("Clicked tbFnPI")
		dlg.ui.textEdit.SetText(`3.1415926`)
	})
	
	dlg.ui.tbOpsAdd.Clicked().Attach(func() {
		fmt.Println("Clicked tbOpsAdd")
		//dlg.ui.textEdit.SetText(``)
		dlg.ops.ops = "+"
		//opsRun(dlg)
	})
	dlg.ui.tbOpsSub.Clicked().Attach(func() {
		fmt.Println("Clicked tbOpsSub")
		dlg.ops.ops = "-"
		//opsRun(dlg)
	})
	dlg.ui.tbOpsMul.Clicked().Attach(func() {
		fmt.Println("Clicked tbOpsMul")
		dlg.ops.ops = "*"
		//opsRun(dlg)
	})
	dlg.ui.tbOpsDiv.Clicked().Attach(func() {
		fmt.Println("Clicked tbOpsDiv")
		dlg.ops.ops = "/"
		//opsRun(dlg)
	})
	dlg.ui.tbOpsMod.Clicked().Attach(func() {
		fmt.Println("Clicked tbOpsMod")
		dlg.ops.ops = "/"
		//opsRun(dlg)
	})
	dlg.ui.tbOpsPolar.Clicked().Attach(func() {
		fmt.Println("Clicked tbOpsPolar")
		dlg.ops.ops = "/"
		//opsRun(dlg)
	})
	dlg.ui.tbOpsEqual.Clicked().Attach(func() {
		fmt.Println("Clicked tbOpsEqual")
		opsRun(dlg)
	})
	
	return dlg.Run(), nil
}
