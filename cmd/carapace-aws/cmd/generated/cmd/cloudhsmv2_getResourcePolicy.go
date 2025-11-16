package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudhsmv2_getResourcePolicyCmd = &cobra.Command{
	Use:   "get-resource-policy",
	Short: "Retrieves the resource policy document attached to a given resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudhsmv2_getResourcePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudhsmv2_getResourcePolicyCmd).Standalone()

		cloudhsmv2_getResourcePolicyCmd.Flags().String("resource-arn", "", "Amazon Resource Name (ARN) of the resource to which a policy is attached.")
	})
	cloudhsmv2Cmd.AddCommand(cloudhsmv2_getResourcePolicyCmd)
}
