package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_listLaunchPathsCmd = &cobra.Command{
	Use:   "list-launch-paths",
	Short: "Lists the paths to the specified product.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_listLaunchPathsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalog_listLaunchPathsCmd).Standalone()

		servicecatalog_listLaunchPathsCmd.Flags().String("accept-language", "", "The language code.")
		servicecatalog_listLaunchPathsCmd.Flags().String("page-size", "", "The maximum number of items to return with this call.")
		servicecatalog_listLaunchPathsCmd.Flags().String("page-token", "", "The page token for the next set of results.")
		servicecatalog_listLaunchPathsCmd.Flags().String("product-id", "", "The product identifier.")
		servicecatalog_listLaunchPathsCmd.MarkFlagRequired("product-id")
	})
	servicecatalogCmd.AddCommand(servicecatalog_listLaunchPathsCmd)
}
