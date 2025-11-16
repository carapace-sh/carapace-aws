package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fms_getProtectionStatusCmd = &cobra.Command{
	Use:   "get-protection-status",
	Short: "If you created a Shield Advanced policy, returns policy-level attack summary information in the event of a potential DDoS attack.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fms_getProtectionStatusCmd).Standalone()

	fms_getProtectionStatusCmd.Flags().String("end-time", "", "The end of the time period to query for the attacks.")
	fms_getProtectionStatusCmd.Flags().String("max-results", "", "Specifies the number of objects that you want Firewall Manager to return for this request.")
	fms_getProtectionStatusCmd.Flags().String("member-account-id", "", "The Amazon Web Services account that is in scope of the policy that you want to get the details for.")
	fms_getProtectionStatusCmd.Flags().String("next-token", "", "If you specify a value for `MaxResults` and you have more objects than the number that you specify for `MaxResults`, Firewall Manager returns a `NextToken` value in the response, which you can use to retrieve another group of objects.")
	fms_getProtectionStatusCmd.Flags().String("policy-id", "", "The ID of the policy for which you want to get the attack information.")
	fms_getProtectionStatusCmd.Flags().String("start-time", "", "The start of the time period to query for the attacks.")
	fms_getProtectionStatusCmd.MarkFlagRequired("policy-id")
	fmsCmd.AddCommand(fms_getProtectionStatusCmd)
}
