package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getCatalogsCmd = &cobra.Command{
	Use:   "get-catalogs",
	Short: "Retrieves all catalogs defined in a catalog in the Glue Data Catalog.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getCatalogsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_getCatalogsCmd).Standalone()

		glue_getCatalogsCmd.Flags().String("include-root", "", "Whether to list the default catalog in the account and region in the response.")
		glue_getCatalogsCmd.Flags().String("max-results", "", "The maximum number of catalogs to return in one response.")
		glue_getCatalogsCmd.Flags().String("next-token", "", "A continuation token, if this is a continuation call.")
		glue_getCatalogsCmd.Flags().Bool("no-recursive", false, "Whether to list all catalogs across the catalog hierarchy, starting from the `ParentCatalogId`.")
		glue_getCatalogsCmd.Flags().String("parent-catalog-id", "", "The ID of the parent catalog in which the catalog resides.")
		glue_getCatalogsCmd.Flags().Bool("recursive", false, "Whether to list all catalogs across the catalog hierarchy, starting from the `ParentCatalogId`.")
		glue_getCatalogsCmd.Flag("no-recursive").Hidden = true
	})
	glueCmd.AddCommand(glue_getCatalogsCmd)
}
