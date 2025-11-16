package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_getLftagCmd = &cobra.Command{
	Use:   "get-lftag",
	Short: "Returns an LF-tag definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_getLftagCmd).Standalone()

	lakeformation_getLftagCmd.Flags().String("catalog-id", "", "The identifier for the Data Catalog.")
	lakeformation_getLftagCmd.Flags().String("tag-key", "", "The key-name for the LF-tag.")
	lakeformation_getLftagCmd.MarkFlagRequired("tag-key")
	lakeformationCmd.AddCommand(lakeformation_getLftagCmd)
}
