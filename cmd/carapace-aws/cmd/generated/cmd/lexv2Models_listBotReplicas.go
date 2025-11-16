package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_listBotReplicasCmd = &cobra.Command{
	Use:   "list-bot-replicas",
	Short: "The action to list the replicated bots.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_listBotReplicasCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexv2Models_listBotReplicasCmd).Standalone()

		lexv2Models_listBotReplicasCmd.Flags().String("bot-id", "", "The request for the unique bot IDs in the list of replicated bots.")
		lexv2Models_listBotReplicasCmd.MarkFlagRequired("bot-id")
	})
	lexv2ModelsCmd.AddCommand(lexv2Models_listBotReplicasCmd)
}
