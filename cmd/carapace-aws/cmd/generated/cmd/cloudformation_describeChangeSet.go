package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_describeChangeSetCmd = &cobra.Command{
	Use:   "describe-change-set",
	Short: "Returns the inputs for the change set and a list of changes that CloudFormation will make if you execute the change set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_describeChangeSetCmd).Standalone()

	cloudformation_describeChangeSetCmd.Flags().String("change-set-name", "", "The name or Amazon Resource Name (ARN) of the change set that you want to describe.")
	cloudformation_describeChangeSetCmd.Flags().String("include-property-values", "", "If `true`, the returned changes include detailed changes in the property values.")
	cloudformation_describeChangeSetCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
	cloudformation_describeChangeSetCmd.Flags().String("stack-name", "", "If you specified the name of a change set, specify the stack name or ID (ARN) of the change set you want to describe.")
	cloudformation_describeChangeSetCmd.MarkFlagRequired("change-set-name")
	cloudformationCmd.AddCommand(cloudformation_describeChangeSetCmd)
}
