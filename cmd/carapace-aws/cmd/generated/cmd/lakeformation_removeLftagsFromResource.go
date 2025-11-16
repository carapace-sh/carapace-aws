package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_removeLftagsFromResourceCmd = &cobra.Command{
	Use:   "remove-lftags-from-resource",
	Short: "Removes an LF-tag from the resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_removeLftagsFromResourceCmd).Standalone()

	lakeformation_removeLftagsFromResourceCmd.Flags().String("catalog-id", "", "The identifier for the Data Catalog.")
	lakeformation_removeLftagsFromResourceCmd.Flags().String("lftags", "", "The LF-tags to be removed from the resource.")
	lakeformation_removeLftagsFromResourceCmd.Flags().String("resource", "", "The database, table, or column resource where you want to remove an LF-tag.")
	lakeformation_removeLftagsFromResourceCmd.MarkFlagRequired("lftags")
	lakeformation_removeLftagsFromResourceCmd.MarkFlagRequired("resource")
	lakeformationCmd.AddCommand(lakeformation_removeLftagsFromResourceCmd)
}
