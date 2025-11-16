package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoSync_subscribeToDatasetCmd = &cobra.Command{
	Use:   "subscribe-to-dataset",
	Short: "Subscribes to receive notifications when a dataset is modified by another device.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoSync_subscribeToDatasetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoSync_subscribeToDatasetCmd).Standalone()

		cognitoSync_subscribeToDatasetCmd.Flags().String("dataset-name", "", "The name of the dataset to subcribe to.")
		cognitoSync_subscribeToDatasetCmd.Flags().String("device-id", "", "The unique ID generated for this device by Cognito.")
		cognitoSync_subscribeToDatasetCmd.Flags().String("identity-id", "", "Unique ID for this identity.")
		cognitoSync_subscribeToDatasetCmd.Flags().String("identity-pool-id", "", "A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE) created by Amazon Cognito.")
		cognitoSync_subscribeToDatasetCmd.MarkFlagRequired("dataset-name")
		cognitoSync_subscribeToDatasetCmd.MarkFlagRequired("device-id")
		cognitoSync_subscribeToDatasetCmd.MarkFlagRequired("identity-id")
		cognitoSync_subscribeToDatasetCmd.MarkFlagRequired("identity-pool-id")
	})
	cognitoSyncCmd.AddCommand(cognitoSync_subscribeToDatasetCmd)
}
