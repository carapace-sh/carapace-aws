package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecatalyst_listWorkflowsCmd = &cobra.Command{
	Use:   "list-workflows",
	Short: "Retrieves a list of workflows in a specified project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecatalyst_listWorkflowsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecatalyst_listWorkflowsCmd).Standalone()

		codecatalyst_listWorkflowsCmd.Flags().String("max-results", "", "The maximum number of results to show in a single call to this API.")
		codecatalyst_listWorkflowsCmd.Flags().String("next-token", "", "A token returned from a call to this API to indicate the next batch of results to return, if any.")
		codecatalyst_listWorkflowsCmd.Flags().String("project-name", "", "The name of the project in the space.")
		codecatalyst_listWorkflowsCmd.Flags().String("sort-by", "", "Information used to sort the items in the returned list.")
		codecatalyst_listWorkflowsCmd.Flags().String("space-name", "", "The name of the space.")
		codecatalyst_listWorkflowsCmd.MarkFlagRequired("project-name")
		codecatalyst_listWorkflowsCmd.MarkFlagRequired("space-name")
	})
	codecatalystCmd.AddCommand(codecatalyst_listWorkflowsCmd)
}
