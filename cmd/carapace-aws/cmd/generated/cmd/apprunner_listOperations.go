package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apprunner_listOperationsCmd = &cobra.Command{
	Use:   "list-operations",
	Short: "Return a list of operations that occurred on an App Runner service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apprunner_listOperationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apprunner_listOperationsCmd).Standalone()

		apprunner_listOperationsCmd.Flags().String("max-results", "", "The maximum number of results to include in each response (result page).")
		apprunner_listOperationsCmd.Flags().String("next-token", "", "A token from a previous result page.")
		apprunner_listOperationsCmd.Flags().String("service-arn", "", "The Amazon Resource Name (ARN) of the App Runner service that you want a list of operations for.")
		apprunner_listOperationsCmd.MarkFlagRequired("service-arn")
	})
	apprunnerCmd.AddCommand(apprunner_listOperationsCmd)
}
