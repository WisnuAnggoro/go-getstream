package getstream

import (
	stream "gopkg.in/GetStream/stream-go2.v3"
)

type service struct {
	getstreamClient *stream.Client
}

type Service interface {
	AddPostByUserSerial(userSerial, postContent string) (*stream.AddActivityResponse, error)
	GetPostByUserSerial(userSerial string) (*stream.FlatFeedResponse, error)
	GetPostDetailByUserSerial(userSerial string) (*stream.EnrichedFlatFeedResponse, error)
	DeletePostByPostID(userSerial, postID string) error
	GetTimelineByUserSerial(userSerial string) (*stream.FlatFeedResponse, error)
	GetDetailTimelineByUserSerial(userSerial string) (*stream.EnrichedFlatFeedResponse, error)
	Follow(ownUserSerial, targetUserSerial string) error
	Unfollow(ownUserSerial, targetUserSerial string) error
	AddLikeToPostID(likerUserSerial, postID string) (*stream.Reaction, error)
	RetrieveLikeDetailOnPostID(postID string, limit int) (*stream.FilterReactionResponse, error)
	RetrieveLikeDetailOnPostIDWithPagination(postID, nextLikeID string, limit int) (*stream.FilterReactionResponse, error)
	RemoveLikeByReactionID(reactionID string) error
}

func NewService(getstreamClient *stream.Client) Service {
	return &service{
		getstreamClient: getstreamClient,
	}
}

func (s *service) AddPostByUserSerial(userSerial, postContent string) (*stream.AddActivityResponse, error) {
	// Get user feed object
	userFlatFeed, err := s.getstreamClient.FlatFeed("user", userSerial)
	if err != nil {
		return nil, err
	}

	// Add post activity to the feed
	resp, err := userFlatFeed.AddActivity(stream.Activity{
		Actor:  userFlatFeed.ID(),
		Verb:   "post",
		Object: "1",
		Extra: map[string]interface{}{
			"post": postContent,
		},
	})

	return resp, err
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

func (s *service) GetPostDetailByUserSerial(userSerial string) (*stream.EnrichedFlatFeedResponse, error) {
	// Get user feed object
	userFlatFeed, err := s.getstreamClient.FlatFeed("user", userSerial)
	if err != nil {
		return nil, err
	}

	// Add enriched option
	opts := []stream.GetActivitiesOption{
		stream.WithEnrichReactionKindsFilter("like"),
		stream.WithEnrichReactionCounts(),
	}

	// Get `enriched post` activity
	return userFlatFeed.GetEnrichedActivities(opts...)
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

func (s *service) AddLikeToPostID(likerUserSerial, postID string) (*stream.Reaction, error) {
	// Create a new `like` reaction
	r := stream.AddReactionRequestObject{
		Kind:       "like",
		ActivityID: postID,
		UserID:     likerUserSerial,
	}

	// Add the reaction to stream
	return s.getstreamClient.Reactions().Add(r)
}

func (s *service) RetrieveLikeDetailOnPostID(postID string, limit int) (*stream.FilterReactionResponse, error) {
	// Retrieve detail likes activity on selected postID
	return s.getstreamClient.
		Reactions().Filter(stream.ByActivityID(postID).ByKind("like"), stream.WithLimit(limit))
}

func (s *service) RetrieveLikeDetailOnPostIDWithPagination(postID, nextLikeID string, limit int) (*stream.FilterReactionResponse, error) {
	// retrieve the next {limit} likes using the id_lt param
	filterAttribute := stream.ByActivityID(postID).ByKind("like")
	limitation := stream.WithLimit(limit)
	pagination := stream.WithIDLT(nextLikeID)

	return s.getstreamClient.Reactions().Filter(filterAttribute, limitation, pagination)
}

func (s *service) RemoveLikeByReactionID(reactionID string) error {
	// Delete reaction by `reactionID`
	return s.getstreamClient.Reactions().Delete(reactionID)
}
