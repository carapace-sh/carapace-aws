package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_putResourcePolicyCmd = &cobra.Command{
	Use:   "put-resource-policy",
	Short: "Creates or updates a resource policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_putResourcePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkmanager_putResourcePolicyCmd).Standalone()

		networkmanager_putResourcePolicyCmd.Flags().String("policy-document", "", "The JSON resource policy document.")
		networkmanager_putResourcePolicyCmd.Flags().String("resource-arn", "", "The ARN of the resource policy.")
		networkmanager_putResourcePolicyCmd.MarkFlagRequired("policy-document")
		networkmanager_putResourcePolicyCmd.MarkFlagRequired("resource-arn")
	})
	networkmanagerCmd.AddCommand(networkmanager_putResourcePolicyCmd)
}
