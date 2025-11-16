package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_updateAccessPolicyCmd = &cobra.Command{
	Use:   "update-access-policy",
	Short: "Updates an existing access policy that specifies an identity's access to an IoT SiteWise Monitor portal or project resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_updateAccessPolicyCmd).Standalone()

	iotsitewise_updateAccessPolicyCmd.Flags().String("access-policy-id", "", "The ID of the access policy.")
	iotsitewise_updateAccessPolicyCmd.Flags().String("access-policy-identity", "", "The identity for this access policy.")
	iotsitewise_updateAccessPolicyCmd.Flags().String("access-policy-permission", "", "The permission level for this access policy.")
	iotsitewise_updateAccessPolicyCmd.Flags().String("access-policy-resource", "", "The IoT SiteWise Monitor resource for this access policy.")
	iotsitewise_updateAccessPolicyCmd.Flags().String("client-token", "", "A unique case-sensitive identifier that you can provide to ensure the idempotency of the request.")
	iotsitewise_updateAccessPolicyCmd.MarkFlagRequired("access-policy-id")
	iotsitewise_updateAccessPolicyCmd.MarkFlagRequired("access-policy-identity")
	iotsitewise_updateAccessPolicyCmd.MarkFlagRequired("access-policy-permission")
	iotsitewise_updateAccessPolicyCmd.MarkFlagRequired("access-policy-resource")
	iotsitewiseCmd.AddCommand(iotsitewise_updateAccessPolicyCmd)
}
