package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_listLambdaFunctionsCmd = &cobra.Command{
	Use:   "list-lambda-functions",
	Short: "This API is in preview release for Amazon Connect and is subject to change.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_listLambdaFunctionsCmd).Standalone()

	connect_listLambdaFunctionsCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_listLambdaFunctionsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	connect_listLambdaFunctionsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	connect_listLambdaFunctionsCmd.MarkFlagRequired("instance-id")
	connectCmd.AddCommand(connect_listLambdaFunctionsCmd)
}
