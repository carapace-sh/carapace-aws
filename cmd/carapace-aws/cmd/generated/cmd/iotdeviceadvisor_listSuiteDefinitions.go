package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotdeviceadvisor_listSuiteDefinitionsCmd = &cobra.Command{
	Use:   "list-suite-definitions",
	Short: "Lists the Device Advisor test suites you have created.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotdeviceadvisor_listSuiteDefinitionsCmd).Standalone()

	iotdeviceadvisor_listSuiteDefinitionsCmd.Flags().String("max-results", "", "The maximum number of results to return at once.")
	iotdeviceadvisor_listSuiteDefinitionsCmd.Flags().String("next-token", "", "A token used to get the next set of results.")
	iotdeviceadvisorCmd.AddCommand(iotdeviceadvisor_listSuiteDefinitionsCmd)
}
