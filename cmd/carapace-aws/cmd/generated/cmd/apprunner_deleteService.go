package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apprunner_deleteServiceCmd = &cobra.Command{
	Use:   "delete-service",
	Short: "Delete an App Runner service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apprunner_deleteServiceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apprunner_deleteServiceCmd).Standalone()

		apprunner_deleteServiceCmd.Flags().String("service-arn", "", "The Amazon Resource Name (ARN) of the App Runner service that you want to delete.")
		apprunner_deleteServiceCmd.MarkFlagRequired("service-arn")
	})
	apprunnerCmd.AddCommand(apprunner_deleteServiceCmd)
}
