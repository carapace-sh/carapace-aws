package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_listServiceActionsCmd = &cobra.Command{
	Use:   "list-service-actions",
	Short: "Lists all self-service actions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_listServiceActionsCmd).Standalone()

	servicecatalog_listServiceActionsCmd.Flags().String("accept-language", "", "The language code.")
	servicecatalog_listServiceActionsCmd.Flags().String("page-size", "", "The maximum number of items to return with this call.")
	servicecatalog_listServiceActionsCmd.Flags().String("page-token", "", "The page token for the next set of results.")
	servicecatalogCmd.AddCommand(servicecatalog_listServiceActionsCmd)
}
