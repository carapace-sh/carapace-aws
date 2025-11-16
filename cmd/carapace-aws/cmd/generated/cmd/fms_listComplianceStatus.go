package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fms_listComplianceStatusCmd = &cobra.Command{
	Use:   "list-compliance-status",
	Short: "Returns an array of `PolicyComplianceStatus` objects.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fms_listComplianceStatusCmd).Standalone()

	fms_listComplianceStatusCmd.Flags().String("max-results", "", "Specifies the number of `PolicyComplianceStatus` objects that you want Firewall Manager to return for this request.")
	fms_listComplianceStatusCmd.Flags().String("next-token", "", "If you specify a value for `MaxResults` and you have more `PolicyComplianceStatus` objects than the number that you specify for `MaxResults`, Firewall Manager returns a `NextToken` value in the response that allows you to list another group of `PolicyComplianceStatus` objects.")
	fms_listComplianceStatusCmd.Flags().String("policy-id", "", "The ID of the Firewall Manager policy that you want the details for.")
	fms_listComplianceStatusCmd.MarkFlagRequired("policy-id")
	fmsCmd.AddCommand(fms_listComplianceStatusCmd)
}
