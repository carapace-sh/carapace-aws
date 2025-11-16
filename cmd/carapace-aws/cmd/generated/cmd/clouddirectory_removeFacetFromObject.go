package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_removeFacetFromObjectCmd = &cobra.Command{
	Use:   "remove-facet-from-object",
	Short: "Removes the specified facet from the specified object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_removeFacetFromObjectCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(clouddirectory_removeFacetFromObjectCmd).Standalone()

		clouddirectory_removeFacetFromObjectCmd.Flags().String("directory-arn", "", "The ARN of the directory in which the object resides.")
		clouddirectory_removeFacetFromObjectCmd.Flags().String("object-reference", "", "A reference to the object to remove the facet from.")
		clouddirectory_removeFacetFromObjectCmd.Flags().String("schema-facet", "", "The facet to remove.")
		clouddirectory_removeFacetFromObjectCmd.MarkFlagRequired("directory-arn")
		clouddirectory_removeFacetFromObjectCmd.MarkFlagRequired("object-reference")
		clouddirectory_removeFacetFromObjectCmd.MarkFlagRequired("schema-facet")
	})
	clouddirectoryCmd.AddCommand(clouddirectory_removeFacetFromObjectCmd)
}
