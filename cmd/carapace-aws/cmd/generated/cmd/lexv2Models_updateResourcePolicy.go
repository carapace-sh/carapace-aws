package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_updateResourcePolicyCmd = &cobra.Command{
	Use:   "update-resource-policy",
	Short: "Replaces the existing resource policy for a bot or bot alias with a new one.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_updateResourcePolicyCmd).Standalone()

	lexv2Models_updateResourcePolicyCmd.Flags().String("expected-revision-id", "", "The identifier of the revision of the policy to update.")
	lexv2Models_updateResourcePolicyCmd.Flags().String("policy", "", "A resource policy to add to the resource.")
	lexv2Models_updateResourcePolicyCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the bot or bot alias that the resource policy is attached to.")
	lexv2Models_updateResourcePolicyCmd.MarkFlagRequired("policy")
	lexv2Models_updateResourcePolicyCmd.MarkFlagRequired("resource-arn")
	lexv2ModelsCmd.AddCommand(lexv2Models_updateResourcePolicyCmd)
}
