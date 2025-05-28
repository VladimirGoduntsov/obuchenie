package main

import (
    "fmt"
    "math"
)

func main() {
    // Целочисленные типы
    fmt.Println("int8   :", math.MinInt8, "to", math.MaxInt8)
    fmt.Println("int16  :", math.MinInt16, "to", math.MaxInt16)
    fmt.Println("int32  :", math.MinInt32, "to", math.MaxInt32)
    fmt.Println("int64  :", math.MinInt64, "to", math.MaxInt64)
    fmt.Println("uint8  :", 0, "to", math.MaxUint8)
    fmt.Println("uint16 :", 0, "to", math.MaxUint16)
    fmt.Println("uint32 :", 0, "to", math.MaxUint32)
    fmt.Println("uint64 :", 0, "to", uint64(math.MaxUint64))
// Числа с плавающей точкой
    fmt.Println("float32:", -math.MaxFloat32, "to", math.MaxFloat32)
    fmt.Println("float64:", -math.MaxFloat64, "to", math.MaxFloat64)
}

