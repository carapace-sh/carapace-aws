package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_listCodeSigningConfigsCmd = &cobra.Command{
	Use:   "list-code-signing-configs",
	Short: "Returns a list of [code signing configurations](https://docs.aws.amazon.com/lambda/latest/dg/configuring-codesigning.html). A request returns up to 10,000 configurations per call.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_listCodeSigningConfigsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lambda_listCodeSigningConfigsCmd).Standalone()

		lambda_listCodeSigningConfigsCmd.Flags().String("marker", "", "Specify the pagination token that's returned by a previous request to retrieve the next page of results.")
		lambda_listCodeSigningConfigsCmd.Flags().String("max-items", "", "Maximum number of items to return.")
	})
	lambdaCmd.AddCommand(lambda_listCodeSigningConfigsCmd)
}
