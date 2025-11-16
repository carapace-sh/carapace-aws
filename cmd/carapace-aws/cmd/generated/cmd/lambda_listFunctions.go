package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_listFunctionsCmd = &cobra.Command{
	Use:   "list-functions",
	Short: "Returns a list of Lambda functions, with the version-specific configuration of each.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_listFunctionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lambda_listFunctionsCmd).Standalone()

		lambda_listFunctionsCmd.Flags().String("function-version", "", "Set to `ALL` to include entries for all published versions of each function.")
		lambda_listFunctionsCmd.Flags().String("marker", "", "Specify the pagination token that's returned by a previous request to retrieve the next page of results.")
		lambda_listFunctionsCmd.Flags().String("master-region", "", "For Lambda@Edge functions, the Amazon Web Services Region of the master function.")
		lambda_listFunctionsCmd.Flags().String("max-items", "", "The maximum number of functions to return in the response.")
	})
	lambdaCmd.AddCommand(lambda_listFunctionsCmd)
}
