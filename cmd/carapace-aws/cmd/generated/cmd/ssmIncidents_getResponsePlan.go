package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmIncidents_getResponsePlanCmd = &cobra.Command{
	Use:   "get-response-plan",
	Short: "Retrieves the details of the specified response plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmIncidents_getResponsePlanCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmIncidents_getResponsePlanCmd).Standalone()

		ssmIncidents_getResponsePlanCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the response plan.")
		ssmIncidents_getResponsePlanCmd.MarkFlagRequired("arn")
	})
	ssmIncidentsCmd.AddCommand(ssmIncidents_getResponsePlanCmd)
}
