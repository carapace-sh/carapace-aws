package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var groundstation_describeContactCmd = &cobra.Command{
	Use:   "describe-contact",
	Short: "Describes an existing contact.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(groundstation_describeContactCmd).Standalone()

	groundstation_describeContactCmd.Flags().String("contact-id", "", "UUID of a contact.")
	groundstation_describeContactCmd.MarkFlagRequired("contact-id")
	groundstationCmd.AddCommand(groundstation_describeContactCmd)
}
