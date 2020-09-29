package modules

func (result *ResponseByList) MergeArticlesWithMarketing(articles []Article, contentMarketing []ContentMarketing,
	contentMarketingPosition int) {
	// by spec we need map each 5 articles to 1 ad, so the number of articles must be no less then 5*len(contentMarketing)
	if len(articles) < len(contentMarketing)*contentMarketingPosition {
		return
	}

	result.Items = make([]interface{}, 0)
	if contentMarketingPosition < 2 { // ToDo - change alg and remove this shit
		return
	}

	cmCounter := 0
	for i := 0; i < len(articles); i++ { // "every N-th position" pattern
		if (i+1)%contentMarketingPosition == 0 && cmCounter < len(contentMarketing) {
			result.Items = append(result.Items, contentMarketing[cmCounter])
			cmCounter++
		} else if (i+1)%contentMarketingPosition == 0 && cmCounter >= len(contentMarketing) {
			result.Items = append(result.Items, EAd)
		} else {
			result.Items = append(result.Items, articles[i])
		}
	}
}
