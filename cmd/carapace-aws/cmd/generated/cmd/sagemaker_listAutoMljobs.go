package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listAutoMljobsCmd = &cobra.Command{
	Use:   "list-auto-mljobs",
	Short: "Request a list of jobs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listAutoMljobsCmd).Standalone()

	sagemaker_listAutoMljobsCmd.Flags().String("creation-time-after", "", "Request a list of jobs, using a filter for time.")
	sagemaker_listAutoMljobsCmd.Flags().String("creation-time-before", "", "Request a list of jobs, using a filter for time.")
	sagemaker_listAutoMljobsCmd.Flags().String("last-modified-time-after", "", "Request a list of jobs, using a filter for time.")
	sagemaker_listAutoMljobsCmd.Flags().String("last-modified-time-before", "", "Request a list of jobs, using a filter for time.")
	sagemaker_listAutoMljobsCmd.Flags().String("max-results", "", "Request a list of jobs up to a specified limit.")
	sagemaker_listAutoMljobsCmd.Flags().String("name-contains", "", "Request a list of jobs, using a search filter for name.")
	sagemaker_listAutoMljobsCmd.Flags().String("next-token", "", "If the previous response was truncated, you receive this token.")
	sagemaker_listAutoMljobsCmd.Flags().String("sort-by", "", "The parameter by which to sort the results.")
	sagemaker_listAutoMljobsCmd.Flags().String("sort-order", "", "The sort order for the results.")
	sagemaker_listAutoMljobsCmd.Flags().String("status-equals", "", "Request a list of jobs, using a filter for status.")
	sagemakerCmd.AddCommand(sagemaker_listAutoMljobsCmd)
}
