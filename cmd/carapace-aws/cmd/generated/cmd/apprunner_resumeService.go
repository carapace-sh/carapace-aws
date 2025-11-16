package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apprunner_resumeServiceCmd = &cobra.Command{
	Use:   "resume-service",
	Short: "Resume an active App Runner service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apprunner_resumeServiceCmd).Standalone()

	apprunner_resumeServiceCmd.Flags().String("service-arn", "", "The Amazon Resource Name (ARN) of the App Runner service that you want to resume.")
	apprunner_resumeServiceCmd.MarkFlagRequired("service-arn")
	apprunnerCmd.AddCommand(apprunner_resumeServiceCmd)
}
