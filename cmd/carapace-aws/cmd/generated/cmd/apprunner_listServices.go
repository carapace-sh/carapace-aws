package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apprunner_listServicesCmd = &cobra.Command{
	Use:   "list-services",
	Short: "Returns a list of running App Runner services in your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apprunner_listServicesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apprunner_listServicesCmd).Standalone()

		apprunner_listServicesCmd.Flags().String("max-results", "", "The maximum number of results to include in each response (result page).")
		apprunner_listServicesCmd.Flags().String("next-token", "", "A token from a previous result page.")
	})
	apprunnerCmd.AddCommand(apprunner_listServicesCmd)
}
