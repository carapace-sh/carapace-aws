package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emrContainers_listJobRunsCmd = &cobra.Command{
	Use:   "list-job-runs",
	Short: "Lists job runs based on a set of parameters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emrContainers_listJobRunsCmd).Standalone()

	emrContainers_listJobRunsCmd.Flags().String("created-after", "", "The date and time after which the job runs were submitted.")
	emrContainers_listJobRunsCmd.Flags().String("created-before", "", "The date and time before which the job runs were submitted.")
	emrContainers_listJobRunsCmd.Flags().String("max-results", "", "The maximum number of job runs that can be listed.")
	emrContainers_listJobRunsCmd.Flags().String("name", "", "The name of the job run.")
	emrContainers_listJobRunsCmd.Flags().String("next-token", "", "The token for the next set of job runs to return.")
	emrContainers_listJobRunsCmd.Flags().String("states", "", "The states of the job run.")
	emrContainers_listJobRunsCmd.Flags().String("virtual-cluster-id", "", "The ID of the virtual cluster for which to list the job run.")
	emrContainers_listJobRunsCmd.MarkFlagRequired("virtual-cluster-id")
	emrContainersCmd.AddCommand(emrContainers_listJobRunsCmd)
}
