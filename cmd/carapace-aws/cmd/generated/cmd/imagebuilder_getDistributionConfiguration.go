package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_getDistributionConfigurationCmd = &cobra.Command{
	Use:   "get-distribution-configuration",
	Short: "Gets a distribution configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_getDistributionConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(imagebuilder_getDistributionConfigurationCmd).Standalone()

		imagebuilder_getDistributionConfigurationCmd.Flags().String("distribution-configuration-arn", "", "The Amazon Resource Name (ARN) of the distribution configuration that you want to retrieve.")
		imagebuilder_getDistributionConfigurationCmd.MarkFlagRequired("distribution-configuration-arn")
	})
	imagebuilderCmd.AddCommand(imagebuilder_getDistributionConfigurationCmd)
}
