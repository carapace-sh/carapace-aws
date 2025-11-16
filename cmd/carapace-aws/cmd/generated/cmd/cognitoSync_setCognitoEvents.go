package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoSync_setCognitoEventsCmd = &cobra.Command{
	Use:   "set-cognito-events",
	Short: "Sets the AWS Lambda function for a given event type for an identity pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoSync_setCognitoEventsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoSync_setCognitoEventsCmd).Standalone()

		cognitoSync_setCognitoEventsCmd.Flags().String("events", "", "The events to configure")
		cognitoSync_setCognitoEventsCmd.Flags().String("identity-pool-id", "", "The Cognito Identity Pool to use when configuring Cognito Events")
		cognitoSync_setCognitoEventsCmd.MarkFlagRequired("events")
		cognitoSync_setCognitoEventsCmd.MarkFlagRequired("identity-pool-id")
	})
	cognitoSyncCmd.AddCommand(cognitoSync_setCognitoEventsCmd)
}
