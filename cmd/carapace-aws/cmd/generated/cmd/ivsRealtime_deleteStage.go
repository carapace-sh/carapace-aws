package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivsRealtime_deleteStageCmd = &cobra.Command{
	Use:   "delete-stage",
	Short: "Shuts down and deletes the specified stage (disconnecting all participants).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivsRealtime_deleteStageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ivsRealtime_deleteStageCmd).Standalone()

		ivsRealtime_deleteStageCmd.Flags().String("arn", "", "ARN of the stage to be deleted.")
		ivsRealtime_deleteStageCmd.MarkFlagRequired("arn")
	})
	ivsRealtimeCmd.AddCommand(ivsRealtime_deleteStageCmd)
}
