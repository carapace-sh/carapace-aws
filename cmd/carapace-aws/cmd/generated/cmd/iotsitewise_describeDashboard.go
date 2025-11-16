package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_describeDashboardCmd = &cobra.Command{
	Use:   "describe-dashboard",
	Short: "Retrieves information about a dashboard.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_describeDashboardCmd).Standalone()

	iotsitewise_describeDashboardCmd.Flags().String("dashboard-id", "", "The ID of the dashboard.")
	iotsitewise_describeDashboardCmd.MarkFlagRequired("dashboard-id")
	iotsitewiseCmd.AddCommand(iotsitewise_describeDashboardCmd)
}
