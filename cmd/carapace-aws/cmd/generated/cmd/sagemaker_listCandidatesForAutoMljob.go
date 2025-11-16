package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listCandidatesForAutoMljobCmd = &cobra.Command{
	Use:   "list-candidates-for-auto-mljob",
	Short: "List the candidates created for the job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listCandidatesForAutoMljobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_listCandidatesForAutoMljobCmd).Standalone()

		sagemaker_listCandidatesForAutoMljobCmd.Flags().String("auto-mljob-name", "", "List the candidates created for the job by providing the job's name.")
		sagemaker_listCandidatesForAutoMljobCmd.Flags().String("candidate-name-equals", "", "List the candidates for the job and filter by candidate name.")
		sagemaker_listCandidatesForAutoMljobCmd.Flags().String("max-results", "", "List the job's candidates up to a specified limit.")
		sagemaker_listCandidatesForAutoMljobCmd.Flags().String("next-token", "", "If the previous response was truncated, you receive this token.")
		sagemaker_listCandidatesForAutoMljobCmd.Flags().String("sort-by", "", "The parameter by which to sort the results.")
		sagemaker_listCandidatesForAutoMljobCmd.Flags().String("sort-order", "", "The sort order for the results.")
		sagemaker_listCandidatesForAutoMljobCmd.Flags().String("status-equals", "", "List the candidates for the job and filter by status.")
		sagemaker_listCandidatesForAutoMljobCmd.MarkFlagRequired("auto-mljob-name")
	})
	sagemakerCmd.AddCommand(sagemaker_listCandidatesForAutoMljobCmd)
}
