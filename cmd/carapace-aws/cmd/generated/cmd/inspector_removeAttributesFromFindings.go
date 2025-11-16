package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector_removeAttributesFromFindingsCmd = &cobra.Command{
	Use:   "remove-attributes-from-findings",
	Short: "Removes entire attributes (key and value pairs) from the findings that are specified by the ARNs of the findings where an attribute with the specified key exists.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector_removeAttributesFromFindingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector_removeAttributesFromFindingsCmd).Standalone()

		inspector_removeAttributesFromFindingsCmd.Flags().String("attribute-keys", "", "The array of attribute keys that you want to remove from specified findings.")
		inspector_removeAttributesFromFindingsCmd.Flags().String("finding-arns", "", "The ARNs that specify the findings that you want to remove attributes from.")
		inspector_removeAttributesFromFindingsCmd.MarkFlagRequired("attribute-keys")
		inspector_removeAttributesFromFindingsCmd.MarkFlagRequired("finding-arns")
	})
	inspectorCmd.AddCommand(inspector_removeAttributesFromFindingsCmd)
}
