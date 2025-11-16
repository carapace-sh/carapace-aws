package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotevents_describeLoggingOptionsCmd = &cobra.Command{
	Use:   "describe-logging-options",
	Short: "Retrieves the current settings of the AWS IoT Events logging options.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotevents_describeLoggingOptionsCmd).Standalone()

	ioteventsCmd.AddCommand(iotevents_describeLoggingOptionsCmd)
}
