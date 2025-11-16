package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_getResourcePolicyCmd = &cobra.Command{
	Use:   "get-resource-policy",
	Short: "Returns information about a resource policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_getResourcePolicyCmd).Standalone()

	networkmanager_getResourcePolicyCmd.Flags().String("resource-arn", "", "The ARN of the resource.")
	networkmanager_getResourcePolicyCmd.MarkFlagRequired("resource-arn")
	networkmanagerCmd.AddCommand(networkmanager_getResourcePolicyCmd)
}
