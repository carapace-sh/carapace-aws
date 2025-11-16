package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var swf_listWorkflowTypesCmd = &cobra.Command{
	Use:   "list-workflow-types",
	Short: "Returns information about workflow types in the specified domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swf_listWorkflowTypesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(swf_listWorkflowTypesCmd).Standalone()

		swf_listWorkflowTypesCmd.Flags().String("domain", "", "The name of the domain in which the workflow types have been registered.")
		swf_listWorkflowTypesCmd.Flags().String("maximum-page-size", "", "The maximum number of results that are returned per call.")
		swf_listWorkflowTypesCmd.Flags().String("name", "", "If specified, lists the workflow type with this name.")
		swf_listWorkflowTypesCmd.Flags().String("next-page-token", "", "If `NextPageToken` is returned there are more results available.")
		swf_listWorkflowTypesCmd.Flags().String("registration-status", "", "Specifies the registration status of the workflow types to list.")
		swf_listWorkflowTypesCmd.Flags().String("reverse-order", "", "When set to `true`, returns the results in reverse order.")
		swf_listWorkflowTypesCmd.MarkFlagRequired("domain")
		swf_listWorkflowTypesCmd.MarkFlagRequired("registration-status")
	})
	swfCmd.AddCommand(swf_listWorkflowTypesCmd)
}
