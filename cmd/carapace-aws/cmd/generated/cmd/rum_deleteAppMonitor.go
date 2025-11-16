package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rum_deleteAppMonitorCmd = &cobra.Command{
	Use:   "delete-app-monitor",
	Short: "Deletes an existing app monitor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rum_deleteAppMonitorCmd).Standalone()

	rum_deleteAppMonitorCmd.Flags().String("name", "", "The name of the app monitor to delete.")
	rum_deleteAppMonitorCmd.MarkFlagRequired("name")
	rumCmd.AddCommand(rum_deleteAppMonitorCmd)
}
