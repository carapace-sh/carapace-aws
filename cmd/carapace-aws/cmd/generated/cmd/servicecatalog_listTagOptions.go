package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_listTagOptionsCmd = &cobra.Command{
	Use:   "list-tag-options",
	Short: "Lists the specified TagOptions or all TagOptions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_listTagOptionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalog_listTagOptionsCmd).Standalone()

		servicecatalog_listTagOptionsCmd.Flags().String("filters", "", "The search filters.")
		servicecatalog_listTagOptionsCmd.Flags().String("page-size", "", "The maximum number of items to return with this call.")
		servicecatalog_listTagOptionsCmd.Flags().String("page-token", "", "The page token for the next set of results.")
	})
	servicecatalogCmd.AddCommand(servicecatalog_listTagOptionsCmd)
}
