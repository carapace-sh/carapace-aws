package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_describeInstancePropertiesCmd = &cobra.Command{
	Use:   "describe-instance-properties",
	Short: "An API operation used by the Systems Manager console to display information about Systems Manager managed nodes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_describeInstancePropertiesCmd).Standalone()

	ssm_describeInstancePropertiesCmd.Flags().String("filters-with-operator", "", "The request filters to use with the operator.")
	ssm_describeInstancePropertiesCmd.Flags().String("instance-property-filter-list", "", "An array of instance property filters.")
	ssm_describeInstancePropertiesCmd.Flags().String("max-results", "", "The maximum number of items to return for the call.")
	ssm_describeInstancePropertiesCmd.Flags().String("next-token", "", "The token provided by a previous request to use to return the next set of properties.")
	ssmCmd.AddCommand(ssm_describeInstancePropertiesCmd)
}
