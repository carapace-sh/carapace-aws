package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotdeviceadvisor_listSuiteRunsCmd = &cobra.Command{
	Use:   "list-suite-runs",
	Short: "Lists runs of the specified Device Advisor test suite.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotdeviceadvisor_listSuiteRunsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotdeviceadvisor_listSuiteRunsCmd).Standalone()

		iotdeviceadvisor_listSuiteRunsCmd.Flags().String("max-results", "", "The maximum number of results to return at once.")
		iotdeviceadvisor_listSuiteRunsCmd.Flags().String("next-token", "", "A token to retrieve the next set of results.")
		iotdeviceadvisor_listSuiteRunsCmd.Flags().String("suite-definition-id", "", "Lists the test suite runs of the specified test suite based on suite definition ID.")
		iotdeviceadvisor_listSuiteRunsCmd.Flags().String("suite-definition-version", "", "Must be passed along with `suiteDefinitionId`.")
	})
	iotdeviceadvisorCmd.AddCommand(iotdeviceadvisor_listSuiteRunsCmd)
}
