package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ram_getResourcePoliciesCmd = &cobra.Command{
	Use:   "get-resource-policies",
	Short: "Retrieves the resource policies for the specified resources that you own and have shared.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ram_getResourcePoliciesCmd).Standalone()

	ram_getResourcePoliciesCmd.Flags().String("max-results", "", "Specifies the total number of results that you want included on each page of the response.")
	ram_getResourcePoliciesCmd.Flags().String("next-token", "", "Specifies that you want to receive the next page of results.")
	ram_getResourcePoliciesCmd.Flags().String("principal", "", "Specifies the principal.")
	ram_getResourcePoliciesCmd.Flags().String("resource-arns", "", "Specifies the [Amazon Resource Names (ARNs)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the resources whose policies you want to retrieve.")
	ram_getResourcePoliciesCmd.MarkFlagRequired("resource-arns")
	ramCmd.AddCommand(ram_getResourcePoliciesCmd)
}
