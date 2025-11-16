package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fms_listPoliciesCmd = &cobra.Command{
	Use:   "list-policies",
	Short: "Returns an array of `PolicySummary` objects.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fms_listPoliciesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fms_listPoliciesCmd).Standalone()

		fms_listPoliciesCmd.Flags().String("max-results", "", "Specifies the number of `PolicySummary` objects that you want Firewall Manager to return for this request.")
		fms_listPoliciesCmd.Flags().String("next-token", "", "If you specify a value for `MaxResults` and you have more `PolicySummary` objects than the number that you specify for `MaxResults`, Firewall Manager returns a `NextToken` value in the response that allows you to list another group of `PolicySummary` objects.")
	})
	fmsCmd.AddCommand(fms_listPoliciesCmd)
}
