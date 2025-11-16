package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotdeviceadvisor_deleteSuiteDefinitionCmd = &cobra.Command{
	Use:   "delete-suite-definition",
	Short: "Deletes a Device Advisor test suite.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotdeviceadvisor_deleteSuiteDefinitionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotdeviceadvisor_deleteSuiteDefinitionCmd).Standalone()

		iotdeviceadvisor_deleteSuiteDefinitionCmd.Flags().String("suite-definition-id", "", "Suite definition ID of the test suite to be deleted.")
		iotdeviceadvisor_deleteSuiteDefinitionCmd.MarkFlagRequired("suite-definition-id")
	})
	iotdeviceadvisorCmd.AddCommand(iotdeviceadvisor_deleteSuiteDefinitionCmd)
}
