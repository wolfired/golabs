package wabf

type opcode byte

const (
	//Control flow operators
	opcode_unreachable opcode = 0x00
	opcode_nop         opcode = 0x01
	opcode_block       opcode = 0x02
	opcode_loop        opcode = 0x03
	opcode_if          opcode = 0x04
	opcode_else        opcode = 0x05
	opcode_end         opcode = 0x0b
	opcode_br          opcode = 0x0c
	opcode_br_if       opcode = 0x0d
	opcode_br_table    opcode = 0x0e
	opcode_return      opcode = 0x0f
	//Call operators
	opcode_call          opcode = 0x10
	opcode_call_indirect opcode = 0x11
	//Parametric operators
	opcode_drop   opcode = 0x1a
	opcode_select opcode = 0x1b
	//Variable access
	opcode_get_local  opcode = 0x20
	opcode_set_local  opcode = 0x21
	opcode_tee_local  opcode = 0x22
	opcode_get_global opcode = 0x23
	opcode_set_global opcode = 0x24
	//Memory-related operators
	opcode_i32_load       opcode = 0x28
	opcode_i64_load       opcode = 0x29
	opcode_f32_load       opcode = 0x2a
	opcode_f64_load       opcode = 0x2b
	opcode_i32_load8_s    opcode = 0x2c
	opcode_i32_load8_u    opcode = 0x2d
	opcode_i32_load16_s   opcode = 0x2e
	opcode_i32_load16_u   opcode = 0x2f
	opcode_i64_load8_s    opcode = 0x30
	opcode_i64_load8_u    opcode = 0x31
	opcode_i64_load16_s   opcode = 0x32
	opcode_i64_load16_u   opcode = 0x33
	opcode_i64_load32_s   opcode = 0x34
	opcode_i64_load32_u   opcode = 0x35
	opcode_i32_store      opcode = 0x36
	opcode_i64_store      opcode = 0x37
	opcode_f32_store      opcode = 0x38
	opcode_f64_store      opcode = 0x39
	opcode_i32_store8     opcode = 0x3a
	opcode_i32_store16    opcode = 0x3b
	opcode_i64_store8     opcode = 0x3c
	opcode_i64_store16    opcode = 0x3d
	opcode_i64_store32    opcode = 0x3e
	opcode_current_memory opcode = 0x3f
	opcode_grow_memory    opcode = 0x40
	//Constants
	opcode_i32_const opcode = 0x41
	opcode_i64_const opcode = 0x42
	opcode_f32_const opcode = 0x43
	opcode_f64_const opcode = 0x44
	//Comparison operators
	opcode_i32_eqz  opcode = 0x45
	opcode_i32_eq   opcode = 0x46
	opcode_i32_ne   opcode = 0x47
	opcode_i32_lt_s opcode = 0x48
	opcode_i32_lt_u opcode = 0x49
	opcode_i32_gt_s opcode = 0x4a
	opcode_i32_gt_u opcode = 0x4b
	opcode_i32_le_s opcode = 0x4c
	opcode_i32_le_u opcode = 0x4d
	opcode_i32_ge_s opcode = 0x4e
	opcode_i32_ge_u opcode = 0x4f
	opcode_i64_eqz  opcode = 0x50
	opcode_i64_eq   opcode = 0x51
	opcode_i64_ne   opcode = 0x52
	opcode_i64_lt_s opcode = 0x53
	opcode_i64_lt_u opcode = 0x54
	opcode_i64_gt_s opcode = 0x55
	opcode_i64_gt_u opcode = 0x56
	opcode_i64_le_s opcode = 0x57
	opcode_i64_le_u opcode = 0x58
	opcode_i64_ge_s opcode = 0x59
	opcode_i64_ge_u opcode = 0x5a
	opcode_f32_eq   opcode = 0x5b
	opcode_f32_ne   opcode = 0x5c
	opcode_f32_lt   opcode = 0x5d
	opcode_f32_gt   opcode = 0x5e
	opcode_f32_le   opcode = 0x5f
	opcode_f32_ge   opcode = 0x60
	opcode_f64_eq   opcode = 0x61
	opcode_f64_ne   opcode = 0x62
	opcode_f64_lt   opcode = 0x63
	opcode_f64_gt   opcode = 0x64
	opcode_f64_le   opcode = 0x65
	opcode_f64_ge   opcode = 0x66
	//Numeric operators
	opcode_i32_clz      opcode = 0x67
	opcode_i32_ctz      opcode = 0x68
	opcode_i32_popcnt   opcode = 0x69
	opcode_i32_add      opcode = 0x6a
	opcode_i32_sub      opcode = 0x6b
	opcode_i32_mul      opcode = 0x6c
	opcode_i32_div_s    opcode = 0x6d
	opcode_i32_div_u    opcode = 0x6e
	opcode_i32_rem_s    opcode = 0x6f
	opcode_i32_rem_u    opcode = 0x70
	opcode_i32_and      opcode = 0x71
	opcode_i32_or       opcode = 0x72
	opcode_i32_xor      opcode = 0x73
	opcode_i32_shl      opcode = 0x74
	opcode_i32_shr_s    opcode = 0x75
	opcode_i32_shr_u    opcode = 0x76
	opcode_i32_rotl     opcode = 0x77
	opcode_i32_rotr     opcode = 0x78
	opcode_i64_clz      opcode = 0x79
	opcode_i64_ctz      opcode = 0x7a
	opcode_i64_popcnt   opcode = 0x7b
	opcode_i64_add      opcode = 0x7c
	opcode_i64_sub      opcode = 0x7d
	opcode_i64_mul      opcode = 0x7e
	opcode_i64_div_s    opcode = 0x7f
	opcode_i64_div_u    opcode = 0x80
	opcode_i64_rem_s    opcode = 0x81
	opcode_i64_rem_u    opcode = 0x82
	opcode_i64_and      opcode = 0x83
	opcode_i64_or       opcode = 0x84
	opcode_i64_xor      opcode = 0x85
	opcode_i64_shl      opcode = 0x86
	opcode_i64_shr_s    opcode = 0x87
	opcode_i64_shr_u    opcode = 0x88
	opcode_i64_rotl     opcode = 0x89
	opcode_i64_rotr     opcode = 0x8a
	opcode_f32_abs      opcode = 0x8b
	opcode_f32_neg      opcode = 0x8c
	opcode_f32_ceil     opcode = 0x8d
	opcode_f32_floor    opcode = 0x8e
	opcode_f32_trunc    opcode = 0x8f
	opcode_f32_nearest  opcode = 0x90
	opcode_f32_sqrt     opcode = 0x91
	opcode_f32_add      opcode = 0x92
	opcode_f32_sub      opcode = 0x93
	opcode_f32_mul      opcode = 0x94
	opcode_f32_div      opcode = 0x95
	opcode_f32_min      opcode = 0x96
	opcode_f32_max      opcode = 0x97
	opcode_f32_copysign opcode = 0x98
	opcode_f64_abs      opcode = 0x99
	opcode_f64_neg      opcode = 0x9a
	opcode_f64_ceil     opcode = 0x9b
	opcode_f64_floor    opcode = 0x9c
	opcode_f64_trunc    opcode = 0x9d
	opcode_f64_nearest  opcode = 0x9e
	opcode_f64_sqrt     opcode = 0x9f
	opcode_f64_add      opcode = 0xa0
	opcode_f64_sub      opcode = 0xa1
	opcode_f64_mul      opcode = 0xa2
	opcode_f64_div      opcode = 0xa3
	opcode_f64_min      opcode = 0xa4
	opcode_f64_max      opcode = 0xa5
	opcode_f64_copysign opcode = 0xa6
	//Conversions
	opcodei32_wrap_i64      opcode = 0xa7
	opcodei32_trunc_s_f32   opcode = 0xa8
	opcodei32_trunc_u_f32   opcode = 0xa9
	opcodei32_trunc_s_f64   opcode = 0xaa
	opcodei32_trunc_u_f64   opcode = 0xab
	opcodei64_extend_s_i32  opcode = 0xac
	opcodei64_extend_u_i32  opcode = 0xad
	opcodei64_trunc_s_f32   opcode = 0xae
	opcodei64_trunc_u_f32   opcode = 0xaf
	opcodei64_trunc_s_f64   opcode = 0xb0
	opcodei64_trunc_u_f64   opcode = 0xb1
	opcodef32_convert_s_i32 opcode = 0xb2
	opcodef32_convert_u_i32 opcode = 0xb3
	opcodef32_convert_s_i64 opcode = 0xb4
	opcodef32_convert_u_i64 opcode = 0xb5
	opcodef32_demote_f64    opcode = 0xb6
	opcodef64_convert_s_i32 opcode = 0xb7
	opcodef64_convert_u_i32 opcode = 0xb8
	opcodef64_convert_s_i64 opcode = 0xb9
	opcodef64_convert_u_i64 opcode = 0xba
	opcodef64_promote_f32   opcode = 0xbb
	//Reinterpretations
	opcodei32_reinterpret_f32 opcode = 0xbc
	opcodei64_reinterpret_f64 opcode = 0xbd
	opcodef32_reinterpret_i32 opcode = 0xbe
	opcodef64_reinterpret_i64 opcode = 0xbf
)

type operator struct {
	name string
	opc  opcode
	desc string
}
