package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_listWorkflowsCmd = &cobra.Command{
	Use:   "list-workflows",
	Short: "Lists workflow build versions based on filtering parameters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_listWorkflowsCmd).Standalone()

	imagebuilder_listWorkflowsCmd.Flags().Bool("by-name", false, "Specify all or part of the workflow name to streamline results.")
	imagebuilder_listWorkflowsCmd.Flags().String("filters", "", "Used to streamline search results.")
	imagebuilder_listWorkflowsCmd.Flags().String("max-results", "", "Specify the maximum number of items to return in a request.")
	imagebuilder_listWorkflowsCmd.Flags().String("next-token", "", "A token to specify where to start paginating.")
	imagebuilder_listWorkflowsCmd.Flags().Bool("no-by-name", false, "Specify all or part of the workflow name to streamline results.")
	imagebuilder_listWorkflowsCmd.Flags().String("owner", "", "Used to get a list of workflow build version filtered by the identity of the creator.")
	imagebuilder_listWorkflowsCmd.Flag("no-by-name").Hidden = true
	imagebuilderCmd.AddCommand(imagebuilder_listWorkflowsCmd)
}
