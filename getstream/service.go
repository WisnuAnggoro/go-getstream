package getstream

import (
	stream "gopkg.in/GetStream/stream-go2.v3"
)

type service struct {
	getstreamClient *stream.Client
}

type Service interface {
	AddPostByUserSerial(userSerial, postContent string) error
	GetPostByUserSerial(userSerial string) (*stream.FlatFeedResponse, error)
	DeletePostByPostID(userSerial, postID string) error
	GetTimelineByUserSerial(userSerial string) (*stream.FlatFeedResponse, error)
	GetDetailTimelineByUserSerial(userSerial string) (*stream.EnrichedFlatFeedResponse, error)
	Follow(ownUserSerial, targetUserSerial string) error
	Unfollow(ownUserSerial, targetUserSerial string) error
	AddLikeToPostID(likerUserSerial, postID string) error
	RetrieveLikeDetailOnPostID(postID string) (*stream.FilterReactionResponse, error)
	RemoveLikeByReactionID(reactionID string) error
}

func NewService(getstreamClient *stream.Client) Service {
	return &service{
		getstreamClient: getstreamClient,
	}
}

func (s *service) AddPostByUserSerial(userSerial, postContent string) error {
	// Get user feed object
	userFlatFeed, err := s.getstreamClient.FlatFeed("user", userSerial)
	if err != nil {
		return err
	}

	// Add post activity to the feed
	_, err = userFlatFeed.AddActivity(stream.Activity{
		Actor:  userFlatFeed.ID(),
		Verb:   "post",
		Object: "1",
		Extra: map[string]interface{}{
			"post": postContent,
		},
	})

	return err
}

func (s *service) GetPostByUserSerial(userSerial string) (*stream.FlatFeedResponse, error) {
	// Get user feed object
	userFlatFeed, err := s.getstreamClient.FlatFeed("user", userSerial)
	if err != nil {
		return nil, err
	}

	// Get `post` activity
	return userFlatFeed.GetActivities()
}

func (s *service) DeletePostByPostID(userSerial, postID string) error {
	// Get user feed object
	userFlatFeed, err := s.getstreamClient.FlatFeed("user", userSerial)
	if err != nil {
		return err
	}

	// Remove `post` activity specified by `activityID`
	return userFlatFeed.RemoveActivityByID(postID)
}

func (s *service) GetTimelineByUserSerial(userSerial string) (*stream.FlatFeedResponse, error) {
	// Get timeline feed object
	userFlatFeed, err := s.getstreamClient.FlatFeed("timeline", userSerial)
	if err != nil {
		return nil, err
	}

	// Get activities on `timeline` feed
	return userFlatFeed.GetActivities()
}

func (s *service) GetDetailTimelineByUserSerial(userSerial string) (*stream.EnrichedFlatFeedResponse, error) {
	// Get timeline feed object
	userFlatFeed, err := s.getstreamClient.FlatFeed("timeline", userSerial)
	if err != nil {
		return nil, err
	}

	// Add enriched option
	opts := []stream.GetActivitiesOption{
		stream.WithEnrichRecentReactions(),
		stream.WithEnrichReactionCounts(),
	}

	// Get `enriched` activities on `timeline` feed
	return userFlatFeed.GetEnrichedActivities(opts...)
}

func (s *service) Follow(ownUserSerial, targetUserSerial string) error {
	// Get timeline feed object
	ownUserFlatFeed, err := s.getstreamClient.FlatFeed("timeline", ownUserSerial)
	if err != nil {
		return err
	}

	// Get user feed object
	targetUserFlatFeed, err := s.getstreamClient.FlatFeed("user", targetUserSerial)
	if err != nil {
		return err
	}

	// `timeline` of `ownUser` will be filled by all activities of `targetUser`
	return ownUserFlatFeed.Follow(targetUserFlatFeed)
}

func (s *service) Unfollow(ownUserSerial, targetUserSerial string) error {
	// Get timeline feed object
	ownUserFlatFeed, err := s.getstreamClient.FlatFeed("timeline", ownUserSerial)
	if err != nil {
		return err
	}

	// Get user feed object
	targetUserFlatFeed, err := s.getstreamClient.FlatFeed("user", targetUserSerial)
	if err != nil {
		return err
	}

	// `timeline` of `ownUser` will no longer be filled by all activities of `targetUser`
	return ownUserFlatFeed.Unfollow(targetUserFlatFeed)
}

func (s *service) AddLikeToPostID(likerUserSerial, postID string) error {
	// Create a new `like` reaction
	r := stream.AddReactionRequestObject{
		Kind:       "like",
		ActivityID: postID,
		UserID:     likerUserSerial,
	}

	// Add the reaction to stream
	_, err := s.getstreamClient.Reactions().Add(r)
	return err
}

func (s *service) RetrieveLikeDetailOnPostID(postID string) (*stream.FilterReactionResponse, error) {
	// Retrieve detail likes activity on selected postID
	return s.getstreamClient.
		Reactions().Filter(stream.ByActivityID(postID).ByKind("like"))
}

func (s *service) RemoveLikeByReactionID(reactionID string) error {
	// Delete reaction by `reactionID`
	return s.getstreamClient.Reactions().Delete(reactionID)
}
