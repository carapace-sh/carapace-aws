package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_listJobParameterDefinitionsCmd = &cobra.Command{
	Use:   "list-job-parameter-definitions",
	Short: "Lists parameter definitions of a job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_listJobParameterDefinitionsCmd).Standalone()

	deadline_listJobParameterDefinitionsCmd.Flags().String("farm-id", "", "The farm ID of the job to list.")
	deadline_listJobParameterDefinitionsCmd.Flags().String("job-id", "", "The job ID to include on the list.")
	deadline_listJobParameterDefinitionsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
	deadline_listJobParameterDefinitionsCmd.Flags().String("next-token", "", "The token for the next set of results, or `null` to start from the beginning.")
	deadline_listJobParameterDefinitionsCmd.Flags().String("queue-id", "", "The queue ID to include on the list.")
	deadline_listJobParameterDefinitionsCmd.MarkFlagRequired("farm-id")
	deadline_listJobParameterDefinitionsCmd.MarkFlagRequired("job-id")
	deadline_listJobParameterDefinitionsCmd.MarkFlagRequired("queue-id")
	deadlineCmd.AddCommand(deadline_listJobParameterDefinitionsCmd)
}
