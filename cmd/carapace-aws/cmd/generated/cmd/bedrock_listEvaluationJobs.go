package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_listEvaluationJobsCmd = &cobra.Command{
	Use:   "list-evaluation-jobs",
	Short: "Lists all existing evaluation jobs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_listEvaluationJobsCmd).Standalone()

	bedrock_listEvaluationJobsCmd.Flags().String("application-type-equals", "", "A filter to only list evaluation jobs that are either model evaluations or knowledge base evaluations.")
	bedrock_listEvaluationJobsCmd.Flags().String("creation-time-after", "", "A filter to only list evaluation jobs created after a specified time.")
	bedrock_listEvaluationJobsCmd.Flags().String("creation-time-before", "", "A filter to only list evaluation jobs created before a specified time.")
	bedrock_listEvaluationJobsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
	bedrock_listEvaluationJobsCmd.Flags().String("name-contains", "", "A filter to only list evaluation jobs that contain a specified string in the job name.")
	bedrock_listEvaluationJobsCmd.Flags().String("next-token", "", "Continuation token from the previous response, for Amazon Bedrock to list the next set of results.")
	bedrock_listEvaluationJobsCmd.Flags().String("sort-by", "", "Specifies a creation time to sort the list of evaluation jobs by when they were created.")
	bedrock_listEvaluationJobsCmd.Flags().String("sort-order", "", "Specifies whether to sort the list of evaluation jobs by either ascending or descending order.")
	bedrock_listEvaluationJobsCmd.Flags().String("status-equals", "", "A filter to only list evaluation jobs that are of a certain status.")
	bedrockCmd.AddCommand(bedrock_listEvaluationJobsCmd)
}
