package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoSync_listRecordsCmd = &cobra.Command{
	Use:   "list-records",
	Short: "Gets paginated records, optionally changed after a particular sync count for a dataset and identity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoSync_listRecordsCmd).Standalone()

	cognitoSync_listRecordsCmd.Flags().String("dataset-name", "", "A string of up to 128 characters.")
	cognitoSync_listRecordsCmd.Flags().String("identity-id", "", "A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE) created by Amazon Cognito.")
	cognitoSync_listRecordsCmd.Flags().String("identity-pool-id", "", "A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE) created by Amazon Cognito.")
	cognitoSync_listRecordsCmd.Flags().String("last-sync-count", "", "The last server sync count for this record.")
	cognitoSync_listRecordsCmd.Flags().String("max-results", "", "The maximum number of results to be returned.")
	cognitoSync_listRecordsCmd.Flags().String("next-token", "", "A pagination token for obtaining the next page of results.")
	cognitoSync_listRecordsCmd.Flags().String("sync-session-token", "", "A token containing a session ID, identity ID, and expiration.")
	cognitoSync_listRecordsCmd.MarkFlagRequired("dataset-name")
	cognitoSync_listRecordsCmd.MarkFlagRequired("identity-id")
	cognitoSync_listRecordsCmd.MarkFlagRequired("identity-pool-id")
	cognitoSyncCmd.AddCommand(cognitoSync_listRecordsCmd)
}
