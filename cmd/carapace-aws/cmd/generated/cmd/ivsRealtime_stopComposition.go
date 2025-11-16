package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivsRealtime_stopCompositionCmd = &cobra.Command{
	Use:   "stop-composition",
	Short: "Stops and deletes a Composition resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivsRealtime_stopCompositionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ivsRealtime_stopCompositionCmd).Standalone()

		ivsRealtime_stopCompositionCmd.Flags().String("arn", "", "ARN of the Composition.")
		ivsRealtime_stopCompositionCmd.MarkFlagRequired("arn")
	})
	ivsRealtimeCmd.AddCommand(ivsRealtime_stopCompositionCmd)
}
