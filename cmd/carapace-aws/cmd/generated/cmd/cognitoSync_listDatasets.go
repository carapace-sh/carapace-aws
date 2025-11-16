package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoSync_listDatasetsCmd = &cobra.Command{
	Use:   "list-datasets",
	Short: "Lists datasets for an identity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoSync_listDatasetsCmd).Standalone()

	cognitoSync_listDatasetsCmd.Flags().String("identity-id", "", "A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE) created by Amazon Cognito.")
	cognitoSync_listDatasetsCmd.Flags().String("identity-pool-id", "", "A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE) created by Amazon Cognito.")
	cognitoSync_listDatasetsCmd.Flags().String("max-results", "", "The maximum number of results to be returned.")
	cognitoSync_listDatasetsCmd.Flags().String("next-token", "", "A pagination token for obtaining the next page of results.")
	cognitoSync_listDatasetsCmd.MarkFlagRequired("identity-id")
	cognitoSync_listDatasetsCmd.MarkFlagRequired("identity-pool-id")
	cognitoSyncCmd.AddCommand(cognitoSync_listDatasetsCmd)
}
