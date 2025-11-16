package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoSync_unsubscribeFromDatasetCmd = &cobra.Command{
	Use:   "unsubscribe-from-dataset",
	Short: "Unsubscribes from receiving notifications when a dataset is modified by another device.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoSync_unsubscribeFromDatasetCmd).Standalone()

	cognitoSync_unsubscribeFromDatasetCmd.Flags().String("dataset-name", "", "The name of the dataset from which to unsubcribe.")
	cognitoSync_unsubscribeFromDatasetCmd.Flags().String("device-id", "", "The unique ID generated for this device by Cognito.")
	cognitoSync_unsubscribeFromDatasetCmd.Flags().String("identity-id", "", "Unique ID for this identity.")
	cognitoSync_unsubscribeFromDatasetCmd.Flags().String("identity-pool-id", "", "A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE) created by Amazon Cognito.")
	cognitoSync_unsubscribeFromDatasetCmd.MarkFlagRequired("dataset-name")
	cognitoSync_unsubscribeFromDatasetCmd.MarkFlagRequired("device-id")
	cognitoSync_unsubscribeFromDatasetCmd.MarkFlagRequired("identity-id")
	cognitoSync_unsubscribeFromDatasetCmd.MarkFlagRequired("identity-pool-id")
	cognitoSyncCmd.AddCommand(cognitoSync_unsubscribeFromDatasetCmd)
}
