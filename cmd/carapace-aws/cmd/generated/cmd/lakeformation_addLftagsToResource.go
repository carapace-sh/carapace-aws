package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_addLftagsToResourceCmd = &cobra.Command{
	Use:   "add-lftags-to-resource",
	Short: "Attaches one or more LF-tags to an existing resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_addLftagsToResourceCmd).Standalone()

	lakeformation_addLftagsToResourceCmd.Flags().String("catalog-id", "", "The identifier for the Data Catalog.")
	lakeformation_addLftagsToResourceCmd.Flags().String("lftags", "", "The LF-tags to attach to the resource.")
	lakeformation_addLftagsToResourceCmd.Flags().String("resource", "", "The database, table, or column resource to which to attach an LF-tag.")
	lakeformation_addLftagsToResourceCmd.MarkFlagRequired("lftags")
	lakeformation_addLftagsToResourceCmd.MarkFlagRequired("resource")
	lakeformationCmd.AddCommand(lakeformation_addLftagsToResourceCmd)
}
