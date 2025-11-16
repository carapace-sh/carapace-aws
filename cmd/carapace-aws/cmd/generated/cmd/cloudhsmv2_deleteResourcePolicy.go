package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudhsmv2_deleteResourcePolicyCmd = &cobra.Command{
	Use:   "delete-resource-policy",
	Short: "Deletes an CloudHSM resource policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudhsmv2_deleteResourcePolicyCmd).Standalone()

	cloudhsmv2_deleteResourcePolicyCmd.Flags().String("resource-arn", "", "Amazon Resource Name (ARN) of the resource from which the policy will be removed.")
	cloudhsmv2Cmd.AddCommand(cloudhsmv2_deleteResourcePolicyCmd)
}
