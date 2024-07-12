package main

import (
    "fmt"
    "strings"
    "testing"
)

func BenchmarkSprintf(b *testing.B) {
    name := "World"
    for i := 0; i < b.N; i++ {
        _ = fmt.Sprintf("Hello, %s!", name)
    }
}

func BenchmarkBuilder(b *testing.B) {
    name := "World"
    for i := 0; i < b.N; i++ {
        var builder strings.Builder
        builder.WriteString("Hello, ")
        builder.WriteString(name)
        builder.WriteString("!")
        _ = builder.String()
    }
}