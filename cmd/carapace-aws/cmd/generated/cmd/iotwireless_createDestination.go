package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_createDestinationCmd = &cobra.Command{
	Use:   "create-destination",
	Short: "Creates a new destination that maps a device message to an AWS IoT rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_createDestinationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_createDestinationCmd).Standalone()

		iotwireless_createDestinationCmd.Flags().String("client-request-token", "", "Each resource must have a unique client request token.")
		iotwireless_createDestinationCmd.Flags().String("description", "", "The description of the new resource.")
		iotwireless_createDestinationCmd.Flags().String("expression", "", "The rule name or topic rule to send messages to.")
		iotwireless_createDestinationCmd.Flags().String("expression-type", "", "The type of value in `Expression`.")
		iotwireless_createDestinationCmd.Flags().String("name", "", "The name of the new resource.")
		iotwireless_createDestinationCmd.Flags().String("role-arn", "", "The ARN of the IAM Role that authorizes the destination.")
		iotwireless_createDestinationCmd.Flags().String("tags", "", "The tags to attach to the new destination.")
		iotwireless_createDestinationCmd.MarkFlagRequired("expression")
		iotwireless_createDestinationCmd.MarkFlagRequired("expression-type")
		iotwireless_createDestinationCmd.MarkFlagRequired("name")
		iotwireless_createDestinationCmd.MarkFlagRequired("role-arn")
	})
	iotwirelessCmd.AddCommand(iotwireless_createDestinationCmd)
}
