package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecr_putImageScanningConfigurationCmd = &cobra.Command{
	Use:   "put-image-scanning-configuration",
	Short: "The `PutImageScanningConfiguration` API is being deprecated, in favor of specifying the image scanning configuration at the registry level.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecr_putImageScanningConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecr_putImageScanningConfigurationCmd).Standalone()

		ecr_putImageScanningConfigurationCmd.Flags().String("image-scanning-configuration", "", "The image scanning configuration for the repository.")
		ecr_putImageScanningConfigurationCmd.Flags().String("registry-id", "", "The Amazon Web Services account ID associated with the registry that contains the repository in which to update the image scanning configuration setting.")
		ecr_putImageScanningConfigurationCmd.Flags().String("repository-name", "", "The name of the repository in which to update the image scanning configuration setting.")
		ecr_putImageScanningConfigurationCmd.MarkFlagRequired("image-scanning-configuration")
		ecr_putImageScanningConfigurationCmd.MarkFlagRequired("repository-name")
	})
	ecrCmd.AddCommand(ecr_putImageScanningConfigurationCmd)
}
