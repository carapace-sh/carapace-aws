package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lookoutequipment_putResourcePolicyCmd = &cobra.Command{
	Use:   "put-resource-policy",
	Short: "Creates a resource control policy for a given resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lookoutequipment_putResourcePolicyCmd).Standalone()

	lookoutequipment_putResourcePolicyCmd.Flags().String("client-token", "", "A unique identifier for the request.")
	lookoutequipment_putResourcePolicyCmd.Flags().String("policy-revision-id", "", "A unique identifier for a revision of the resource policy.")
	lookoutequipment_putResourcePolicyCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource for which the policy is being created.")
	lookoutequipment_putResourcePolicyCmd.Flags().String("resource-policy", "", "The JSON-formatted resource policy to create.")
	lookoutequipment_putResourcePolicyCmd.MarkFlagRequired("client-token")
	lookoutequipment_putResourcePolicyCmd.MarkFlagRequired("resource-arn")
	lookoutequipment_putResourcePolicyCmd.MarkFlagRequired("resource-policy")
	lookoutequipmentCmd.AddCommand(lookoutequipment_putResourcePolicyCmd)
}
