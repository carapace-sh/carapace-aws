package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_getJobCmd = &cobra.Command{
	Use:   "get-job",
	Short: "Gets a Deadline Cloud job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_getJobCmd).Standalone()

	deadline_getJobCmd.Flags().String("farm-id", "", "The farm ID of the farm in the job.")
	deadline_getJobCmd.Flags().String("job-id", "", "The job ID.")
	deadline_getJobCmd.Flags().String("queue-id", "", "The queue ID associated with the job.")
	deadline_getJobCmd.MarkFlagRequired("farm-id")
	deadline_getJobCmd.MarkFlagRequired("job-id")
	deadline_getJobCmd.MarkFlagRequired("queue-id")
	deadlineCmd.AddCommand(deadline_getJobCmd)
}
