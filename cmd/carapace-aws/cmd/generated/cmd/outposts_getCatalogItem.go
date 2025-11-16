package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var outposts_getCatalogItemCmd = &cobra.Command{
	Use:   "get-catalog-item",
	Short: "Gets information about the specified catalog item.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(outposts_getCatalogItemCmd).Standalone()

	outposts_getCatalogItemCmd.Flags().String("catalog-item-id", "", "The ID of the catalog item.")
	outposts_getCatalogItemCmd.MarkFlagRequired("catalog-item-id")
	outpostsCmd.AddCommand(outposts_getCatalogItemCmd)
}
