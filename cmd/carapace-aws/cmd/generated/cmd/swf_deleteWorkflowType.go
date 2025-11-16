package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var swf_deleteWorkflowTypeCmd = &cobra.Command{
	Use:   "delete-workflow-type",
	Short: "Deletes the specified *workflow type*.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swf_deleteWorkflowTypeCmd).Standalone()

	swf_deleteWorkflowTypeCmd.Flags().String("domain", "", "The name of the domain in which the workflow type is registered.")
	swf_deleteWorkflowTypeCmd.Flags().String("workflow-type", "", "The workflow type to delete.")
	swf_deleteWorkflowTypeCmd.MarkFlagRequired("domain")
	swf_deleteWorkflowTypeCmd.MarkFlagRequired("workflow-type")
	swfCmd.AddCommand(swf_deleteWorkflowTypeCmd)
}
