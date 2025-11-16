package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_describeJobFlowsCmd = &cobra.Command{
	Use:   "describe-job-flows",
	Short: "This API is no longer supported and will eventually be removed.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_describeJobFlowsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(emr_describeJobFlowsCmd).Standalone()

		emr_describeJobFlowsCmd.Flags().String("created-after", "", "Return only job flows created after this date and time.")
		emr_describeJobFlowsCmd.Flags().String("created-before", "", "Return only job flows created before this date and time.")
		emr_describeJobFlowsCmd.Flags().String("job-flow-ids", "", "Return only job flows whose job flow ID is contained in this list.")
		emr_describeJobFlowsCmd.Flags().String("job-flow-states", "", "Return only job flows whose state is contained in this list.")
	})
	emrCmd.AddCommand(emr_describeJobFlowsCmd)
}
