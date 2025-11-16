package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivsRealtime_listStagesCmd = &cobra.Command{
	Use:   "list-stages",
	Short: "Gets summary information about all stages in your account, in the AWS region where the API request is processed.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivsRealtime_listStagesCmd).Standalone()

	ivsRealtime_listStagesCmd.Flags().String("max-results", "", "Maximum number of results to return.")
	ivsRealtime_listStagesCmd.Flags().String("next-token", "", "The first stage to retrieve.")
	ivsRealtimeCmd.AddCommand(ivsRealtime_listStagesCmd)
}
