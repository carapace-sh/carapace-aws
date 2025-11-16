package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_updateLftagCmd = &cobra.Command{
	Use:   "update-lftag",
	Short: "Updates the list of possible values for the specified LF-tag key.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_updateLftagCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lakeformation_updateLftagCmd).Standalone()

		lakeformation_updateLftagCmd.Flags().String("catalog-id", "", "The identifier for the Data Catalog.")
		lakeformation_updateLftagCmd.Flags().String("tag-key", "", "The key-name for the LF-tag for which to add or delete values.")
		lakeformation_updateLftagCmd.Flags().String("tag-values-to-add", "", "A list of LF-tag values to add from the LF-tag.")
		lakeformation_updateLftagCmd.Flags().String("tag-values-to-delete", "", "A list of LF-tag values to delete from the LF-tag.")
		lakeformation_updateLftagCmd.MarkFlagRequired("tag-key")
	})
	lakeformationCmd.AddCommand(lakeformation_updateLftagCmd)
}
