package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_listLifecycleExecutionResourcesCmd = &cobra.Command{
	Use:   "list-lifecycle-execution-resources",
	Short: "List resources that the runtime instance of the image lifecycle identified for lifecycle actions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_listLifecycleExecutionResourcesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(imagebuilder_listLifecycleExecutionResourcesCmd).Standalone()

		imagebuilder_listLifecycleExecutionResourcesCmd.Flags().String("lifecycle-execution-id", "", "Use the unique identifier for a runtime instance of the lifecycle policy to get runtime details.")
		imagebuilder_listLifecycleExecutionResourcesCmd.Flags().String("max-results", "", "Specify the maximum number of items to return in a request.")
		imagebuilder_listLifecycleExecutionResourcesCmd.Flags().String("next-token", "", "A token to specify where to start paginating.")
		imagebuilder_listLifecycleExecutionResourcesCmd.Flags().String("parent-resource-id", "", "You can leave this empty to get a list of Image Builder resources that were identified for lifecycle actions.")
		imagebuilder_listLifecycleExecutionResourcesCmd.MarkFlagRequired("lifecycle-execution-id")
	})
	imagebuilderCmd.AddCommand(imagebuilder_listLifecycleExecutionResourcesCmd)
}
