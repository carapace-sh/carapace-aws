package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_describeStackResourceDriftsCmd = &cobra.Command{
	Use:   "describe-stack-resource-drifts",
	Short: "Returns drift information for the resources that have been checked for drift in the specified stack.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_describeStackResourceDriftsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudformation_describeStackResourceDriftsCmd).Standalone()

		cloudformation_describeStackResourceDriftsCmd.Flags().String("max-results", "", "The maximum number of results to be returned with a single call.")
		cloudformation_describeStackResourceDriftsCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
		cloudformation_describeStackResourceDriftsCmd.Flags().String("stack-name", "", "The name of the stack for which you want drift information.")
		cloudformation_describeStackResourceDriftsCmd.Flags().String("stack-resource-drift-status-filters", "", "The resource drift status values to use as filters for the resource drift results returned.")
		cloudformation_describeStackResourceDriftsCmd.MarkFlagRequired("stack-name")
	})
	cloudformationCmd.AddCommand(cloudformation_describeStackResourceDriftsCmd)
}
