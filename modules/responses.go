package modules

func (result *ResponseByList) MergeArticlesWithMarketing(articles []Article, contentMarketing []ContentMarketing,
	contentMarketingPosition int) {

	// it doesn't make any sense to place marketing object on every 1st position,
	// 'cause we will have a list of Marketing objects followed by {"type": "Ad"} objects
	if contentMarketingPosition < 2 {
		result.Items = make([]interface{}, 0)
		return
	}

	cmCounter := 0
	arCounter := 0
	totalLength := len(articles) + len(articles)/contentMarketingPosition + 1
	result.Items = make([]interface{}, totalLength)

	for i := 0; i < totalLength; i++ { // probably, not the best implementation of "every N-th position" pattern
		if (i+1)%contentMarketingPosition == 0 && cmCounter < len(contentMarketing) {
			result.Items[i] = contentMarketing[cmCounter]
			cmCounter++
		} else if (i+1)%contentMarketingPosition == 0 && cmCounter >= len(contentMarketing) {
			result.Items[i] = EAd
		} else if arCounter < len(articles) {
			result.Items[i] = articles[arCounter]
			arCounter++
		}
	}

	// if there is a situation, that we have more Marketing objects that articles,
	// I will simply add all of them to the end
	for ; cmCounter < len(contentMarketing); cmCounter++ {
		result.Items = append(result.Items, contentMarketing[cmCounter])
	}
}
