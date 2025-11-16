package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_updateContactAttributesCmd = &cobra.Command{
	Use:   "update-contact-attributes",
	Short: "Creates or updates user-defined contact attributes associated with the specified contact.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_updateContactAttributesCmd).Standalone()

	connect_updateContactAttributesCmd.Flags().String("attributes", "", "The Amazon Connect attributes.")
	connect_updateContactAttributesCmd.Flags().String("initial-contact-id", "", "The identifier of the contact.")
	connect_updateContactAttributesCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_updateContactAttributesCmd.MarkFlagRequired("attributes")
	connect_updateContactAttributesCmd.MarkFlagRequired("initial-contact-id")
	connect_updateContactAttributesCmd.MarkFlagRequired("instance-id")
	connectCmd.AddCommand(connect_updateContactAttributesCmd)
}
