package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_describeChangeSetHooksCmd = &cobra.Command{
	Use:   "describe-change-set-hooks",
	Short: "Returns Hook-related information for the change set and a list of changes that CloudFormation makes when you run the change set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_describeChangeSetHooksCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudformation_describeChangeSetHooksCmd).Standalone()

		cloudformation_describeChangeSetHooksCmd.Flags().String("change-set-name", "", "The name or Amazon Resource Name (ARN) of the change set that you want to describe.")
		cloudformation_describeChangeSetHooksCmd.Flags().String("logical-resource-id", "", "If specified, lists only the Hooks related to the specified `LogicalResourceId`.")
		cloudformation_describeChangeSetHooksCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
		cloudformation_describeChangeSetHooksCmd.Flags().String("stack-name", "", "If you specified the name of a change set, specify the stack name or stack ID (ARN) of the change set you want to describe.")
		cloudformation_describeChangeSetHooksCmd.MarkFlagRequired("change-set-name")
	})
	cloudformationCmd.AddCommand(cloudformation_describeChangeSetHooksCmd)
}
