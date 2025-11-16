package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_createDeliverabilityTestReportCmd = &cobra.Command{
	Use:   "create-deliverability-test-report",
	Short: "Create a new predictive inbox placement test.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_createDeliverabilityTestReportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sesv2_createDeliverabilityTestReportCmd).Standalone()

		sesv2_createDeliverabilityTestReportCmd.Flags().String("content", "", "The HTML body of the message that you sent when you performed the predictive inbox placement test.")
		sesv2_createDeliverabilityTestReportCmd.Flags().String("from-email-address", "", "The email address that the predictive inbox placement test email was sent from.")
		sesv2_createDeliverabilityTestReportCmd.Flags().String("report-name", "", "A unique name that helps you to identify the predictive inbox placement test when you retrieve the results.")
		sesv2_createDeliverabilityTestReportCmd.Flags().String("tags", "", "An array of objects that define the tags (keys and values) that you want to associate with the predictive inbox placement test.")
		sesv2_createDeliverabilityTestReportCmd.MarkFlagRequired("content")
		sesv2_createDeliverabilityTestReportCmd.MarkFlagRequired("from-email-address")
	})
	sesv2Cmd.AddCommand(sesv2_createDeliverabilityTestReportCmd)
}
