package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoSync_describeDatasetCmd = &cobra.Command{
	Use:   "describe-dataset",
	Short: "Gets meta data about a dataset by identity and dataset name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoSync_describeDatasetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoSync_describeDatasetCmd).Standalone()

		cognitoSync_describeDatasetCmd.Flags().String("dataset-name", "", "A string of up to 128 characters.")
		cognitoSync_describeDatasetCmd.Flags().String("identity-id", "", "A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE) created by Amazon Cognito.")
		cognitoSync_describeDatasetCmd.Flags().String("identity-pool-id", "", "A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE) created by Amazon Cognito.")
		cognitoSync_describeDatasetCmd.MarkFlagRequired("dataset-name")
		cognitoSync_describeDatasetCmd.MarkFlagRequired("identity-id")
		cognitoSync_describeDatasetCmd.MarkFlagRequired("identity-pool-id")
	})
	cognitoSyncCmd.AddCommand(cognitoSync_describeDatasetCmd)
}
