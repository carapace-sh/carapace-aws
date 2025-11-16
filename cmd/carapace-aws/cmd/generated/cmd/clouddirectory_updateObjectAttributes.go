package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_updateObjectAttributesCmd = &cobra.Command{
	Use:   "update-object-attributes",
	Short: "Updates a given object's attributes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_updateObjectAttributesCmd).Standalone()

	clouddirectory_updateObjectAttributesCmd.Flags().String("attribute-updates", "", "The attributes update structure.")
	clouddirectory_updateObjectAttributesCmd.Flags().String("directory-arn", "", "The Amazon Resource Name (ARN) that is associated with the [Directory]() where the object resides.")
	clouddirectory_updateObjectAttributesCmd.Flags().String("object-reference", "", "The reference that identifies the object.")
	clouddirectory_updateObjectAttributesCmd.MarkFlagRequired("attribute-updates")
	clouddirectory_updateObjectAttributesCmd.MarkFlagRequired("directory-arn")
	clouddirectory_updateObjectAttributesCmd.MarkFlagRequired("object-reference")
	clouddirectoryCmd.AddCommand(clouddirectory_updateObjectAttributesCmd)
}
