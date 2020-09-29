package modules

func (result *ResponseByList) MergeArticlesWithMarketing(articles []Article, contentMarketing []ContentMarketing,
	contentMarketingPosition int) {

	result.Items = make([]interface{}, 0)

	// it doesn't make any sense to place marketing object on every 1st position,
	// 'cause we will have a list of Marketing objects followed by {"type": "Ad"} objects
	if contentMarketingPosition < 2 {
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

	// if there is a situation, that we have more Marketing objects that articles,
	// I will simply add all of them to the end
	for ; cmCounter < len(contentMarketing); cmCounter++ {
		result.Items = append(result.Items, contentMarketing[cmCounter])
	}
}
