package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_describeStacksCmd = &cobra.Command{
	Use:   "describe-stacks",
	Short: "Returns the description for the specified stack; if no stack name was specified, then it returns the description for all the stacks created.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_describeStacksCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudformation_describeStacksCmd).Standalone()

		cloudformation_describeStacksCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
		cloudformation_describeStacksCmd.Flags().String("stack-name", "", "If you don't pass a parameter to `StackName`, the API returns a response that describes all resources in the account, which can impact performance.")
	})
	cloudformationCmd.AddCommand(cloudformation_describeStacksCmd)
}
