package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var launchWizard_listWorkloadsCmd = &cobra.Command{
	Use:   "list-workloads",
	Short: "Lists the available workload names.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(launchWizard_listWorkloadsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(launchWizard_listWorkloadsCmd).Standalone()

		launchWizard_listWorkloadsCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
		launchWizard_listWorkloadsCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
	})
	launchWizardCmd.AddCommand(launchWizard_listWorkloadsCmd)
}
