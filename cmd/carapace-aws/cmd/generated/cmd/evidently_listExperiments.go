package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var evidently_listExperimentsCmd = &cobra.Command{
	Use:   "list-experiments",
	Short: "Returns configuration details about all the experiments in the specified project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(evidently_listExperimentsCmd).Standalone()

	evidently_listExperimentsCmd.Flags().String("max-results", "", "The maximum number of results to include in the response.")
	evidently_listExperimentsCmd.Flags().String("next-token", "", "The token to use when requesting the next set of results.")
	evidently_listExperimentsCmd.Flags().String("project", "", "The name or ARN of the project to return the experiment list from.")
	evidently_listExperimentsCmd.Flags().String("status", "", "Use this optional parameter to limit the returned results to only the experiments with the status that you specify here.")
	evidently_listExperimentsCmd.MarkFlagRequired("project")
	evidentlyCmd.AddCommand(evidently_listExperimentsCmd)
}
