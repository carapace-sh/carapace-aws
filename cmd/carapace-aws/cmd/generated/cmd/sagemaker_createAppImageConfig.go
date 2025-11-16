package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createAppImageConfigCmd = &cobra.Command{
	Use:   "create-app-image-config",
	Short: "Creates a configuration for running a SageMaker AI image as a KernelGateway app.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createAppImageConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_createAppImageConfigCmd).Standalone()

		sagemaker_createAppImageConfigCmd.Flags().String("app-image-config-name", "", "The name of the AppImageConfig.")
		sagemaker_createAppImageConfigCmd.Flags().String("code-editor-app-image-config", "", "The `CodeEditorAppImageConfig`.")
		sagemaker_createAppImageConfigCmd.Flags().String("jupyter-lab-app-image-config", "", "The `JupyterLabAppImageConfig`.")
		sagemaker_createAppImageConfigCmd.Flags().String("kernel-gateway-image-config", "", "The KernelGatewayImageConfig.")
		sagemaker_createAppImageConfigCmd.Flags().String("tags", "", "A list of tags to apply to the AppImageConfig.")
		sagemaker_createAppImageConfigCmd.MarkFlagRequired("app-image-config-name")
	})
	sagemakerCmd.AddCommand(sagemaker_createAppImageConfigCmd)
}
