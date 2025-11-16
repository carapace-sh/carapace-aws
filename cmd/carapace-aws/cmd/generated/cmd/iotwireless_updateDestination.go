package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_updateDestinationCmd = &cobra.Command{
	Use:   "update-destination",
	Short: "Updates properties of a destination.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_updateDestinationCmd).Standalone()

	iotwireless_updateDestinationCmd.Flags().String("description", "", "A new description of the resource.")
	iotwireless_updateDestinationCmd.Flags().String("expression", "", "The new rule name or topic rule to send messages to.")
	iotwireless_updateDestinationCmd.Flags().String("expression-type", "", "The type of value in `Expression`.")
	iotwireless_updateDestinationCmd.Flags().String("name", "", "The new name of the resource.")
	iotwireless_updateDestinationCmd.Flags().String("role-arn", "", "The ARN of the IAM Role that authorizes the destination.")
	iotwireless_updateDestinationCmd.MarkFlagRequired("name")
	iotwirelessCmd.AddCommand(iotwireless_updateDestinationCmd)
}
