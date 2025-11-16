package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_terminateJobFlowsCmd = &cobra.Command{
	Use:   "terminate-job-flows",
	Short: "TerminateJobFlows shuts a list of clusters (job flows) down.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_terminateJobFlowsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(emr_terminateJobFlowsCmd).Standalone()

		emr_terminateJobFlowsCmd.Flags().String("job-flow-ids", "", "A list of job flows to be shut down.")
		emr_terminateJobFlowsCmd.MarkFlagRequired("job-flow-ids")
	})
	emrCmd.AddCommand(emr_terminateJobFlowsCmd)
}
