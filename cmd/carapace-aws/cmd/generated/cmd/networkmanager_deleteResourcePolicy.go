package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_deleteResourcePolicyCmd = &cobra.Command{
	Use:   "delete-resource-policy",
	Short: "Deletes a resource policy for the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_deleteResourcePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkmanager_deleteResourcePolicyCmd).Standalone()

		networkmanager_deleteResourcePolicyCmd.Flags().String("resource-arn", "", "The ARN of the policy to delete.")
		networkmanager_deleteResourcePolicyCmd.MarkFlagRequired("resource-arn")
	})
	networkmanagerCmd.AddCommand(networkmanager_deleteResourcePolicyCmd)
}
