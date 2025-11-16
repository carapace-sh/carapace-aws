package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var swf_deprecateWorkflowTypeCmd = &cobra.Command{
	Use:   "deprecate-workflow-type",
	Short: "Deprecates the specified *workflow type*.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swf_deprecateWorkflowTypeCmd).Standalone()

	swf_deprecateWorkflowTypeCmd.Flags().String("domain", "", "The name of the domain in which the workflow type is registered.")
	swf_deprecateWorkflowTypeCmd.Flags().String("workflow-type", "", "The workflow type to deprecate.")
	swf_deprecateWorkflowTypeCmd.MarkFlagRequired("domain")
	swf_deprecateWorkflowTypeCmd.MarkFlagRequired("workflow-type")
	swfCmd.AddCommand(swf_deprecateWorkflowTypeCmd)
}
