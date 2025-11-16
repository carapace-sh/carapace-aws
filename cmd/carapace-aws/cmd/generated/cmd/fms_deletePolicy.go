package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fms_deletePolicyCmd = &cobra.Command{
	Use:   "delete-policy",
	Short: "Permanently deletes an Firewall Manager policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fms_deletePolicyCmd).Standalone()

	fms_deletePolicyCmd.Flags().Bool("delete-all-policy-resources", false, "If `True`, the request performs cleanup according to the policy type.")
	fms_deletePolicyCmd.Flags().Bool("no-delete-all-policy-resources", false, "If `True`, the request performs cleanup according to the policy type.")
	fms_deletePolicyCmd.Flags().String("policy-id", "", "The ID of the policy that you want to delete.")
	fms_deletePolicyCmd.Flag("no-delete-all-policy-resources").Hidden = true
	fms_deletePolicyCmd.MarkFlagRequired("policy-id")
	fmsCmd.AddCommand(fms_deletePolicyCmd)
}
