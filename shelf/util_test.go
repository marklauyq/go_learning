package shelf

import (
    "testing"
)

func TestTitleToFile(t *testing.T) {
    title := " Hello World "

    titleToTest := titleToFile(title)

    if titleToTest != "hello_world"{
        t.Errorf("Expected 'hello_world' got %s" , titleToTest)
    }
}
