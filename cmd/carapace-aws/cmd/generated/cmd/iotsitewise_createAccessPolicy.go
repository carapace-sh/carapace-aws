package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_createAccessPolicyCmd = &cobra.Command{
	Use:   "create-access-policy",
	Short: "Creates an access policy that grants the specified identity (IAM Identity Center user, IAM Identity Center group, or IAM user) access to the specified IoT SiteWise Monitor portal or project resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_createAccessPolicyCmd).Standalone()

	iotsitewise_createAccessPolicyCmd.Flags().String("access-policy-identity", "", "The identity for this access policy.")
	iotsitewise_createAccessPolicyCmd.Flags().String("access-policy-permission", "", "The permission level for this access policy.")
	iotsitewise_createAccessPolicyCmd.Flags().String("access-policy-resource", "", "The IoT SiteWise Monitor resource for this access policy.")
	iotsitewise_createAccessPolicyCmd.Flags().String("client-token", "", "A unique case-sensitive identifier that you can provide to ensure the idempotency of the request.")
	iotsitewise_createAccessPolicyCmd.Flags().String("tags", "", "A list of key-value pairs that contain metadata for the access policy.")
	iotsitewise_createAccessPolicyCmd.MarkFlagRequired("access-policy-identity")
	iotsitewise_createAccessPolicyCmd.MarkFlagRequired("access-policy-permission")
	iotsitewise_createAccessPolicyCmd.MarkFlagRequired("access-policy-resource")
	iotsitewiseCmd.AddCommand(iotsitewise_createAccessPolicyCmd)
}
