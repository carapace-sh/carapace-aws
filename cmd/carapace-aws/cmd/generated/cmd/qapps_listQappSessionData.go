package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qapps_listQappSessionDataCmd = &cobra.Command{
	Use:   "list-qapp-session-data",
	Short: "Lists the collected data of a Q App data collection session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qapps_listQappSessionDataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qapps_listQappSessionDataCmd).Standalone()

		qapps_listQappSessionDataCmd.Flags().String("instance-id", "", "The unique identifier of the Amazon Q Business application environment instance.")
		qapps_listQappSessionDataCmd.Flags().String("session-id", "", "The unique identifier of the Q App data collection session.")
		qapps_listQappSessionDataCmd.MarkFlagRequired("instance-id")
		qapps_listQappSessionDataCmd.MarkFlagRequired("session-id")
	})
	qappsCmd.AddCommand(qapps_listQappSessionDataCmd)
}
