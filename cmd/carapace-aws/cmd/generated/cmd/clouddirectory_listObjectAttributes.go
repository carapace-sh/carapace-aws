package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_listObjectAttributesCmd = &cobra.Command{
	Use:   "list-object-attributes",
	Short: "Lists all attributes that are associated with an object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_listObjectAttributesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(clouddirectory_listObjectAttributesCmd).Standalone()

		clouddirectory_listObjectAttributesCmd.Flags().String("consistency-level", "", "Represents the manner and timing in which the successful write or update of an object is reflected in a subsequent read operation of that same object.")
		clouddirectory_listObjectAttributesCmd.Flags().String("directory-arn", "", "The Amazon Resource Name (ARN) that is associated with the [Directory]() where the object resides.")
		clouddirectory_listObjectAttributesCmd.Flags().String("facet-filter", "", "Used to filter the list of object attributes that are associated with a certain facet.")
		clouddirectory_listObjectAttributesCmd.Flags().String("max-results", "", "The maximum number of items to be retrieved in a single call.")
		clouddirectory_listObjectAttributesCmd.Flags().String("next-token", "", "The pagination token.")
		clouddirectory_listObjectAttributesCmd.Flags().String("object-reference", "", "The reference that identifies the object whose attributes will be listed.")
		clouddirectory_listObjectAttributesCmd.MarkFlagRequired("directory-arn")
		clouddirectory_listObjectAttributesCmd.MarkFlagRequired("object-reference")
	})
	clouddirectoryCmd.AddCommand(clouddirectory_listObjectAttributesCmd)
}
