package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoSync_bulkPublishCmd = &cobra.Command{
	Use:   "bulk-publish",
	Short: "Initiates a bulk publish of all existing datasets for an Identity Pool to the configured stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoSync_bulkPublishCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoSync_bulkPublishCmd).Standalone()

		cognitoSync_bulkPublishCmd.Flags().String("identity-pool-id", "", "A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE) created by Amazon Cognito.")
		cognitoSync_bulkPublishCmd.MarkFlagRequired("identity-pool-id")
	})
	cognitoSyncCmd.AddCommand(cognitoSync_bulkPublishCmd)
}
