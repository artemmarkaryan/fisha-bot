package reaction

import "github.com/artemmarkaryan/fisha-facade/pkg/pb/gen/api"

type Reaction struct {
	api.ReactRequest_Reaction
	Code  string
	Emoji string
}

const LikeCode = "like"
const DislikeCode = "dislike"

var (
	LikeReaction = Reaction{
		Emoji:                 "👍",
		Code:                  LikeCode,
		ReactRequest_Reaction: api.ReactRequest_LIKE,
	}
	DislikeReaction = Reaction{
		Emoji:                 "👎",
		Code:                  DislikeCode,
		ReactRequest_Reaction: api.ReactRequest_DISLIKE,
	}
)

var Mapping = map[string]Reaction{
	LikeCode:    LikeReaction,
	DislikeCode: DislikeReaction,
}
