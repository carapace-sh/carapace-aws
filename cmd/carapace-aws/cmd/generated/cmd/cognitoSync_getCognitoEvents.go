package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoSync_getCognitoEventsCmd = &cobra.Command{
	Use:   "get-cognito-events",
	Short: "Gets the events and the corresponding Lambda functions associated with an identity pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoSync_getCognitoEventsCmd).Standalone()

	cognitoSync_getCognitoEventsCmd.Flags().String("identity-pool-id", "", "The Cognito Identity Pool ID for the request")
	cognitoSync_getCognitoEventsCmd.MarkFlagRequired("identity-pool-id")
	cognitoSyncCmd.AddCommand(cognitoSync_getCognitoEventsCmd)
}
