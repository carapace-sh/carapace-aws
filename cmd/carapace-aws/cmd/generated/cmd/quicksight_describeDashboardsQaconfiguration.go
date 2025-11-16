package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_describeDashboardsQaconfigurationCmd = &cobra.Command{
	Use:   "describe-dashboards-qaconfiguration",
	Short: "Describes an existing dashboard QA configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_describeDashboardsQaconfigurationCmd).Standalone()

	quicksight_describeDashboardsQaconfigurationCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the dashboard QA configuration that you want described.")
	quicksight_describeDashboardsQaconfigurationCmd.MarkFlagRequired("aws-account-id")
	quicksightCmd.AddCommand(quicksight_describeDashboardsQaconfigurationCmd)
}
