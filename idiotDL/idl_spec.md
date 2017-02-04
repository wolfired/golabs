[TOC]

## Basic Type

**Integer**: `i8 i16 i32 i64` `u8 u16 u32 u64`

**Floating**: `f32 f64`

**Boolean**: `b`

**Bit**: `bi`

**Byte**: `by`

**Character**: `c`

**String**: `s`

**Pointer**: type`*`

**Reference**: type`&`

## Complex Type

**Array**: `[`length?`]`vtype

**Map**: `[`ktype`]`vtype

## Custom Type

**Function/Method**
```
pkg.name (
    in param_name param_type
    out return_name return_type
)
```

**Struct/Class/Interface**
```
pkg.Name {
    anonymous_custom_type
    field_name field_type
    method_name (in param_name param_type, out return_name return_type)
}
```
