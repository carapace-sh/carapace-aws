package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_createObjectCmd = &cobra.Command{
	Use:   "create-object",
	Short: "Creates an object in a [Directory]().",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_createObjectCmd).Standalone()

	clouddirectory_createObjectCmd.Flags().String("directory-arn", "", "The Amazon Resource Name (ARN) that is associated with the [Directory]() in which the object will be created.")
	clouddirectory_createObjectCmd.Flags().String("link-name", "", "The name of link that is used to attach this object to a parent.")
	clouddirectory_createObjectCmd.Flags().String("object-attribute-list", "", "The attribute map whose attribute ARN contains the key and attribute value as the map value.")
	clouddirectory_createObjectCmd.Flags().String("parent-reference", "", "If specified, the parent reference to which this object will be attached.")
	clouddirectory_createObjectCmd.Flags().String("schema-facets", "", "A list of schema facets to be associated with the object.")
	clouddirectory_createObjectCmd.MarkFlagRequired("directory-arn")
	clouddirectory_createObjectCmd.MarkFlagRequired("schema-facets")
	clouddirectoryCmd.AddCommand(clouddirectory_createObjectCmd)
}
