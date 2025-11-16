package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivsRealtime_getStageSessionCmd = &cobra.Command{
	Use:   "get-stage-session",
	Short: "Gets information for the specified stage session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivsRealtime_getStageSessionCmd).Standalone()

	ivsRealtime_getStageSessionCmd.Flags().String("session-id", "", "ID of a session within the stage.")
	ivsRealtime_getStageSessionCmd.Flags().String("stage-arn", "", "ARN of the stage for which the information is to be retrieved.")
	ivsRealtime_getStageSessionCmd.MarkFlagRequired("session-id")
	ivsRealtime_getStageSessionCmd.MarkFlagRequired("stage-arn")
	ivsRealtimeCmd.AddCommand(ivsRealtime_getStageSessionCmd)
}
