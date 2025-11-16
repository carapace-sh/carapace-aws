package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivsRealtime_listStageSessionsCmd = &cobra.Command{
	Use:   "list-stage-sessions",
	Short: "Gets all sessions for a specified stage.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivsRealtime_listStageSessionsCmd).Standalone()

	ivsRealtime_listStageSessionsCmd.Flags().String("max-results", "", "Maximum number of results to return.")
	ivsRealtime_listStageSessionsCmd.Flags().String("next-token", "", "The first stage session to retrieve.")
	ivsRealtime_listStageSessionsCmd.Flags().String("stage-arn", "", "Stage ARN.")
	ivsRealtime_listStageSessionsCmd.MarkFlagRequired("stage-arn")
	ivsRealtimeCmd.AddCommand(ivsRealtime_listStageSessionsCmd)
}
