package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apprunner_describeCustomDomainsCmd = &cobra.Command{
	Use:   "describe-custom-domains",
	Short: "Return a description of custom domain names that are associated with an App Runner service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apprunner_describeCustomDomainsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apprunner_describeCustomDomainsCmd).Standalone()

		apprunner_describeCustomDomainsCmd.Flags().String("max-results", "", "The maximum number of results that each response (result page) can include.")
		apprunner_describeCustomDomainsCmd.Flags().String("next-token", "", "A token from a previous result page.")
		apprunner_describeCustomDomainsCmd.Flags().String("service-arn", "", "The Amazon Resource Name (ARN) of the App Runner service that you want associated custom domain names to be described for.")
		apprunner_describeCustomDomainsCmd.MarkFlagRequired("service-arn")
	})
	apprunnerCmd.AddCommand(apprunner_describeCustomDomainsCmd)
}
