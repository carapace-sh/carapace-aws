package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_getSendStatisticsCmd = &cobra.Command{
	Use:   "get-send-statistics",
	Short: "Provides sending statistics for the current Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_getSendStatisticsCmd).Standalone()

	sesCmd.AddCommand(ses_getSendStatisticsCmd)
}
