package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_getLinkAttributesCmd = &cobra.Command{
	Use:   "get-link-attributes",
	Short: "Retrieves attributes that are associated with a typed link.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_getLinkAttributesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(clouddirectory_getLinkAttributesCmd).Standalone()

		clouddirectory_getLinkAttributesCmd.Flags().String("attribute-names", "", "A list of attribute names whose values will be retrieved.")
		clouddirectory_getLinkAttributesCmd.Flags().String("consistency-level", "", "The consistency level at which to retrieve the attributes on a typed link.")
		clouddirectory_getLinkAttributesCmd.Flags().String("directory-arn", "", "The Amazon Resource Name (ARN) that is associated with the Directory where the typed link resides.")
		clouddirectory_getLinkAttributesCmd.Flags().String("typed-link-specifier", "", "Allows a typed link specifier to be accepted as input.")
		clouddirectory_getLinkAttributesCmd.MarkFlagRequired("attribute-names")
		clouddirectory_getLinkAttributesCmd.MarkFlagRequired("directory-arn")
		clouddirectory_getLinkAttributesCmd.MarkFlagRequired("typed-link-specifier")
	})
	clouddirectoryCmd.AddCommand(clouddirectory_getLinkAttributesCmd)
}
