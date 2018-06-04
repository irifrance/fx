// Package fx provides fixed point numbers based on int64s.
//
// The fixed point numbers support multiplication, division, float64
// conversion, Sin, Cos, Tan, Atan, const Pi, Sqrt, const Sqrt2
//
// Build flags may be used to define different q's (number of fractional bits),
// and a generator program generates the necessary constants for more options of
// q values.
//
// Using Fixed Point Numbers
//
// Fixed point numbers offer the following avantages
//
//  - uniform precision
//  - replicability across platforms
//  - built-in addition, subtraction, shifting using machine integer ops.
//
// with the following costs w.r.t. floats
//
//  - reduced dynamic range
//  - slower multiplication (5 native multiplies + shifts and adds)
//  - even slower division
//  - slower trigonometric functions (based on Cordic rotation currently)
//  - less general math support
//
// Package fx was developed primarily in support of an application
// which needed replicability accross platforms (and programming languages)
// of some trigonometric functions and which targeted hardware implementations.
package fx
