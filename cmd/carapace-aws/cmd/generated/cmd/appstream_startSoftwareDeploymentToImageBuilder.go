package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_startSoftwareDeploymentToImageBuilderCmd = &cobra.Command{
	Use:   "start-software-deployment-to-image-builder",
	Short: "Initiates license included applications deployment to an image builder instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_startSoftwareDeploymentToImageBuilderCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appstream_startSoftwareDeploymentToImageBuilderCmd).Standalone()

		appstream_startSoftwareDeploymentToImageBuilderCmd.Flags().String("image-builder-name", "", "The name of the target image builder instance.")
		appstream_startSoftwareDeploymentToImageBuilderCmd.Flags().Bool("no-retry-failed-deployments", false, "Whether to retry previously failed license included application deployments.")
		appstream_startSoftwareDeploymentToImageBuilderCmd.Flags().Bool("retry-failed-deployments", false, "Whether to retry previously failed license included application deployments.")
		appstream_startSoftwareDeploymentToImageBuilderCmd.MarkFlagRequired("image-builder-name")
		appstream_startSoftwareDeploymentToImageBuilderCmd.Flag("no-retry-failed-deployments").Hidden = true
	})
	appstreamCmd.AddCommand(appstream_startSoftwareDeploymentToImageBuilderCmd)
}
