package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivsRealtime_updateStageCmd = &cobra.Command{
	Use:   "update-stage",
	Short: "Updates a stage’s configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivsRealtime_updateStageCmd).Standalone()

	ivsRealtime_updateStageCmd.Flags().String("arn", "", "ARN of the stage to be updated.")
	ivsRealtime_updateStageCmd.Flags().String("auto-participant-recording-configuration", "", "Configuration object for individual participant recording, to attach to the stage.")
	ivsRealtime_updateStageCmd.Flags().String("name", "", "Name of the stage to be updated.")
	ivsRealtime_updateStageCmd.MarkFlagRequired("arn")
	ivsRealtimeCmd.AddCommand(ivsRealtime_updateStageCmd)
}
