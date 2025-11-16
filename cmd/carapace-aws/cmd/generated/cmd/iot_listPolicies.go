package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listPoliciesCmd = &cobra.Command{
	Use:   "list-policies",
	Short: "Lists your policies.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listPoliciesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_listPoliciesCmd).Standalone()

		iot_listPoliciesCmd.Flags().String("ascending-order", "", "Specifies the order for results.")
		iot_listPoliciesCmd.Flags().String("marker", "", "The marker for the next set of results.")
		iot_listPoliciesCmd.Flags().String("page-size", "", "The result page size.")
	})
	iotCmd.AddCommand(iot_listPoliciesCmd)
}
