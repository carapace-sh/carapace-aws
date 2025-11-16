package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elbv2_getResourcePolicyCmd = &cobra.Command{
	Use:   "get-resource-policy",
	Short: "Retrieves the resource policy for a specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elbv2_getResourcePolicyCmd).Standalone()

	elbv2_getResourcePolicyCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
	elbv2_getResourcePolicyCmd.MarkFlagRequired("resource-arn")
	elbv2Cmd.AddCommand(elbv2_getResourcePolicyCmd)
}
