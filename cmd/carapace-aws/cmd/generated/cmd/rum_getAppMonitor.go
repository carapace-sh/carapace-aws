package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rum_getAppMonitorCmd = &cobra.Command{
	Use:   "get-app-monitor",
	Short: "Retrieves the complete configuration information for one app monitor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rum_getAppMonitorCmd).Standalone()

	rum_getAppMonitorCmd.Flags().String("name", "", "The app monitor to retrieve information for.")
	rum_getAppMonitorCmd.MarkFlagRequired("name")
	rumCmd.AddCommand(rum_getAppMonitorCmd)
}
