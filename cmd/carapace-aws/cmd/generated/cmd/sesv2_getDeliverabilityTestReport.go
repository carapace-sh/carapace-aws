package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_getDeliverabilityTestReportCmd = &cobra.Command{
	Use:   "get-deliverability-test-report",
	Short: "Retrieve the results of a predictive inbox placement test.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_getDeliverabilityTestReportCmd).Standalone()

	sesv2_getDeliverabilityTestReportCmd.Flags().String("report-id", "", "A unique string that identifies the predictive inbox placement test.")
	sesv2_getDeliverabilityTestReportCmd.MarkFlagRequired("report-id")
	sesv2Cmd.AddCommand(sesv2_getDeliverabilityTestReportCmd)
}
