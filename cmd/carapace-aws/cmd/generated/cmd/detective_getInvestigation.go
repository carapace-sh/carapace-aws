package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var detective_getInvestigationCmd = &cobra.Command{
	Use:   "get-investigation",
	Short: "Detective investigations lets you investigate IAM users and IAM roles using indicators of compromise.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(detective_getInvestigationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(detective_getInvestigationCmd).Standalone()

		detective_getInvestigationCmd.Flags().String("graph-arn", "", "The Amazon Resource Name (ARN) of the behavior graph.")
		detective_getInvestigationCmd.Flags().String("investigation-id", "", "The investigation ID of the investigation report.")
		detective_getInvestigationCmd.MarkFlagRequired("graph-arn")
		detective_getInvestigationCmd.MarkFlagRequired("investigation-id")
	})
	detectiveCmd.AddCommand(detective_getInvestigationCmd)
}
