package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_deleteBotReplicaCmd = &cobra.Command{
	Use:   "delete-bot-replica",
	Short: "The action to delete the replicated bot in the secondary region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_deleteBotReplicaCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexv2Models_deleteBotReplicaCmd).Standalone()

		lexv2Models_deleteBotReplicaCmd.Flags().String("bot-id", "", "The unique ID of the replicated bot to be deleted from the secondary region")
		lexv2Models_deleteBotReplicaCmd.Flags().String("replica-region", "", "The secondary region of the replicated bot that will be deleted.")
		lexv2Models_deleteBotReplicaCmd.MarkFlagRequired("bot-id")
		lexv2Models_deleteBotReplicaCmd.MarkFlagRequired("replica-region")
	})
	lexv2ModelsCmd.AddCommand(lexv2Models_deleteBotReplicaCmd)
}
