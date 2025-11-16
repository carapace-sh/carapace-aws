package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivsRealtime_createStageCmd = &cobra.Command{
	Use:   "create-stage",
	Short: "Creates a new stage (and optionally participant tokens).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivsRealtime_createStageCmd).Standalone()

	ivsRealtime_createStageCmd.Flags().String("auto-participant-recording-configuration", "", "Configuration object for individual participant recording, to attach to the new stage.")
	ivsRealtime_createStageCmd.Flags().String("name", "", "Optional name that can be specified for the stage being created.")
	ivsRealtime_createStageCmd.Flags().String("participant-token-configurations", "", "Array of participant token configuration objects to attach to the new stage.")
	ivsRealtime_createStageCmd.Flags().String("tags", "", "Tags attached to the resource.")
	ivsRealtimeCmd.AddCommand(ivsRealtime_createStageCmd)
}
