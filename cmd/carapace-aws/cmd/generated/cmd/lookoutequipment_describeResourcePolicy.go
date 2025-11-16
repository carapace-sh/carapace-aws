package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lookoutequipment_describeResourcePolicyCmd = &cobra.Command{
	Use:   "describe-resource-policy",
	Short: "Provides the details of a resource policy attached to a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lookoutequipment_describeResourcePolicyCmd).Standalone()

	lookoutequipment_describeResourcePolicyCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource that is associated with the resource policy.")
	lookoutequipment_describeResourcePolicyCmd.MarkFlagRequired("resource-arn")
	lookoutequipmentCmd.AddCommand(lookoutequipment_describeResourcePolicyCmd)
}
