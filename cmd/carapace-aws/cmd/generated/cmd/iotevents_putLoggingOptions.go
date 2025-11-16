package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotevents_putLoggingOptionsCmd = &cobra.Command{
	Use:   "put-logging-options",
	Short: "Sets or updates the AWS IoT Events logging options.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotevents_putLoggingOptionsCmd).Standalone()

	iotevents_putLoggingOptionsCmd.Flags().String("logging-options", "", "The new values of the AWS IoT Events logging options.")
	iotevents_putLoggingOptionsCmd.MarkFlagRequired("logging-options")
	ioteventsCmd.AddCommand(iotevents_putLoggingOptionsCmd)
}
