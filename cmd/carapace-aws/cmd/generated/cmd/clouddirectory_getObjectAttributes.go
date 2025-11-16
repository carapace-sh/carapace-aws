package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_getObjectAttributesCmd = &cobra.Command{
	Use:   "get-object-attributes",
	Short: "Retrieves attributes within a facet that are associated with an object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_getObjectAttributesCmd).Standalone()

	clouddirectory_getObjectAttributesCmd.Flags().String("attribute-names", "", "List of attribute names whose values will be retrieved.")
	clouddirectory_getObjectAttributesCmd.Flags().String("consistency-level", "", "The consistency level at which to retrieve the attributes on an object.")
	clouddirectory_getObjectAttributesCmd.Flags().String("directory-arn", "", "The Amazon Resource Name (ARN) that is associated with the [Directory]() where the object resides.")
	clouddirectory_getObjectAttributesCmd.Flags().String("object-reference", "", "Reference that identifies the object whose attributes will be retrieved.")
	clouddirectory_getObjectAttributesCmd.Flags().String("schema-facet", "", "Identifier for the facet whose attributes will be retrieved.")
	clouddirectory_getObjectAttributesCmd.MarkFlagRequired("attribute-names")
	clouddirectory_getObjectAttributesCmd.MarkFlagRequired("directory-arn")
	clouddirectory_getObjectAttributesCmd.MarkFlagRequired("object-reference")
	clouddirectory_getObjectAttributesCmd.MarkFlagRequired("schema-facet")
	clouddirectoryCmd.AddCommand(clouddirectory_getObjectAttributesCmd)
}
