package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_createBotReplicaCmd = &cobra.Command{
	Use:   "create-bot-replica",
	Short: "Action to create a replication of the source bot in the secondary region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_createBotReplicaCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexv2Models_createBotReplicaCmd).Standalone()

		lexv2Models_createBotReplicaCmd.Flags().String("bot-id", "", "The request for the unique bot ID of the source bot to be replicated in the secondary region.")
		lexv2Models_createBotReplicaCmd.Flags().String("replica-region", "", "The request for the secondary region that will be used in the replication of the source bot.")
		lexv2Models_createBotReplicaCmd.MarkFlagRequired("bot-id")
		lexv2Models_createBotReplicaCmd.MarkFlagRequired("replica-region")
	})
	lexv2ModelsCmd.AddCommand(lexv2Models_createBotReplicaCmd)
}
