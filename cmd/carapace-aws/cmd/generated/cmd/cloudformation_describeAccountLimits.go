package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_describeAccountLimitsCmd = &cobra.Command{
	Use:   "describe-account-limits",
	Short: "Retrieves your account's CloudFormation limits, such as the maximum number of stacks that you can create in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_describeAccountLimitsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudformation_describeAccountLimitsCmd).Standalone()

		cloudformation_describeAccountLimitsCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
	})
	cloudformationCmd.AddCommand(cloudformation_describeAccountLimitsCmd)
}
