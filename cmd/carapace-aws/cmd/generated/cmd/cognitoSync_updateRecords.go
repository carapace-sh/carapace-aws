package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoSync_updateRecordsCmd = &cobra.Command{
	Use:   "update-records",
	Short: "Posts updates to records and adds and deletes records for a dataset and user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoSync_updateRecordsCmd).Standalone()

	cognitoSync_updateRecordsCmd.Flags().String("client-context", "", "Intended to supply a device ID that will populate the lastModifiedBy field referenced in other methods.")
	cognitoSync_updateRecordsCmd.Flags().String("dataset-name", "", "A string of up to 128 characters.")
	cognitoSync_updateRecordsCmd.Flags().String("device-id", "", "The unique ID generated for this device by Cognito.")
	cognitoSync_updateRecordsCmd.Flags().String("identity-id", "", "A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE) created by Amazon Cognito.")
	cognitoSync_updateRecordsCmd.Flags().String("identity-pool-id", "", "A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE) created by Amazon Cognito.")
	cognitoSync_updateRecordsCmd.Flags().String("record-patches", "", "A list of patch operations.")
	cognitoSync_updateRecordsCmd.Flags().String("sync-session-token", "", "The SyncSessionToken returned by a previous call to ListRecords for this dataset and identity.")
	cognitoSync_updateRecordsCmd.MarkFlagRequired("dataset-name")
	cognitoSync_updateRecordsCmd.MarkFlagRequired("identity-id")
	cognitoSync_updateRecordsCmd.MarkFlagRequired("identity-pool-id")
	cognitoSync_updateRecordsCmd.MarkFlagRequired("sync-session-token")
	cognitoSyncCmd.AddCommand(cognitoSync_updateRecordsCmd)
}
