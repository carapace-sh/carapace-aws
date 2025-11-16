package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmIncidents_deleteResponsePlanCmd = &cobra.Command{
	Use:   "delete-response-plan",
	Short: "Deletes the specified response plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmIncidents_deleteResponsePlanCmd).Standalone()

	ssmIncidents_deleteResponsePlanCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the response plan.")
	ssmIncidents_deleteResponsePlanCmd.MarkFlagRequired("arn")
	ssmIncidentsCmd.AddCommand(ssmIncidents_deleteResponsePlanCmd)
}
