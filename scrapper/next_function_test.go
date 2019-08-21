package scrapper

import "testing"

func Test_Wuxiaworld_Next_Button_Works(t *testing.T){
	c , _ := getChapter("https://www.wuxiaworld.com/novel/dragon-prince-yuan/yz-chapter-1")

	nextChapter := WuxiaWorldNextSelector(c)

	if nextChapter != "https://www.wuxiaworld.com/novel/dragon-prince-yuan/yz-chapter-2"{
		t.Error("Expected to retrieve second chapter but received : ", nextChapter)
	}
}

func Test_Wuxiaworld_Next_Button_Can_Detect_Previews(t *testing.T){

}

func Test_Boxnovel_next_button_can_be_found(t *testing.T) {

}

