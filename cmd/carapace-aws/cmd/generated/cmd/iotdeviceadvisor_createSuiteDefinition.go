package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotdeviceadvisor_createSuiteDefinitionCmd = &cobra.Command{
	Use:   "create-suite-definition",
	Short: "Creates a Device Advisor test suite.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotdeviceadvisor_createSuiteDefinitionCmd).Standalone()

	iotdeviceadvisor_createSuiteDefinitionCmd.Flags().String("client-token", "", "The client token for the test suite definition creation.")
	iotdeviceadvisor_createSuiteDefinitionCmd.Flags().String("suite-definition-configuration", "", "Creates a Device Advisor test suite with suite definition configuration.")
	iotdeviceadvisor_createSuiteDefinitionCmd.Flags().String("tags", "", "The tags to be attached to the suite definition.")
	iotdeviceadvisor_createSuiteDefinitionCmd.MarkFlagRequired("suite-definition-configuration")
	iotdeviceadvisorCmd.AddCommand(iotdeviceadvisor_createSuiteDefinitionCmd)
}
