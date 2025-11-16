package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_updateLinkAttributesCmd = &cobra.Command{
	Use:   "update-link-attributes",
	Short: "Updates a given typed link’s attributes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_updateLinkAttributesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(clouddirectory_updateLinkAttributesCmd).Standalone()

		clouddirectory_updateLinkAttributesCmd.Flags().String("attribute-updates", "", "The attributes update structure.")
		clouddirectory_updateLinkAttributesCmd.Flags().String("directory-arn", "", "The Amazon Resource Name (ARN) that is associated with the Directory where the updated typed link resides.")
		clouddirectory_updateLinkAttributesCmd.Flags().String("typed-link-specifier", "", "Allows a typed link specifier to be accepted as input.")
		clouddirectory_updateLinkAttributesCmd.MarkFlagRequired("attribute-updates")
		clouddirectory_updateLinkAttributesCmd.MarkFlagRequired("directory-arn")
		clouddirectory_updateLinkAttributesCmd.MarkFlagRequired("typed-link-specifier")
	})
	clouddirectoryCmd.AddCommand(clouddirectory_updateLinkAttributesCmd)
}
