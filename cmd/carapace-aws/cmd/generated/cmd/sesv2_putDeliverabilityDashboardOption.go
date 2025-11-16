package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_putDeliverabilityDashboardOptionCmd = &cobra.Command{
	Use:   "put-deliverability-dashboard-option",
	Short: "Enable or disable the Deliverability dashboard.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_putDeliverabilityDashboardOptionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sesv2_putDeliverabilityDashboardOptionCmd).Standalone()

		sesv2_putDeliverabilityDashboardOptionCmd.Flags().String("dashboard-enabled", "", "Specifies whether to enable the Deliverability dashboard.")
		sesv2_putDeliverabilityDashboardOptionCmd.Flags().String("subscribed-domains", "", "An array of objects, one for each verified domain that you use to send email and enabled the Deliverability dashboard for.")
		sesv2_putDeliverabilityDashboardOptionCmd.MarkFlagRequired("dashboard-enabled")
	})
	sesv2Cmd.AddCommand(sesv2_putDeliverabilityDashboardOptionCmd)
}
