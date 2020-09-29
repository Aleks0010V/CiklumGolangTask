package modules

import (
	"encoding/json"
	"testing"
)

func TestMergeArticlesWithMarketing(t *testing.T) {
	var response ResponseByList
	articles := make([]Article, 18)
	marketingContent := make([]ContentMarketing, 2)

	testArticle := Article{
		Type:          "Article",
		HarvesterId:   "dagbladet.no/72043690",
		Cerebro_score: 3.198918510034566,
		URL:           "https://www.dagbladet.no/kjendis/nedslaende-meghan-tall/72043690",
		Title:         "Nedslående Meghan-tall",
		CleanImage:    "https://dbstatic.no/?imageId=72043713&panoy=32.795698924731&panox=0&panow=100&panoh=52.688172043011&heighty=0&heightw=41.016949152542&heighth=100&heightx=22.71186440678",
	}

	testMC := ContentMarketing{
		Type:              "ContentMarketing",
		HarvesterId:       "norsk-tipping.no/180120",
		CommercialPartner: "Norsk",
		LogoURL:           "https://www.dagbladet.no/files/2018/11/20/norsk%20tipping%20logo%202.png",
		Cerebro_score:     0.1,
		URL:               "https://www.norsk-tipping.no/artikler/lotto180120?WT.mc_id=Dagbladet_dagbladet_ekomm_lotto_vinnerhistorielotto_DB_ekomm&utm_source=dagbladet&utm_medium=ekomm&utm_content=lotto_vinnerhistorielotto_DB&utm_campaign=ekomm",
		Title:             "Disse tre vant 5 millioner hver!",
		CleanImage:        "https://dbstatic.no/72043321.jpg?imageId=72043321&x=0.000000&y=2.444988&cropw=100.000000&croph=80.440098",
	}

	for i := 0; i < 18; i++ {
		articles[i] = testArticle
	}
	for i := 0; i < 2; i++ {
		marketingContent[i] = testMC
	}

	response.MergeArticlesWithMarketing(articles, marketingContent, 6)
	for i := 0; i < 5; i++ {
		if response.Items[i] != testArticle {
			t.Error("Pattern is broken in articles part")
		}
	}
	if response.Items[5] != testMC {
		t.Error("Pattern is broken in MC part")
	}
	if response.Items[len(response.Items)-1] != EAd {
		t.Error("Pattern is broken in empty ad part")
	}
	_, err := json.Marshal(response.Items)
	if err != nil {
		t.Fatalf("Error while JSON encoding: %v", err)
	} else {
		t.Log("JSON is valid")
	}
}
