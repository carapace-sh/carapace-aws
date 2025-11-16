package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_deleteResourcePolicyCmd = &cobra.Command{
	Use:   "delete-resource-policy",
	Short: "Removes an existing policy from a bot or bot alias.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_deleteResourcePolicyCmd).Standalone()

	lexv2Models_deleteResourcePolicyCmd.Flags().String("expected-revision-id", "", "The identifier of the revision to edit.")
	lexv2Models_deleteResourcePolicyCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the bot or bot alias that has the resource policy attached.")
	lexv2Models_deleteResourcePolicyCmd.MarkFlagRequired("resource-arn")
	lexv2ModelsCmd.AddCommand(lexv2Models_deleteResourcePolicyCmd)
}
