package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_updateMatchmakingConfigurationCmd = &cobra.Command{
	Use:   "update-matchmaking-configuration",
	Short: "**This API works with the following fleet types:** EC2, Anywhere, Container\n\nUpdates settings for a FlexMatch matchmaking configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_updateMatchmakingConfigurationCmd).Standalone()

	gamelift_updateMatchmakingConfigurationCmd.Flags().String("acceptance-required", "", "A flag that indicates whether a match that was created with this configuration must be accepted by the matched players.")
	gamelift_updateMatchmakingConfigurationCmd.Flags().String("acceptance-timeout-seconds", "", "The length of time (in seconds) to wait for players to accept a proposed match, if acceptance is required.")
	gamelift_updateMatchmakingConfigurationCmd.Flags().String("additional-player-count", "", "The number of player slots in a match to keep open for future players.")
	gamelift_updateMatchmakingConfigurationCmd.Flags().String("backfill-mode", "", "The method that is used to backfill game sessions created with this matchmaking configuration.")
	gamelift_updateMatchmakingConfigurationCmd.Flags().String("custom-event-data", "", "Information to add to all events related to the matchmaking configuration.")
	gamelift_updateMatchmakingConfigurationCmd.Flags().String("description", "", "A description for the matchmaking configuration.")
	gamelift_updateMatchmakingConfigurationCmd.Flags().String("flex-match-mode", "", "Indicates whether this matchmaking configuration is being used with Amazon GameLift Servers hosting or as a standalone matchmaking solution.")
	gamelift_updateMatchmakingConfigurationCmd.Flags().String("game-properties", "", "A set of key-value pairs that can store custom data in a game session.")
	gamelift_updateMatchmakingConfigurationCmd.Flags().String("game-session-data", "", "A set of custom game session properties, formatted as a single string value.")
	gamelift_updateMatchmakingConfigurationCmd.Flags().String("game-session-queue-arns", "", "The Amazon Resource Name ([ARN](https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-arn-format.html)) that is assigned to a Amazon GameLift Servers game session queue resource and uniquely identifies it.")
	gamelift_updateMatchmakingConfigurationCmd.Flags().String("name", "", "A unique identifier for the matchmaking configuration to update.")
	gamelift_updateMatchmakingConfigurationCmd.Flags().String("notification-target", "", "An SNS topic ARN that is set up to receive matchmaking notifications.")
	gamelift_updateMatchmakingConfigurationCmd.Flags().String("request-timeout-seconds", "", "The maximum duration, in seconds, that a matchmaking ticket can remain in process before timing out.")
	gamelift_updateMatchmakingConfigurationCmd.Flags().String("rule-set-name", "", "A unique identifier for the matchmaking rule set to use with this configuration.")
	gamelift_updateMatchmakingConfigurationCmd.MarkFlagRequired("name")
	gameliftCmd.AddCommand(gamelift_updateMatchmakingConfigurationCmd)
}
