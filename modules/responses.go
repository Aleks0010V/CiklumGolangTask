package modules

func (result *ResponseByList) MergeArticlesWithMarketing(articles []Article, contentMarketing []ContentMarketing,
	contentMarketingPosition int) {
	result.Items = make([]interface{}, 0)

	// with 0 we will have a list of contentMarketing objects followed by empty ads
	if contentMarketingPosition < 1 {
		return
	}

	cmCounter := 0
	for i := 0; i < len(articles); i++ { // probably, not the best implementation of "every N-th position" pattern
		if i != 0 && i%contentMarketingPosition == 0 && cmCounter < len(contentMarketing) {
			result.Items = append(result.Items, contentMarketing[cmCounter])
			cmCounter++
		} else if i%contentMarketingPosition == 0 && cmCounter >= len(contentMarketing) {
			result.Items = append(result.Items, EAd)
		}
		result.Items = append(result.Items, articles[i])
	}

	// if there is a situation, that we have more Marketing objects that articles,
	// I will simply add all of them to the end
	for ; cmCounter < len(contentMarketing); cmCounter++ {
		result.Items = append(result.Items, contentMarketing[cmCounter])
	}
}
