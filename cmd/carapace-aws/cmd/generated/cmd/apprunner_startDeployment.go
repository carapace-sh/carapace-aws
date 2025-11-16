package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apprunner_startDeploymentCmd = &cobra.Command{
	Use:   "start-deployment",
	Short: "Initiate a manual deployment of the latest commit in a source code repository or the latest image in a source image repository to an App Runner service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apprunner_startDeploymentCmd).Standalone()

	apprunner_startDeploymentCmd.Flags().String("service-arn", "", "The Amazon Resource Name (ARN) of the App Runner service that you want to manually deploy to.")
	apprunner_startDeploymentCmd.MarkFlagRequired("service-arn")
	apprunnerCmd.AddCommand(apprunner_startDeploymentCmd)
}
