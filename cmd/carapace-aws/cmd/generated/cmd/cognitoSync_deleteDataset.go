package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoSync_deleteDatasetCmd = &cobra.Command{
	Use:   "delete-dataset",
	Short: "Deletes the specific dataset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoSync_deleteDatasetCmd).Standalone()

	cognitoSync_deleteDatasetCmd.Flags().String("dataset-name", "", "A string of up to 128 characters.")
	cognitoSync_deleteDatasetCmd.Flags().String("identity-id", "", "A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE) created by Amazon Cognito.")
	cognitoSync_deleteDatasetCmd.Flags().String("identity-pool-id", "", "A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE) created by Amazon Cognito.")
	cognitoSync_deleteDatasetCmd.MarkFlagRequired("dataset-name")
	cognitoSync_deleteDatasetCmd.MarkFlagRequired("identity-id")
	cognitoSync_deleteDatasetCmd.MarkFlagRequired("identity-pool-id")
	cognitoSyncCmd.AddCommand(cognitoSync_deleteDatasetCmd)
}
