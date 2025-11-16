package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_stopRunCmd = &cobra.Command{
	Use:   "stop-run",
	Short: "Initiates a stop request for the current test run.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_stopRunCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devicefarm_stopRunCmd).Standalone()

		devicefarm_stopRunCmd.Flags().String("arn", "", "Represents the Amazon Resource Name (ARN) of the Device Farm run to stop.")
		devicefarm_stopRunCmd.MarkFlagRequired("arn")
	})
	devicefarmCmd.AddCommand(devicefarm_stopRunCmd)
}
