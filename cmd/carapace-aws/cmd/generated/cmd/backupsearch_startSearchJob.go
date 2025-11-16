package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backupsearch_startSearchJobCmd = &cobra.Command{
	Use:   "start-search-job",
	Short: "This operation creates a search job which returns recovery points filtered by SearchScope and items filtered by ItemFilters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backupsearch_startSearchJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backupsearch_startSearchJobCmd).Standalone()

		backupsearch_startSearchJobCmd.Flags().String("client-token", "", "Include this parameter to allow multiple identical calls for idempotency.")
		backupsearch_startSearchJobCmd.Flags().String("encryption-key-arn", "", "The encryption key for the specified search job.")
		backupsearch_startSearchJobCmd.Flags().String("item-filters", "", "Item Filters represent all input item properties specified when the search was created.")
		backupsearch_startSearchJobCmd.Flags().String("name", "", "Include alphanumeric characters to create a name for this search job.")
		backupsearch_startSearchJobCmd.Flags().String("search-scope", "", "This object can contain BackupResourceTypes, BackupResourceArns, BackupResourceCreationTime, BackupResourceTags, and SourceResourceArns to filter the recovery points returned by the search job.")
		backupsearch_startSearchJobCmd.Flags().String("tags", "", "List of tags returned by the operation.")
		backupsearch_startSearchJobCmd.MarkFlagRequired("search-scope")
	})
	backupsearchCmd.AddCommand(backupsearch_startSearchJobCmd)
}
