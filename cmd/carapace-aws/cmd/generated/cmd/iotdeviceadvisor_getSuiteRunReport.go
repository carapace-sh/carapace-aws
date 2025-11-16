package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotdeviceadvisor_getSuiteRunReportCmd = &cobra.Command{
	Use:   "get-suite-run-report",
	Short: "Gets a report download link for a successful Device Advisor qualifying test suite run.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotdeviceadvisor_getSuiteRunReportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotdeviceadvisor_getSuiteRunReportCmd).Standalone()

		iotdeviceadvisor_getSuiteRunReportCmd.Flags().String("suite-definition-id", "", "Suite definition ID of the test suite.")
		iotdeviceadvisor_getSuiteRunReportCmd.Flags().String("suite-run-id", "", "Suite run ID of the test suite run.")
		iotdeviceadvisor_getSuiteRunReportCmd.MarkFlagRequired("suite-definition-id")
		iotdeviceadvisor_getSuiteRunReportCmd.MarkFlagRequired("suite-run-id")
	})
	iotdeviceadvisorCmd.AddCommand(iotdeviceadvisor_getSuiteRunReportCmd)
}
