package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var outposts_listCatalogItemsCmd = &cobra.Command{
	Use:   "list-catalog-items",
	Short: "Lists the items in the catalog.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(outposts_listCatalogItemsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(outposts_listCatalogItemsCmd).Standalone()

		outposts_listCatalogItemsCmd.Flags().String("ec2-family-filter", "", "Filters the results by EC2 family (for example, M5).")
		outposts_listCatalogItemsCmd.Flags().String("item-class-filter", "", "Filters the results by item class.")
		outposts_listCatalogItemsCmd.Flags().String("max-results", "", "")
		outposts_listCatalogItemsCmd.Flags().String("next-token", "", "")
		outposts_listCatalogItemsCmd.Flags().String("supported-storage-filter", "", "Filters the results by storage option.")
	})
	outpostsCmd.AddCommand(outposts_listCatalogItemsCmd)
}
