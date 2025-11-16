package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudhsmv2_putResourcePolicyCmd = &cobra.Command{
	Use:   "put-resource-policy",
	Short: "Creates or updates an CloudHSM resource policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudhsmv2_putResourcePolicyCmd).Standalone()

	cloudhsmv2_putResourcePolicyCmd.Flags().String("policy", "", "The policy you want to associate with a resource.")
	cloudhsmv2_putResourcePolicyCmd.Flags().String("resource-arn", "", "Amazon Resource Name (ARN) of the resource to which you want to attach a policy.")
	cloudhsmv2Cmd.AddCommand(cloudhsmv2_putResourcePolicyCmd)
}
