package main

import "testing"

// Unit-tests
// Case 1 - valid directory with several files.
func TestCase1(t *testing.T) {
    var input = "../testing/case1"
    var expected = "foo\nbar\n.invisible\nconfig\n"

    if err, output := print_dir(input, true); output != expected || err != nil {
         t.Error("Test Failed: {} input, {} expected, received: {} and error was {}", input, expected, output, err)
    }
}

// Case 2 - invalid directory
func TestCase2(t *testing.T) {
    var input = "never seen"
    var expected = ""

    if err, output := print_dir(input, true); output != expected || err == nil {
         t.Error("Test Failed: {} input, {} expected, received: {} and error was {}", input, expected, output, err)
    }
}

// Case 3 - empty string.
func TestCase3(t *testing.T) {
    var input = ""
    var expected = ""

    if err, output := print_dir(input, true); output != expected || err == nil {
         t.Error("Test Failed: {} input, {} expected, received: {} and error was {}", input, expected, output, err)
    }
}

// Case 4 - valid directory with one file.
func TestCase4(t *testing.T) {
    var input = "../testing/case4"
    var expected = "foo\n"

    if err, output := print_dir(input, true); output != expected || err != nil {
         t.Error("Test Failed: {} input, {} expected, received: {} and error was {}", input, expected, output, err)
    }
}

// Case 5 - valid directory with no files.
func TestCase5(t *testing.T) {
    var input = "../testing/case5"
    var expected = ""

    if err, output := print_dir(input, true); output != expected || err != nil {
         t.Error("Test Failed: {} input, {} expected, received: {} and error was {}", input, expected, output, err)
    }
}

// Benchmarks
// Benchmark 1 - run print_dir on a directory with several files
func BenchmarkPrintDirSeveralFiles(b *testing.B) {
    for n := 0; n < b.N; n++ {
        print_dir("../testing/case1", true)
    }
}

// Benchmark 2 - run print_dir on a directory with one file
func BenchmarkPrintDirOneFile(b *testing.B) {
    for n := 0; n < b.N; n++ {
        print_dir("../testing/case4", true)
    }
}
