package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotdeviceadvisor_stopSuiteRunCmd = &cobra.Command{
	Use:   "stop-suite-run",
	Short: "Stops a Device Advisor test suite run that is currently running.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotdeviceadvisor_stopSuiteRunCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotdeviceadvisor_stopSuiteRunCmd).Standalone()

		iotdeviceadvisor_stopSuiteRunCmd.Flags().String("suite-definition-id", "", "Suite definition ID of the test suite run to be stopped.")
		iotdeviceadvisor_stopSuiteRunCmd.Flags().String("suite-run-id", "", "Suite run ID of the test suite run to be stopped.")
		iotdeviceadvisor_stopSuiteRunCmd.MarkFlagRequired("suite-definition-id")
		iotdeviceadvisor_stopSuiteRunCmd.MarkFlagRequired("suite-run-id")
	})
	iotdeviceadvisorCmd.AddCommand(iotdeviceadvisor_stopSuiteRunCmd)
}
