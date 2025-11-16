package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_deleteFacetCmd = &cobra.Command{
	Use:   "delete-facet",
	Short: "Deletes a given [Facet]().",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_deleteFacetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(clouddirectory_deleteFacetCmd).Standalone()

		clouddirectory_deleteFacetCmd.Flags().String("name", "", "The name of the facet to delete.")
		clouddirectory_deleteFacetCmd.Flags().String("schema-arn", "", "The Amazon Resource Name (ARN) that is associated with the [Facet]().")
		clouddirectory_deleteFacetCmd.MarkFlagRequired("name")
		clouddirectory_deleteFacetCmd.MarkFlagRequired("schema-arn")
	})
	clouddirectoryCmd.AddCommand(clouddirectory_deleteFacetCmd)
}
