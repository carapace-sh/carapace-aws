package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_createLftagCmd = &cobra.Command{
	Use:   "create-lftag",
	Short: "Creates an LF-tag with the specified name and values.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_createLftagCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lakeformation_createLftagCmd).Standalone()

		lakeformation_createLftagCmd.Flags().String("catalog-id", "", "The identifier for the Data Catalog.")
		lakeformation_createLftagCmd.Flags().String("tag-key", "", "The key-name for the LF-tag.")
		lakeformation_createLftagCmd.Flags().String("tag-values", "", "A list of possible values an attribute can take.")
		lakeformation_createLftagCmd.MarkFlagRequired("tag-key")
		lakeformation_createLftagCmd.MarkFlagRequired("tag-values")
	})
	lakeformationCmd.AddCommand(lakeformation_createLftagCmd)
}
