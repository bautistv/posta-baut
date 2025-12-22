package lookup

// TeamsLookup defines methods for looking up team and channel names by their IDs.
type TeamsLookupClient interface {
	TeamNameFromID(teamID string) (string, error)
	ChannelNameFromID(teamID string, channelID string) (string, error)
}
