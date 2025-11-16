package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listPolicyPrincipalsCmd = &cobra.Command{
	Use:   "list-policy-principals",
	Short: "Lists the principals associated with the specified policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listPolicyPrincipalsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_listPolicyPrincipalsCmd).Standalone()

		iot_listPolicyPrincipalsCmd.Flags().String("ascending-order", "", "Specifies the order for results.")
		iot_listPolicyPrincipalsCmd.Flags().String("marker", "", "The marker for the next set of results.")
		iot_listPolicyPrincipalsCmd.Flags().String("page-size", "", "The result page size.")
		iot_listPolicyPrincipalsCmd.Flags().String("policy-name", "", "The policy name.")
		iot_listPolicyPrincipalsCmd.MarkFlagRequired("policy-name")
	})
	iotCmd.AddCommand(iot_listPolicyPrincipalsCmd)
}
