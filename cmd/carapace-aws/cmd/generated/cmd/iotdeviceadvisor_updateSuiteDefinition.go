package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotdeviceadvisor_updateSuiteDefinitionCmd = &cobra.Command{
	Use:   "update-suite-definition",
	Short: "Updates a Device Advisor test suite.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotdeviceadvisor_updateSuiteDefinitionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotdeviceadvisor_updateSuiteDefinitionCmd).Standalone()

		iotdeviceadvisor_updateSuiteDefinitionCmd.Flags().String("suite-definition-configuration", "", "Updates a Device Advisor test suite with suite definition configuration.")
		iotdeviceadvisor_updateSuiteDefinitionCmd.Flags().String("suite-definition-id", "", "Suite definition ID of the test suite to be updated.")
		iotdeviceadvisor_updateSuiteDefinitionCmd.MarkFlagRequired("suite-definition-configuration")
		iotdeviceadvisor_updateSuiteDefinitionCmd.MarkFlagRequired("suite-definition-id")
	})
	iotdeviceadvisorCmd.AddCommand(iotdeviceadvisor_updateSuiteDefinitionCmd)
}
