package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_updateAppImageConfigCmd = &cobra.Command{
	Use:   "update-app-image-config",
	Short: "Updates the properties of an AppImageConfig.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_updateAppImageConfigCmd).Standalone()

	sagemaker_updateAppImageConfigCmd.Flags().String("app-image-config-name", "", "The name of the AppImageConfig to update.")
	sagemaker_updateAppImageConfigCmd.Flags().String("code-editor-app-image-config", "", "The Code Editor app running on the image.")
	sagemaker_updateAppImageConfigCmd.Flags().String("jupyter-lab-app-image-config", "", "The JupyterLab app running on the image.")
	sagemaker_updateAppImageConfigCmd.Flags().String("kernel-gateway-image-config", "", "The new KernelGateway app to run on the image.")
	sagemaker_updateAppImageConfigCmd.MarkFlagRequired("app-image-config-name")
	sagemakerCmd.AddCommand(sagemaker_updateAppImageConfigCmd)
}
