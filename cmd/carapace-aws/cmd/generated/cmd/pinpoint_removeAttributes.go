package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_removeAttributesCmd = &cobra.Command{
	Use:   "remove-attributes",
	Short: "Removes one or more custom attributes, of the same attribute type, from the application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_removeAttributesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpoint_removeAttributesCmd).Standalone()

		pinpoint_removeAttributesCmd.Flags().String("application-id", "", "The unique identifier for the application.")
		pinpoint_removeAttributesCmd.Flags().String("attribute-type", "", "The type of attribute or attributes to remove.")
		pinpoint_removeAttributesCmd.Flags().String("update-attributes-request", "", "")
		pinpoint_removeAttributesCmd.MarkFlagRequired("application-id")
		pinpoint_removeAttributesCmd.MarkFlagRequired("attribute-type")
		pinpoint_removeAttributesCmd.MarkFlagRequired("update-attributes-request")
	})
	pinpointCmd.AddCommand(pinpoint_removeAttributesCmd)
}
