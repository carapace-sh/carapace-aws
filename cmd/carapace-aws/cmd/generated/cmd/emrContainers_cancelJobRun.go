package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emrContainers_cancelJobRunCmd = &cobra.Command{
	Use:   "cancel-job-run",
	Short: "Cancels a job run.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emrContainers_cancelJobRunCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(emrContainers_cancelJobRunCmd).Standalone()

		emrContainers_cancelJobRunCmd.Flags().String("id", "", "The ID of the job run to cancel.")
		emrContainers_cancelJobRunCmd.Flags().String("virtual-cluster-id", "", "The ID of the virtual cluster for which the job run will be canceled.")
		emrContainers_cancelJobRunCmd.MarkFlagRequired("id")
		emrContainers_cancelJobRunCmd.MarkFlagRequired("virtual-cluster-id")
	})
	emrContainersCmd.AddCommand(emrContainers_cancelJobRunCmd)
}
