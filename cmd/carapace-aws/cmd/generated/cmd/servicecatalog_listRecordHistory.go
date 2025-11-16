package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_listRecordHistoryCmd = &cobra.Command{
	Use:   "list-record-history",
	Short: "Lists the specified requests or all performed requests.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_listRecordHistoryCmd).Standalone()

	servicecatalog_listRecordHistoryCmd.Flags().String("accept-language", "", "The language code.")
	servicecatalog_listRecordHistoryCmd.Flags().String("access-level-filter", "", "The access level to use to obtain results.")
	servicecatalog_listRecordHistoryCmd.Flags().String("page-size", "", "The maximum number of items to return with this call.")
	servicecatalog_listRecordHistoryCmd.Flags().String("page-token", "", "The page token for the next set of results.")
	servicecatalog_listRecordHistoryCmd.Flags().String("search-filter", "", "The search filter to scope the results.")
	servicecatalogCmd.AddCommand(servicecatalog_listRecordHistoryCmd)
}
