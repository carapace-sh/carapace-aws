package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointEmail_getDeliverabilityTestReportCmd = &cobra.Command{
	Use:   "get-deliverability-test-report",
	Short: "Retrieve the results of a predictive inbox placement test.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointEmail_getDeliverabilityTestReportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpointEmail_getDeliverabilityTestReportCmd).Standalone()

		pinpointEmail_getDeliverabilityTestReportCmd.Flags().String("report-id", "", "A unique string that identifies the predictive inbox placement test.")
		pinpointEmail_getDeliverabilityTestReportCmd.MarkFlagRequired("report-id")
	})
	pinpointEmailCmd.AddCommand(pinpointEmail_getDeliverabilityTestReportCmd)
}
