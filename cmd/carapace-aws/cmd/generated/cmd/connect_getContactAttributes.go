package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_getContactAttributesCmd = &cobra.Command{
	Use:   "get-contact-attributes",
	Short: "Retrieves the contact attributes for the specified contact.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_getContactAttributesCmd).Standalone()

	connect_getContactAttributesCmd.Flags().String("initial-contact-id", "", "The identifier of the initial contact.")
	connect_getContactAttributesCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_getContactAttributesCmd.MarkFlagRequired("initial-contact-id")
	connect_getContactAttributesCmd.MarkFlagRequired("instance-id")
	connectCmd.AddCommand(connect_getContactAttributesCmd)
}
