package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_listWorkflowsCmd = &cobra.Command{
	Use:   "list-workflows",
	Short: "Retrieves a list of existing workflows.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_listWorkflowsCmd).Standalone()

	omics_listWorkflowsCmd.Flags().String("max-results", "", "The maximum number of workflows to return in one page of results.")
	omics_listWorkflowsCmd.Flags().String("name", "", "Filter the list by workflow name.")
	omics_listWorkflowsCmd.Flags().String("starting-token", "", "Specify the pagination token from a previous request to retrieve the next page of results.")
	omics_listWorkflowsCmd.Flags().String("type", "", "Filter the list by workflow type.")
	omicsCmd.AddCommand(omics_listWorkflowsCmd)
}
