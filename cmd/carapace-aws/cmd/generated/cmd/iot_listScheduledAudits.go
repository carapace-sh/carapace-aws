package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listScheduledAuditsCmd = &cobra.Command{
	Use:   "list-scheduled-audits",
	Short: "Lists all of your scheduled audits.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listScheduledAuditsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_listScheduledAuditsCmd).Standalone()

		iot_listScheduledAuditsCmd.Flags().String("max-results", "", "The maximum number of results to return at one time.")
		iot_listScheduledAuditsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	})
	iotCmd.AddCommand(iot_listScheduledAuditsCmd)
}
