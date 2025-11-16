package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var swf_countOpenWorkflowExecutionsCmd = &cobra.Command{
	Use:   "count-open-workflow-executions",
	Short: "Returns the number of open workflow executions within the given domain that meet the specified filtering criteria.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swf_countOpenWorkflowExecutionsCmd).Standalone()

	swf_countOpenWorkflowExecutionsCmd.Flags().String("domain", "", "The name of the domain containing the workflow executions to count.")
	swf_countOpenWorkflowExecutionsCmd.Flags().String("execution-filter", "", "If specified, only workflow executions matching the `WorkflowId` in the filter are counted.")
	swf_countOpenWorkflowExecutionsCmd.Flags().String("start-time-filter", "", "Specifies the start time criteria that workflow executions must meet in order to be counted.")
	swf_countOpenWorkflowExecutionsCmd.Flags().String("tag-filter", "", "If specified, only executions that have a tag that matches the filter are counted.")
	swf_countOpenWorkflowExecutionsCmd.Flags().String("type-filter", "", "Specifies the type of the workflow executions to be counted.")
	swf_countOpenWorkflowExecutionsCmd.MarkFlagRequired("domain")
	swf_countOpenWorkflowExecutionsCmd.MarkFlagRequired("start-time-filter")
	swfCmd.AddCommand(swf_countOpenWorkflowExecutionsCmd)
}
