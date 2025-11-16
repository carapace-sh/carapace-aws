package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qapps_exportQappSessionDataCmd = &cobra.Command{
	Use:   "export-qapp-session-data",
	Short: "Exports the collected data of a Q App data collection session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qapps_exportQappSessionDataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qapps_exportQappSessionDataCmd).Standalone()

		qapps_exportQappSessionDataCmd.Flags().String("instance-id", "", "The unique identifier of the Amazon Q Business application environment instance.")
		qapps_exportQappSessionDataCmd.Flags().String("session-id", "", "The unique identifier of the Q App data collection session.")
		qapps_exportQappSessionDataCmd.MarkFlagRequired("instance-id")
		qapps_exportQappSessionDataCmd.MarkFlagRequired("session-id")
	})
	qappsCmd.AddCommand(qapps_exportQappSessionDataCmd)
}
