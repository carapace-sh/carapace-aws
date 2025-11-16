package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointEmail_putDeliverabilityDashboardOptionCmd = &cobra.Command{
	Use:   "put-deliverability-dashboard-option",
	Short: "Enable or disable the Deliverability dashboard for your Amazon Pinpoint account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointEmail_putDeliverabilityDashboardOptionCmd).Standalone()

	pinpointEmail_putDeliverabilityDashboardOptionCmd.Flags().String("dashboard-enabled", "", "Specifies whether to enable the Deliverability dashboard for your Amazon Pinpoint account.")
	pinpointEmail_putDeliverabilityDashboardOptionCmd.Flags().String("subscribed-domains", "", "An array of objects, one for each verified domain that you use to send email and enabled the Deliverability dashboard for.")
	pinpointEmail_putDeliverabilityDashboardOptionCmd.MarkFlagRequired("dashboard-enabled")
	pinpointEmailCmd.AddCommand(pinpointEmail_putDeliverabilityDashboardOptionCmd)
}
