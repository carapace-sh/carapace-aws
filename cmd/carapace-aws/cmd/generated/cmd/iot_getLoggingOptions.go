package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_getLoggingOptionsCmd = &cobra.Command{
	Use:   "get-logging-options",
	Short: "Gets the logging options.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_getLoggingOptionsCmd).Standalone()

	iotCmd.AddCommand(iot_getLoggingOptionsCmd)
}
