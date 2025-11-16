package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var detective_updateInvestigationStateCmd = &cobra.Command{
	Use:   "update-investigation-state",
	Short: "Updates the state of an investigation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(detective_updateInvestigationStateCmd).Standalone()

	detective_updateInvestigationStateCmd.Flags().String("graph-arn", "", "The Amazon Resource Name (ARN) of the behavior graph.")
	detective_updateInvestigationStateCmd.Flags().String("investigation-id", "", "The investigation ID of the investigation report.")
	detective_updateInvestigationStateCmd.Flags().String("state", "", "The current state of the investigation.")
	detective_updateInvestigationStateCmd.MarkFlagRequired("graph-arn")
	detective_updateInvestigationStateCmd.MarkFlagRequired("investigation-id")
	detective_updateInvestigationStateCmd.MarkFlagRequired("state")
	detectiveCmd.AddCommand(detective_updateInvestigationStateCmd)
}
