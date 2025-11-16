package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apprunner_describeServiceCmd = &cobra.Command{
	Use:   "describe-service",
	Short: "Return a full description of an App Runner service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apprunner_describeServiceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apprunner_describeServiceCmd).Standalone()

		apprunner_describeServiceCmd.Flags().String("service-arn", "", "The Amazon Resource Name (ARN) of the App Runner service that you want a description for.")
		apprunner_describeServiceCmd.MarkFlagRequired("service-arn")
	})
	apprunnerCmd.AddCommand(apprunner_describeServiceCmd)
}
