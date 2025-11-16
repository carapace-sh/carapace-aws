package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotdeviceadvisor_startSuiteRunCmd = &cobra.Command{
	Use:   "start-suite-run",
	Short: "Starts a Device Advisor test suite run.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotdeviceadvisor_startSuiteRunCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotdeviceadvisor_startSuiteRunCmd).Standalone()

		iotdeviceadvisor_startSuiteRunCmd.Flags().String("suite-definition-id", "", "Suite definition ID of the test suite.")
		iotdeviceadvisor_startSuiteRunCmd.Flags().String("suite-definition-version", "", "Suite definition version of the test suite.")
		iotdeviceadvisor_startSuiteRunCmd.Flags().String("suite-run-configuration", "", "Suite run configuration.")
		iotdeviceadvisor_startSuiteRunCmd.Flags().String("tags", "", "The tags to be attached to the suite run.")
		iotdeviceadvisor_startSuiteRunCmd.MarkFlagRequired("suite-definition-id")
		iotdeviceadvisor_startSuiteRunCmd.MarkFlagRequired("suite-run-configuration")
	})
	iotdeviceadvisorCmd.AddCommand(iotdeviceadvisor_startSuiteRunCmd)
}
