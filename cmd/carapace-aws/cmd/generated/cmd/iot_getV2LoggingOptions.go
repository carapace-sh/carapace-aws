package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_getV2LoggingOptionsCmd = &cobra.Command{
	Use:   "get-v2-logging-options",
	Short: "Gets the fine grained logging options.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_getV2LoggingOptionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_getV2LoggingOptionsCmd).Standalone()

	})
	iotCmd.AddCommand(iot_getV2LoggingOptionsCmd)
}
