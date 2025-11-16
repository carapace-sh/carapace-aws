package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_renderUiTemplateCmd = &cobra.Command{
	Use:   "render-ui-template",
	Short: "Renders the UI template so that you can preview the worker's experience.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_renderUiTemplateCmd).Standalone()

	sagemaker_renderUiTemplateCmd.Flags().String("human-task-ui-arn", "", "The `HumanTaskUiArn` of the worker UI that you want to render.")
	sagemaker_renderUiTemplateCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) that has access to the S3 objects that are used by the template.")
	sagemaker_renderUiTemplateCmd.Flags().String("task", "", "A `RenderableTask` object containing a representative task to render.")
	sagemaker_renderUiTemplateCmd.Flags().String("ui-template", "", "A `Template` object containing the worker UI template to render.")
	sagemaker_renderUiTemplateCmd.MarkFlagRequired("role-arn")
	sagemaker_renderUiTemplateCmd.MarkFlagRequired("task")
	sagemakerCmd.AddCommand(sagemaker_renderUiTemplateCmd)
}
