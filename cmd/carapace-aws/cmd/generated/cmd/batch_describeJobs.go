package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var batch_describeJobsCmd = &cobra.Command{
	Use:   "describe-jobs",
	Short: "Describes a list of Batch jobs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(batch_describeJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(batch_describeJobsCmd).Standalone()

		batch_describeJobsCmd.Flags().String("jobs", "", "A list of up to 100 job IDs.")
		batch_describeJobsCmd.MarkFlagRequired("jobs")
	})
	batchCmd.AddCommand(batch_describeJobsCmd)
}
