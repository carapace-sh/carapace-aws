package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_listTestSetRecordsCmd = &cobra.Command{
	Use:   "list-test-set-records",
	Short: "The list of test set records.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_listTestSetRecordsCmd).Standalone()

	lexv2Models_listTestSetRecordsCmd.Flags().String("max-results", "", "The maximum number of test set records to return in each page.")
	lexv2Models_listTestSetRecordsCmd.Flags().String("next-token", "", "If the response from the ListTestSetRecords operation contains more results than specified in the maxResults parameter, a token is returned in the response.")
	lexv2Models_listTestSetRecordsCmd.Flags().String("test-set-id", "", "The identifier of the test set to list its test set records.")
	lexv2Models_listTestSetRecordsCmd.MarkFlagRequired("test-set-id")
	lexv2ModelsCmd.AddCommand(lexv2Models_listTestSetRecordsCmd)
}
