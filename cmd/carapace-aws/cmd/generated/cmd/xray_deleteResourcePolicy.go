package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var xray_deleteResourcePolicyCmd = &cobra.Command{
	Use:   "delete-resource-policy",
	Short: "Deletes a resource policy from the target Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(xray_deleteResourcePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(xray_deleteResourcePolicyCmd).Standalone()

		xray_deleteResourcePolicyCmd.Flags().String("policy-name", "", "The name of the resource policy to delete.")
		xray_deleteResourcePolicyCmd.Flags().String("policy-revision-id", "", "Specifies a specific policy revision to delete.")
		xray_deleteResourcePolicyCmd.MarkFlagRequired("policy-name")
	})
	xrayCmd.AddCommand(xray_deleteResourcePolicyCmd)
}
