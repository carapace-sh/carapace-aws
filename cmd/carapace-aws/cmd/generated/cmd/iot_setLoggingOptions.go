package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_setLoggingOptionsCmd = &cobra.Command{
	Use:   "set-logging-options",
	Short: "Sets the logging options.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_setLoggingOptionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_setLoggingOptionsCmd).Standalone()

		iot_setLoggingOptionsCmd.Flags().String("logging-options-payload", "", "The logging options payload.")
		iot_setLoggingOptionsCmd.MarkFlagRequired("logging-options-payload")
	})
	iotCmd.AddCommand(iot_setLoggingOptionsCmd)
}
