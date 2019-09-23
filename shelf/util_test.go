package shelf

import (
    "testing"
)

func TestTitleToFile(t *testing.T) {
    title := " Hello World "

    titleToTest := titleToFile(title)

    if titleToTest != "hello_world.json"{
        t.Errorf("Expected 'hello_world.json' got %s" , titleToTest)
    }
}

func TestTitleToFileWithExtension(t *testing.T) {
    title := " Hello World "

    titleToTest := titleToFileWithExtensions(title, ".txt")

    if titleToTest != "hello_world.txt"{
        t.Errorf("Expected 'hello_world.txt' got %s" , titleToTest)
    }
}
