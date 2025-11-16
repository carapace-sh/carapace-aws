package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_listBotVersionReplicasCmd = &cobra.Command{
	Use:   "list-bot-version-replicas",
	Short: "Contains information about all the versions replication statuses applicable for Global Resiliency.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_listBotVersionReplicasCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexv2Models_listBotVersionReplicasCmd).Standalone()

		lexv2Models_listBotVersionReplicasCmd.Flags().String("bot-id", "", "The request for the unique ID in the list of replicated bots.")
		lexv2Models_listBotVersionReplicasCmd.Flags().String("max-results", "", "The maximum results given in the list of replicated bots.")
		lexv2Models_listBotVersionReplicasCmd.Flags().String("next-token", "", "The next token given in the list of replicated bots.")
		lexv2Models_listBotVersionReplicasCmd.Flags().String("replica-region", "", "The request for the region used in the list of replicated bots.")
		lexv2Models_listBotVersionReplicasCmd.Flags().String("sort-by", "", "The requested sort category for the list of replicated bots.")
		lexv2Models_listBotVersionReplicasCmd.MarkFlagRequired("bot-id")
		lexv2Models_listBotVersionReplicasCmd.MarkFlagRequired("replica-region")
	})
	lexv2ModelsCmd.AddCommand(lexv2Models_listBotVersionReplicasCmd)
}
