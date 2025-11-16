package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_deleteUsagePlanCmd = &cobra.Command{
	Use:   "delete-usage-plan",
	Short: "Deletes a usage plan of a given plan Id.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_deleteUsagePlanCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_deleteUsagePlanCmd).Standalone()

		apigateway_deleteUsagePlanCmd.Flags().String("usage-plan-id", "", "The Id of the to-be-deleted usage plan.")
		apigateway_deleteUsagePlanCmd.MarkFlagRequired("usage-plan-id")
	})
	apigatewayCmd.AddCommand(apigateway_deleteUsagePlanCmd)
}
