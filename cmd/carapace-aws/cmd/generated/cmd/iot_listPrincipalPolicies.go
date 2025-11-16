package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listPrincipalPoliciesCmd = &cobra.Command{
	Use:   "list-principal-policies",
	Short: "Lists the policies attached to the specified principal.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listPrincipalPoliciesCmd).Standalone()

	iot_listPrincipalPoliciesCmd.Flags().String("ascending-order", "", "Specifies the order for results.")
	iot_listPrincipalPoliciesCmd.Flags().String("marker", "", "The marker for the next set of results.")
	iot_listPrincipalPoliciesCmd.Flags().String("page-size", "", "The result page size.")
	iot_listPrincipalPoliciesCmd.Flags().String("principal", "", "The principal.")
	iot_listPrincipalPoliciesCmd.MarkFlagRequired("principal")
	iotCmd.AddCommand(iot_listPrincipalPoliciesCmd)
}
