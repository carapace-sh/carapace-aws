package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_describeEventConfigurationsCmd = &cobra.Command{
	Use:   "describe-event-configurations",
	Short: "Describes event configurations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_describeEventConfigurationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_describeEventConfigurationsCmd).Standalone()

	})
	iotCmd.AddCommand(iot_describeEventConfigurationsCmd)
}
