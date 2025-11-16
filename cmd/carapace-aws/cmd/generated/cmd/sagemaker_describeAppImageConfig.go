package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeAppImageConfigCmd = &cobra.Command{
	Use:   "describe-app-image-config",
	Short: "Describes an AppImageConfig.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeAppImageConfigCmd).Standalone()

	sagemaker_describeAppImageConfigCmd.Flags().String("app-image-config-name", "", "The name of the AppImageConfig to describe.")
	sagemaker_describeAppImageConfigCmd.MarkFlagRequired("app-image-config-name")
	sagemakerCmd.AddCommand(sagemaker_describeAppImageConfigCmd)
}
