package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var launchWizard_getWorkloadCmd = &cobra.Command{
	Use:   "get-workload",
	Short: "Returns information about a workload.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(launchWizard_getWorkloadCmd).Standalone()

	launchWizard_getWorkloadCmd.Flags().String("workload-name", "", "The name of the workload.")
	launchWizard_getWorkloadCmd.MarkFlagRequired("workload-name")
	launchWizardCmd.AddCommand(launchWizard_getWorkloadCmd)
}
