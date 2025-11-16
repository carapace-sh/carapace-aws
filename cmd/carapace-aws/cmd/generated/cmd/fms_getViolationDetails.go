package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fms_getViolationDetailsCmd = &cobra.Command{
	Use:   "get-violation-details",
	Short: "Retrieves violations for a resource based on the specified Firewall Manager policy and Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fms_getViolationDetailsCmd).Standalone()

	fms_getViolationDetailsCmd.Flags().String("member-account", "", "The Amazon Web Services account ID that you want the details for.")
	fms_getViolationDetailsCmd.Flags().String("policy-id", "", "The ID of the Firewall Manager policy that you want the details for.")
	fms_getViolationDetailsCmd.Flags().String("resource-id", "", "The ID of the resource that has violations.")
	fms_getViolationDetailsCmd.Flags().String("resource-type", "", "The resource type.")
	fms_getViolationDetailsCmd.MarkFlagRequired("member-account")
	fms_getViolationDetailsCmd.MarkFlagRequired("policy-id")
	fms_getViolationDetailsCmd.MarkFlagRequired("resource-id")
	fms_getViolationDetailsCmd.MarkFlagRequired("resource-type")
	fmsCmd.AddCommand(fms_getViolationDetailsCmd)
}
