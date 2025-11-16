package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoSync_getBulkPublishDetailsCmd = &cobra.Command{
	Use:   "get-bulk-publish-details",
	Short: "Get the status of the last BulkPublish operation for an identity pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoSync_getBulkPublishDetailsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoSync_getBulkPublishDetailsCmd).Standalone()

		cognitoSync_getBulkPublishDetailsCmd.Flags().String("identity-pool-id", "", "A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE) created by Amazon Cognito.")
		cognitoSync_getBulkPublishDetailsCmd.MarkFlagRequired("identity-pool-id")
	})
	cognitoSyncCmd.AddCommand(cognitoSync_getBulkPublishDetailsCmd)
}
