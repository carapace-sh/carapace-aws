package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var detective_startInvestigationCmd = &cobra.Command{
	Use:   "start-investigation",
	Short: "Detective investigations lets you investigate IAM users and IAM roles using indicators of compromise.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(detective_startInvestigationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(detective_startInvestigationCmd).Standalone()

		detective_startInvestigationCmd.Flags().String("entity-arn", "", "The unique Amazon Resource Name (ARN) of the IAM user and IAM role.")
		detective_startInvestigationCmd.Flags().String("graph-arn", "", "The Amazon Resource Name (ARN) of the behavior graph.")
		detective_startInvestigationCmd.Flags().String("scope-end-time", "", "The data and time when the investigation ended.")
		detective_startInvestigationCmd.Flags().String("scope-start-time", "", "The data and time when the investigation began.")
		detective_startInvestigationCmd.MarkFlagRequired("entity-arn")
		detective_startInvestigationCmd.MarkFlagRequired("graph-arn")
		detective_startInvestigationCmd.MarkFlagRequired("scope-end-time")
		detective_startInvestigationCmd.MarkFlagRequired("scope-start-time")
	})
	detectiveCmd.AddCommand(detective_startInvestigationCmd)
}
