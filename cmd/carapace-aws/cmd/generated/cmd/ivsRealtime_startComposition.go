package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivsRealtime_startCompositionCmd = &cobra.Command{
	Use:   "start-composition",
	Short: "Starts a Composition from a stage based on the configuration provided in the request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivsRealtime_startCompositionCmd).Standalone()

	ivsRealtime_startCompositionCmd.Flags().String("destinations", "", "Array of destination configuration.")
	ivsRealtime_startCompositionCmd.Flags().String("idempotency-token", "", "Idempotency token.")
	ivsRealtime_startCompositionCmd.Flags().String("layout", "", "Layout object to configure composition parameters.")
	ivsRealtime_startCompositionCmd.Flags().String("stage-arn", "", "ARN of the stage to be used for compositing.")
	ivsRealtime_startCompositionCmd.Flags().String("tags", "", "Tags attached to the resource.")
	ivsRealtime_startCompositionCmd.MarkFlagRequired("destinations")
	ivsRealtime_startCompositionCmd.MarkFlagRequired("stage-arn")
	ivsRealtimeCmd.AddCommand(ivsRealtime_startCompositionCmd)
}
