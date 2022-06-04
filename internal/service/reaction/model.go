package reaction

import "github.com/artemmarkaryan/fisha-facade/pkg/pb/gen/api"

type Reaction struct {
	api.ReactRequest_Reaction
	Code  string
	Emoji string
}

var (
	Like = Reaction{
		Emoji:                 "👍",
		Code:                  "like",
		ReactRequest_Reaction: api.ReactRequest_LIKE,
	}
	Dislike = Reaction{
		Emoji:                 "👎",
		Code:                  "dislike",
		ReactRequest_Reaction: api.ReactRequest_DISLIKE,
	}
)
