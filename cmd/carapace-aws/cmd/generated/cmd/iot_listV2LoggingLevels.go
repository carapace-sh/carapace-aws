package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listV2LoggingLevelsCmd = &cobra.Command{
	Use:   "list-v2-logging-levels",
	Short: "Lists logging levels.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listV2LoggingLevelsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_listV2LoggingLevelsCmd).Standalone()

		iot_listV2LoggingLevelsCmd.Flags().String("max-results", "", "The maximum number of results to return at one time.")
		iot_listV2LoggingLevelsCmd.Flags().String("next-token", "", "To retrieve the next set of results, the `nextToken` value from a previous response; otherwise **null** to receive the first set of results.")
		iot_listV2LoggingLevelsCmd.Flags().String("target-type", "", "The type of resource for which you are configuring logging.")
	})
	iotCmd.AddCommand(iot_listV2LoggingLevelsCmd)
}
