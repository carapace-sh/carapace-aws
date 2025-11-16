package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lookoutequipment_deleteResourcePolicyCmd = &cobra.Command{
	Use:   "delete-resource-policy",
	Short: "Deletes the resource policy attached to the resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lookoutequipment_deleteResourcePolicyCmd).Standalone()

	lookoutequipment_deleteResourcePolicyCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource for which the resource policy should be deleted.")
	lookoutequipment_deleteResourcePolicyCmd.MarkFlagRequired("resource-arn")
	lookoutequipmentCmd.AddCommand(lookoutequipment_deleteResourcePolicyCmd)
}
