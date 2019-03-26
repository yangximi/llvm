// Code generated by "string2enum -linecomment -type DwarfOp ../../ir/enum"; DO NOT EDIT.

package enum

import "fmt"
import "github.com/llir/llvm/ir/enum"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the string2enum command to generate them again.
	var x [1]struct{}
	_ = x[enum.DwarfOpAddr-3]
	_ = x[enum.DwarfOpDeref-6]
	_ = x[enum.DwarfOpConst1u-8]
	_ = x[enum.DwarfOpConst1s-9]
	_ = x[enum.DwarfOpConst2u-10]
	_ = x[enum.DwarfOpConst2s-11]
	_ = x[enum.DwarfOpConst4u-12]
	_ = x[enum.DwarfOpConst4s-13]
	_ = x[enum.DwarfOpConst8u-14]
	_ = x[enum.DwarfOpConst8s-15]
	_ = x[enum.DwarfOpConstu-16]
	_ = x[enum.DwarfOpConsts-17]
	_ = x[enum.DwarfOpDup-18]
	_ = x[enum.DwarfOpDrop-19]
	_ = x[enum.DwarfOpOver-20]
	_ = x[enum.DwarfOpPick-21]
	_ = x[enum.DwarfOpSwap-22]
	_ = x[enum.DwarfOpRot-23]
	_ = x[enum.DwarfOpXderef-24]
	_ = x[enum.DwarfOpAbs-25]
	_ = x[enum.DwarfOpAnd-26]
	_ = x[enum.DwarfOpDiv-27]
	_ = x[enum.DwarfOpMinus-28]
	_ = x[enum.DwarfOpMod-29]
	_ = x[enum.DwarfOpMul-30]
	_ = x[enum.DwarfOpNeg-31]
	_ = x[enum.DwarfOpNot-32]
	_ = x[enum.DwarfOpOr-33]
	_ = x[enum.DwarfOpPlus-34]
	_ = x[enum.DwarfOpPlusUconst-35]
	_ = x[enum.DwarfOpShl-36]
	_ = x[enum.DwarfOpShr-37]
	_ = x[enum.DwarfOpShra-38]
	_ = x[enum.DwarfOpXor-39]
	_ = x[enum.DwarfOpBra-40]
	_ = x[enum.DwarfOpEq-41]
	_ = x[enum.DwarfOpGe-42]
	_ = x[enum.DwarfOpGt-43]
	_ = x[enum.DwarfOpLe-44]
	_ = x[enum.DwarfOpLt-45]
	_ = x[enum.DwarfOpNe-46]
	_ = x[enum.DwarfOpSkip-47]
	_ = x[enum.DwarfOpLit0-48]
	_ = x[enum.DwarfOpLit1-49]
	_ = x[enum.DwarfOpLit2-50]
	_ = x[enum.DwarfOpLit3-51]
	_ = x[enum.DwarfOpLit4-52]
	_ = x[enum.DwarfOpLit5-53]
	_ = x[enum.DwarfOpLit6-54]
	_ = x[enum.DwarfOpLit7-55]
	_ = x[enum.DwarfOpLit8-56]
	_ = x[enum.DwarfOpLit9-57]
	_ = x[enum.DwarfOpLit10-58]
	_ = x[enum.DwarfOpLit11-59]
	_ = x[enum.DwarfOpLit12-60]
	_ = x[enum.DwarfOpLit13-61]
	_ = x[enum.DwarfOpLit14-62]
	_ = x[enum.DwarfOpLit15-63]
	_ = x[enum.DwarfOpLit16-64]
	_ = x[enum.DwarfOpLit17-65]
	_ = x[enum.DwarfOpLit18-66]
	_ = x[enum.DwarfOpLit19-67]
	_ = x[enum.DwarfOpLit20-68]
	_ = x[enum.DwarfOpLit21-69]
	_ = x[enum.DwarfOpLit22-70]
	_ = x[enum.DwarfOpLit23-71]
	_ = x[enum.DwarfOpLit24-72]
	_ = x[enum.DwarfOpLit25-73]
	_ = x[enum.DwarfOpLit26-74]
	_ = x[enum.DwarfOpLit27-75]
	_ = x[enum.DwarfOpLit28-76]
	_ = x[enum.DwarfOpLit29-77]
	_ = x[enum.DwarfOpLit30-78]
	_ = x[enum.DwarfOpLit31-79]
	_ = x[enum.DwarfOpReg0-80]
	_ = x[enum.DwarfOpReg1-81]
	_ = x[enum.DwarfOpReg2-82]
	_ = x[enum.DwarfOpReg3-83]
	_ = x[enum.DwarfOpReg4-84]
	_ = x[enum.DwarfOpReg5-85]
	_ = x[enum.DwarfOpReg6-86]
	_ = x[enum.DwarfOpReg7-87]
	_ = x[enum.DwarfOpReg8-88]
	_ = x[enum.DwarfOpReg9-89]
	_ = x[enum.DwarfOpReg10-90]
	_ = x[enum.DwarfOpReg11-91]
	_ = x[enum.DwarfOpReg12-92]
	_ = x[enum.DwarfOpReg13-93]
	_ = x[enum.DwarfOpReg14-94]
	_ = x[enum.DwarfOpReg15-95]
	_ = x[enum.DwarfOpReg16-96]
	_ = x[enum.DwarfOpReg17-97]
	_ = x[enum.DwarfOpReg18-98]
	_ = x[enum.DwarfOpReg19-99]
	_ = x[enum.DwarfOpReg20-100]
	_ = x[enum.DwarfOpReg21-101]
	_ = x[enum.DwarfOpReg22-102]
	_ = x[enum.DwarfOpReg23-103]
	_ = x[enum.DwarfOpReg24-104]
	_ = x[enum.DwarfOpReg25-105]
	_ = x[enum.DwarfOpReg26-106]
	_ = x[enum.DwarfOpReg27-107]
	_ = x[enum.DwarfOpReg28-108]
	_ = x[enum.DwarfOpReg29-109]
	_ = x[enum.DwarfOpReg30-110]
	_ = x[enum.DwarfOpReg31-111]
	_ = x[enum.DwarfOpBreg0-112]
	_ = x[enum.DwarfOpBreg1-113]
	_ = x[enum.DwarfOpBreg2-114]
	_ = x[enum.DwarfOpBreg3-115]
	_ = x[enum.DwarfOpBreg4-116]
	_ = x[enum.DwarfOpBreg5-117]
	_ = x[enum.DwarfOpBreg6-118]
	_ = x[enum.DwarfOpBreg7-119]
	_ = x[enum.DwarfOpBreg8-120]
	_ = x[enum.DwarfOpBreg9-121]
	_ = x[enum.DwarfOpBreg10-122]
	_ = x[enum.DwarfOpBreg11-123]
	_ = x[enum.DwarfOpBreg12-124]
	_ = x[enum.DwarfOpBreg13-125]
	_ = x[enum.DwarfOpBreg14-126]
	_ = x[enum.DwarfOpBreg15-127]
	_ = x[enum.DwarfOpBreg16-128]
	_ = x[enum.DwarfOpBreg17-129]
	_ = x[enum.DwarfOpBreg18-130]
	_ = x[enum.DwarfOpBreg19-131]
	_ = x[enum.DwarfOpBreg20-132]
	_ = x[enum.DwarfOpBreg21-133]
	_ = x[enum.DwarfOpBreg22-134]
	_ = x[enum.DwarfOpBreg23-135]
	_ = x[enum.DwarfOpBreg24-136]
	_ = x[enum.DwarfOpBreg25-137]
	_ = x[enum.DwarfOpBreg26-138]
	_ = x[enum.DwarfOpBreg27-139]
	_ = x[enum.DwarfOpBreg28-140]
	_ = x[enum.DwarfOpBreg29-141]
	_ = x[enum.DwarfOpBreg30-142]
	_ = x[enum.DwarfOpBreg31-143]
	_ = x[enum.DwarfOpRegx-144]
	_ = x[enum.DwarfOpFbreg-145]
	_ = x[enum.DwarfOpBregx-146]
	_ = x[enum.DwarfOpPiece-147]
	_ = x[enum.DwarfOpDerefSize-148]
	_ = x[enum.DwarfOpXderefSize-149]
	_ = x[enum.DwarfOpNop-150]
	_ = x[enum.DwarfOpPushObjectAddress-151]
	_ = x[enum.DwarfOpCall2-152]
	_ = x[enum.DwarfOpCall4-153]
	_ = x[enum.DwarfOpCallRef-154]
	_ = x[enum.DwarfOpFormTLSAddress-155]
	_ = x[enum.DwarfOpCallFrameCFA-156]
	_ = x[enum.DwarfOpBitPiece-157]
	_ = x[enum.DwarfOpImplicitValue-158]
	_ = x[enum.DwarfOpStackValue-159]
	_ = x[enum.DwarfOpImplicitPointer-160]
	_ = x[enum.DwarfOpAddrx-161]
	_ = x[enum.DwarfOpConstx-162]
	_ = x[enum.DwarfOpEntryValue-163]
	_ = x[enum.DwarfOpConstType-164]
	_ = x[enum.DwarfOpRegvalType-165]
	_ = x[enum.DwarfOpDerefType-166]
	_ = x[enum.DwarfOpXderefType-167]
	_ = x[enum.DwarfOpConvert-168]
	_ = x[enum.DwarfOpReinterpret-169]
	_ = x[enum.DwarfOpGNUPushTLSAddress-224]
	_ = x[enum.DwarfOpGNUAddrIndex-251]
	_ = x[enum.DwarfOpGNUConstIndex-252]
	_ = x[enum.DwarfOpLLVMFragment-4096]
}

const (
	_DwarfOp_name_0 = "DW_OP_addr"
	_DwarfOp_name_1 = "DW_OP_deref"
	_DwarfOp_name_2 = "DW_OP_const1uDW_OP_const1sDW_OP_const2uDW_OP_const2sDW_OP_const4uDW_OP_const4sDW_OP_const8uDW_OP_const8sDW_OP_constuDW_OP_constsDW_OP_dupDW_OP_dropDW_OP_overDW_OP_pickDW_OP_swapDW_OP_rotDW_OP_xderefDW_OP_absDW_OP_andDW_OP_divDW_OP_minusDW_OP_modDW_OP_mulDW_OP_negDW_OP_notDW_OP_orDW_OP_plusDW_OP_plus_uconstDW_OP_shlDW_OP_shrDW_OP_shraDW_OP_xorDW_OP_braDW_OP_eqDW_OP_geDW_OP_gtDW_OP_leDW_OP_ltDW_OP_neDW_OP_skipDW_OP_lit0DW_OP_lit1DW_OP_lit2DW_OP_lit3DW_OP_lit4DW_OP_lit5DW_OP_lit6DW_OP_lit7DW_OP_lit8DW_OP_lit9DW_OP_lit10DW_OP_lit11DW_OP_lit12DW_OP_lit13DW_OP_lit14DW_OP_lit15DW_OP_lit16DW_OP_lit17DW_OP_lit18DW_OP_lit19DW_OP_lit20DW_OP_lit21DW_OP_lit22DW_OP_lit23DW_OP_lit24DW_OP_lit25DW_OP_lit26DW_OP_lit27DW_OP_lit28DW_OP_lit29DW_OP_lit30DW_OP_lit31DW_OP_reg0DW_OP_reg1DW_OP_reg2DW_OP_reg3DW_OP_reg4DW_OP_reg5DW_OP_reg6DW_OP_reg7DW_OP_reg8DW_OP_reg9DW_OP_reg10DW_OP_reg11DW_OP_reg12DW_OP_reg13DW_OP_reg14DW_OP_reg15DW_OP_reg16DW_OP_reg17DW_OP_reg18DW_OP_reg19DW_OP_reg20DW_OP_reg21DW_OP_reg22DW_OP_reg23DW_OP_reg24DW_OP_reg25DW_OP_reg26DW_OP_reg27DW_OP_reg28DW_OP_reg29DW_OP_reg30DW_OP_reg31DW_OP_breg0DW_OP_breg1DW_OP_breg2DW_OP_breg3DW_OP_breg4DW_OP_breg5DW_OP_breg6DW_OP_breg7DW_OP_breg8DW_OP_breg9DW_OP_breg10DW_OP_breg11DW_OP_breg12DW_OP_breg13DW_OP_breg14DW_OP_breg15DW_OP_breg16DW_OP_breg17DW_OP_breg18DW_OP_breg19DW_OP_breg20DW_OP_breg21DW_OP_breg22DW_OP_breg23DW_OP_breg24DW_OP_breg25DW_OP_breg26DW_OP_breg27DW_OP_breg28DW_OP_breg29DW_OP_breg30DW_OP_breg31DW_OP_regxDW_OP_fbregDW_OP_bregxDW_OP_pieceDW_OP_deref_sizeDW_OP_xderef_sizeDW_OP_nopDW_OP_push_object_addressDW_OP_call2DW_OP_call4DW_OP_call_refDW_OP_form_tls_addressDW_OP_call_frame_cfaDW_OP_bit_pieceDW_OP_implicit_valueDW_OP_stack_valueDW_OP_implicit_pointerDW_OP_addrxDW_OP_constxDW_OP_entry_valueDW_OP_const_typeDW_OP_regval_typeDW_OP_deref_typeDW_OP_xderef_typeDW_OP_convertDW_OP_reinterpret"
	_DwarfOp_name_3 = "DW_OP_GNU_push_tls_address"
	_DwarfOp_name_4 = "DW_OP_GNU_addr_indexDW_OP_GNU_const_index"
	_DwarfOp_name_5 = "DW_OP_LLVM_fragment"
)

var (
	_DwarfOp_index_2 = [...]uint16{0, 13, 26, 39, 52, 65, 78, 91, 104, 116, 128, 137, 147, 157, 167, 177, 186, 198, 207, 216, 225, 236, 245, 254, 263, 272, 280, 290, 307, 316, 325, 335, 344, 353, 361, 369, 377, 385, 393, 401, 411, 421, 431, 441, 451, 461, 471, 481, 491, 501, 511, 522, 533, 544, 555, 566, 577, 588, 599, 610, 621, 632, 643, 654, 665, 676, 687, 698, 709, 720, 731, 742, 753, 763, 773, 783, 793, 803, 813, 823, 833, 843, 853, 864, 875, 886, 897, 908, 919, 930, 941, 952, 963, 974, 985, 996, 1007, 1018, 1029, 1040, 1051, 1062, 1073, 1084, 1095, 1106, 1117, 1128, 1139, 1150, 1161, 1172, 1183, 1194, 1205, 1217, 1229, 1241, 1253, 1265, 1277, 1289, 1301, 1313, 1325, 1337, 1349, 1361, 1373, 1385, 1397, 1409, 1421, 1433, 1445, 1457, 1469, 1479, 1490, 1501, 1512, 1528, 1545, 1554, 1579, 1590, 1601, 1615, 1637, 1657, 1672, 1692, 1709, 1731, 1742, 1754, 1771, 1787, 1804, 1820, 1837, 1850, 1867}
	_DwarfOp_index_4 = [...]uint8{0, 20, 41}
)

func DwarfOpFromString(s string) enum.DwarfOp {
	if len(s) == 0 {
		return 0
	}
	if s == _DwarfOp_name_0 {
		return enum.DwarfOp(3)
	}
	if s == _DwarfOp_name_1 {
		return enum.DwarfOp(6)
	}
	for i := range _DwarfOp_index_2[:len(_DwarfOp_index_2)-1] {
		if s == _DwarfOp_name_2[_DwarfOp_index_2[i]:_DwarfOp_index_2[i+1]] {
			return enum.DwarfOp(i + 8)
		}
	}
	if s == _DwarfOp_name_3 {
		return enum.DwarfOp(224)
	}
	for i := range _DwarfOp_index_4[:len(_DwarfOp_index_4)-1] {
		if s == _DwarfOp_name_4[_DwarfOp_index_4[i]:_DwarfOp_index_4[i+1]] {
			return enum.DwarfOp(i + 251)
		}
	}
	if s == _DwarfOp_name_5 {
		return enum.DwarfOp(4096)
	}
	panic(fmt.Errorf("unable to locate DwarfOp enum corresponding to %q", s))
}
