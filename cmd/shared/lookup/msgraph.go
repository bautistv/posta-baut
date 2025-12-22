package lookup

import (
	"context"
	"log"

	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
)

type MSGraphLookup struct {
	// graph client
	client msgraphsdk.GraphServiceClient
}

// TeamNameFromID looks up the team name given its ID.
func (l *MSGraphLookup) TeamNameFromID(ctx context.Context, teamID string) (string, error) {
	teamName := ""
	team, err := l.client.Teams().ByTeamId(teamID).Get(ctx, nil)
	if err != nil {
		log.Printf("failed to get team for team ID %s: %v", teamID, err)
		return "", err
	} else {
		teamName = *team.GetDisplayName()
	}

	return teamName, nil
}

// ChannelNameFromID looks up the channel name given its team ID and channel ID.
func (l *MSGraphLookup) ChannelNameFromID(ctx context.Context, teamID string, channelID string) (string, error) {
	channelName := ""
	channel, err := l.client.
		Teams().
		ByTeamId(teamID).
		Channels().
		ByChannelId(channelID).
		Get(ctx, nil)
	if err != nil {
		log.Printf("failed to get channel for channel ID %s: %v", channelID, err)
		return "", err
	} else {
		channelName = *channel.GetDisplayName()
	}

	return channelName, nil
}
