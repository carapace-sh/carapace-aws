package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emrContainers_describeJobRunCmd = &cobra.Command{
	Use:   "describe-job-run",
	Short: "Displays detailed information about a job run.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emrContainers_describeJobRunCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(emrContainers_describeJobRunCmd).Standalone()

		emrContainers_describeJobRunCmd.Flags().String("id", "", "The ID of the job run request.")
		emrContainers_describeJobRunCmd.Flags().String("virtual-cluster-id", "", "The ID of the virtual cluster for which the job run is submitted.")
		emrContainers_describeJobRunCmd.MarkFlagRequired("id")
		emrContainers_describeJobRunCmd.MarkFlagRequired("virtual-cluster-id")
	})
	emrContainersCmd.AddCommand(emrContainers_describeJobRunCmd)
}
