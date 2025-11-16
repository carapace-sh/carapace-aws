package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_listIntentPathsCmd = &cobra.Command{
	Use:   "list-intent-paths",
	Short: "Retrieves summary statistics for a path of intents that users take over sessions with your bot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_listIntentPathsCmd).Standalone()

	lexv2Models_listIntentPathsCmd.Flags().String("bot-id", "", "The identifier for the bot for which you want to retrieve intent path metrics.")
	lexv2Models_listIntentPathsCmd.Flags().String("end-date-time", "", "The date and time that marks the end of the range of time for which you want to see intent path metrics.")
	lexv2Models_listIntentPathsCmd.Flags().String("filters", "", "A list of objects, each describes a condition by which you want to filter the results.")
	lexv2Models_listIntentPathsCmd.Flags().String("intent-path", "", "The intent path for which you want to retrieve metrics.")
	lexv2Models_listIntentPathsCmd.Flags().String("start-date-time", "", "The date and time that marks the beginning of the range of time for which you want to see intent path metrics.")
	lexv2Models_listIntentPathsCmd.MarkFlagRequired("bot-id")
	lexv2Models_listIntentPathsCmd.MarkFlagRequired("end-date-time")
	lexv2Models_listIntentPathsCmd.MarkFlagRequired("intent-path")
	lexv2Models_listIntentPathsCmd.MarkFlagRequired("start-date-time")
	lexv2ModelsCmd.AddCommand(lexv2Models_listIntentPathsCmd)
}
