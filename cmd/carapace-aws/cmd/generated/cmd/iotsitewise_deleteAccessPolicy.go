package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_deleteAccessPolicyCmd = &cobra.Command{
	Use:   "delete-access-policy",
	Short: "Deletes an access policy that grants the specified identity access to the specified IoT SiteWise Monitor resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_deleteAccessPolicyCmd).Standalone()

	iotsitewise_deleteAccessPolicyCmd.Flags().String("access-policy-id", "", "The ID of the access policy to be deleted.")
	iotsitewise_deleteAccessPolicyCmd.Flags().String("client-token", "", "A unique case-sensitive identifier that you can provide to ensure the idempotency of the request.")
	iotsitewise_deleteAccessPolicyCmd.MarkFlagRequired("access-policy-id")
	iotsitewiseCmd.AddCommand(iotsitewise_deleteAccessPolicyCmd)
}
