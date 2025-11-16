package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_listBotAliasReplicasCmd = &cobra.Command{
	Use:   "list-bot-alias-replicas",
	Short: "The action to list the replicated bots created from the source bot alias.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_listBotAliasReplicasCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexv2Models_listBotAliasReplicasCmd).Standalone()

		lexv2Models_listBotAliasReplicasCmd.Flags().String("bot-id", "", "The request for the unique bot ID of the replicated bot created from the source bot alias.")
		lexv2Models_listBotAliasReplicasCmd.Flags().String("max-results", "", "The request for maximum results to list the replicated bots created from the source bot alias.")
		lexv2Models_listBotAliasReplicasCmd.Flags().String("next-token", "", "The request for the next token for the replicated bot created from the source bot alias.")
		lexv2Models_listBotAliasReplicasCmd.Flags().String("replica-region", "", "The request for the secondary region of the replicated bot created from the source bot alias.")
		lexv2Models_listBotAliasReplicasCmd.MarkFlagRequired("bot-id")
		lexv2Models_listBotAliasReplicasCmd.MarkFlagRequired("replica-region")
	})
	lexv2ModelsCmd.AddCommand(lexv2Models_listBotAliasReplicasCmd)
}
