package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotdeviceadvisor_getSuiteDefinitionCmd = &cobra.Command{
	Use:   "get-suite-definition",
	Short: "Gets information about a Device Advisor test suite.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotdeviceadvisor_getSuiteDefinitionCmd).Standalone()

	iotdeviceadvisor_getSuiteDefinitionCmd.Flags().String("suite-definition-id", "", "Suite definition ID of the test suite to get.")
	iotdeviceadvisor_getSuiteDefinitionCmd.Flags().String("suite-definition-version", "", "Suite definition version of the test suite to get.")
	iotdeviceadvisor_getSuiteDefinitionCmd.MarkFlagRequired("suite-definition-id")
	iotdeviceadvisorCmd.AddCommand(iotdeviceadvisor_getSuiteDefinitionCmd)
}
