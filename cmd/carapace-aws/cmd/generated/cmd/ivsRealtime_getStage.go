package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivsRealtime_getStageCmd = &cobra.Command{
	Use:   "get-stage",
	Short: "Gets information for the specified stage.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivsRealtime_getStageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ivsRealtime_getStageCmd).Standalone()

		ivsRealtime_getStageCmd.Flags().String("arn", "", "ARN of the stage for which the information is to be retrieved.")
		ivsRealtime_getStageCmd.MarkFlagRequired("arn")
	})
	ivsRealtimeCmd.AddCommand(ivsRealtime_getStageCmd)
}
