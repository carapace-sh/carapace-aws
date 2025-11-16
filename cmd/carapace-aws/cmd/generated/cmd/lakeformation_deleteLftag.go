package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_deleteLftagCmd = &cobra.Command{
	Use:   "delete-lftag",
	Short: "Deletes the specified LF-tag given a key name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_deleteLftagCmd).Standalone()

	lakeformation_deleteLftagCmd.Flags().String("catalog-id", "", "The identifier for the Data Catalog.")
	lakeformation_deleteLftagCmd.Flags().String("tag-key", "", "The key-name for the LF-tag to delete.")
	lakeformation_deleteLftagCmd.MarkFlagRequired("tag-key")
	lakeformationCmd.AddCommand(lakeformation_deleteLftagCmd)
}
