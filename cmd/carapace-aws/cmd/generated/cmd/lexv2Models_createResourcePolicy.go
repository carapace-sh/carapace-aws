package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_createResourcePolicyCmd = &cobra.Command{
	Use:   "create-resource-policy",
	Short: "Creates a new resource policy with the specified policy statements.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_createResourcePolicyCmd).Standalone()

	lexv2Models_createResourcePolicyCmd.Flags().String("policy", "", "A resource policy to add to the resource.")
	lexv2Models_createResourcePolicyCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the bot or bot alias that the resource policy is attached to.")
	lexv2Models_createResourcePolicyCmd.MarkFlagRequired("policy")
	lexv2Models_createResourcePolicyCmd.MarkFlagRequired("resource-arn")
	lexv2ModelsCmd.AddCommand(lexv2Models_createResourcePolicyCmd)
}
