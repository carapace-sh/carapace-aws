package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmDashboardsCmd = &cobra.Command{
	Use:   "bcm-dashboards",
	Short: "Amazon Web Services Billing and Cost Management Dashboards is a service that enables you to create, manage, and share dashboards that combine multiple visualizations of your Amazon Web Services cost and usage data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmDashboardsCmd).Standalone()

	rootCmd.AddCommand(bcmDashboardsCmd)
}
