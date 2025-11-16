package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var swf_undeprecateWorkflowTypeCmd = &cobra.Command{
	Use:   "undeprecate-workflow-type",
	Short: "Undeprecates a previously deprecated *workflow type*.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swf_undeprecateWorkflowTypeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(swf_undeprecateWorkflowTypeCmd).Standalone()

		swf_undeprecateWorkflowTypeCmd.Flags().String("domain", "", "The name of the domain of the deprecated workflow type.")
		swf_undeprecateWorkflowTypeCmd.Flags().String("workflow-type", "", "The name of the domain of the deprecated workflow type.")
		swf_undeprecateWorkflowTypeCmd.MarkFlagRequired("domain")
		swf_undeprecateWorkflowTypeCmd.MarkFlagRequired("workflow-type")
	})
	swfCmd.AddCommand(swf_undeprecateWorkflowTypeCmd)
}
