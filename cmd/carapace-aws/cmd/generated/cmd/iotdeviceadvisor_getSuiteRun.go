package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotdeviceadvisor_getSuiteRunCmd = &cobra.Command{
	Use:   "get-suite-run",
	Short: "Gets information about a Device Advisor test suite run.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotdeviceadvisor_getSuiteRunCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotdeviceadvisor_getSuiteRunCmd).Standalone()

		iotdeviceadvisor_getSuiteRunCmd.Flags().String("suite-definition-id", "", "Suite definition ID for the test suite run.")
		iotdeviceadvisor_getSuiteRunCmd.Flags().String("suite-run-id", "", "Suite run ID for the test suite run.")
		iotdeviceadvisor_getSuiteRunCmd.MarkFlagRequired("suite-definition-id")
		iotdeviceadvisor_getSuiteRunCmd.MarkFlagRequired("suite-run-id")
	})
	iotdeviceadvisorCmd.AddCommand(iotdeviceadvisor_getSuiteRunCmd)
}
