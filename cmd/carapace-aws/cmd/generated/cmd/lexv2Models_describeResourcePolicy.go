package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_describeResourcePolicyCmd = &cobra.Command{
	Use:   "describe-resource-policy",
	Short: "Gets the resource policy and policy revision for a bot or bot alias.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_describeResourcePolicyCmd).Standalone()

	lexv2Models_describeResourcePolicyCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the bot or bot alias that the resource policy is attached to.")
	lexv2Models_describeResourcePolicyCmd.MarkFlagRequired("resource-arn")
	lexv2ModelsCmd.AddCommand(lexv2Models_describeResourcePolicyCmd)
}
