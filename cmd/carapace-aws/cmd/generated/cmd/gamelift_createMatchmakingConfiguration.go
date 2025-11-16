package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_createMatchmakingConfigurationCmd = &cobra.Command{
	Use:   "create-matchmaking-configuration",
	Short: "**This API works with the following fleet types:** EC2, Anywhere, Container\n\nDefines a new matchmaking configuration for use with FlexMatch.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_createMatchmakingConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_createMatchmakingConfigurationCmd).Standalone()

		gamelift_createMatchmakingConfigurationCmd.Flags().String("acceptance-required", "", "A flag that determines whether a match that was created with this configuration must be accepted by the matched players.")
		gamelift_createMatchmakingConfigurationCmd.Flags().String("acceptance-timeout-seconds", "", "The length of time (in seconds) to wait for players to accept a proposed match, if acceptance is required.")
		gamelift_createMatchmakingConfigurationCmd.Flags().String("additional-player-count", "", "The number of player slots in a match to keep open for future players.")
		gamelift_createMatchmakingConfigurationCmd.Flags().String("backfill-mode", "", "The method used to backfill game sessions that are created with this matchmaking configuration.")
		gamelift_createMatchmakingConfigurationCmd.Flags().String("custom-event-data", "", "Information to be added to all events related to this matchmaking configuration.")
		gamelift_createMatchmakingConfigurationCmd.Flags().String("description", "", "A human-readable description of the matchmaking configuration.")
		gamelift_createMatchmakingConfigurationCmd.Flags().String("flex-match-mode", "", "Indicates whether this matchmaking configuration is being used with Amazon GameLift Servers hosting or as a standalone matchmaking solution.")
		gamelift_createMatchmakingConfigurationCmd.Flags().String("game-properties", "", "A set of key-value pairs that can store custom data in a game session.")
		gamelift_createMatchmakingConfigurationCmd.Flags().String("game-session-data", "", "A set of custom game session properties, formatted as a single string value.")
		gamelift_createMatchmakingConfigurationCmd.Flags().String("game-session-queue-arns", "", "The Amazon Resource Name ([ARN](https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-arn-format.html)) that is assigned to a Amazon GameLift Servers game session queue resource and uniquely identifies it.")
		gamelift_createMatchmakingConfigurationCmd.Flags().String("name", "", "A unique identifier for the matchmaking configuration.")
		gamelift_createMatchmakingConfigurationCmd.Flags().String("notification-target", "", "An SNS topic ARN that is set up to receive matchmaking notifications.")
		gamelift_createMatchmakingConfigurationCmd.Flags().String("request-timeout-seconds", "", "The maximum duration, in seconds, that a matchmaking ticket can remain in process before timing out.")
		gamelift_createMatchmakingConfigurationCmd.Flags().String("rule-set-name", "", "A unique identifier for the matchmaking rule set to use with this configuration.")
		gamelift_createMatchmakingConfigurationCmd.Flags().String("tags", "", "A list of labels to assign to the new matchmaking configuration resource.")
		gamelift_createMatchmakingConfigurationCmd.MarkFlagRequired("acceptance-required")
		gamelift_createMatchmakingConfigurationCmd.MarkFlagRequired("name")
		gamelift_createMatchmakingConfigurationCmd.MarkFlagRequired("request-timeout-seconds")
		gamelift_createMatchmakingConfigurationCmd.MarkFlagRequired("rule-set-name")
	})
	gameliftCmd.AddCommand(gamelift_createMatchmakingConfigurationCmd)
}
