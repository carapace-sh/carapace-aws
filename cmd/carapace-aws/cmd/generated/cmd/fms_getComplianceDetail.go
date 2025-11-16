package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fms_getComplianceDetailCmd = &cobra.Command{
	Use:   "get-compliance-detail",
	Short: "Returns detailed compliance information about the specified member account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fms_getComplianceDetailCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fms_getComplianceDetailCmd).Standalone()

		fms_getComplianceDetailCmd.Flags().String("member-account", "", "The Amazon Web Services account that owns the resources that you want to get the details for.")
		fms_getComplianceDetailCmd.Flags().String("policy-id", "", "The ID of the policy that you want to get the details for.")
		fms_getComplianceDetailCmd.MarkFlagRequired("member-account")
		fms_getComplianceDetailCmd.MarkFlagRequired("policy-id")
	})
	fmsCmd.AddCommand(fms_getComplianceDetailCmd)
}
