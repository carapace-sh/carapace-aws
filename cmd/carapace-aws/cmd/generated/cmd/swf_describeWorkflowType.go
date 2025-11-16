package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var swf_describeWorkflowTypeCmd = &cobra.Command{
	Use:   "describe-workflow-type",
	Short: "Returns information about the specified *workflow type*.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swf_describeWorkflowTypeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(swf_describeWorkflowTypeCmd).Standalone()

		swf_describeWorkflowTypeCmd.Flags().String("domain", "", "The name of the domain in which this workflow type is registered.")
		swf_describeWorkflowTypeCmd.Flags().String("workflow-type", "", "The workflow type to describe.")
		swf_describeWorkflowTypeCmd.MarkFlagRequired("domain")
		swf_describeWorkflowTypeCmd.MarkFlagRequired("workflow-type")
	})
	swfCmd.AddCommand(swf_describeWorkflowTypeCmd)
}
