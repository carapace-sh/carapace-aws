package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_describeBotReplicaCmd = &cobra.Command{
	Use:   "describe-bot-replica",
	Short: "Monitors the bot replication status through the UI console.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_describeBotReplicaCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexv2Models_describeBotReplicaCmd).Standalone()

		lexv2Models_describeBotReplicaCmd.Flags().String("bot-id", "", "The request for the unique bot ID of the replicated bot being monitored.")
		lexv2Models_describeBotReplicaCmd.Flags().String("replica-region", "", "The request for the region of the replicated bot being monitored.")
		lexv2Models_describeBotReplicaCmd.MarkFlagRequired("bot-id")
		lexv2Models_describeBotReplicaCmd.MarkFlagRequired("replica-region")
	})
	lexv2ModelsCmd.AddCommand(lexv2Models_describeBotReplicaCmd)
}
