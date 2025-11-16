package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_listWaitingWorkflowStepsCmd = &cobra.Command{
	Use:   "list-waiting-workflow-steps",
	Short: "Get a list of workflow steps that are waiting for action for workflows in your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_listWaitingWorkflowStepsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(imagebuilder_listWaitingWorkflowStepsCmd).Standalone()

		imagebuilder_listWaitingWorkflowStepsCmd.Flags().String("max-results", "", "Specify the maximum number of items to return in a request.")
		imagebuilder_listWaitingWorkflowStepsCmd.Flags().String("next-token", "", "A token to specify where to start paginating.")
	})
	imagebuilderCmd.AddCommand(imagebuilder_listWaitingWorkflowStepsCmd)
}
