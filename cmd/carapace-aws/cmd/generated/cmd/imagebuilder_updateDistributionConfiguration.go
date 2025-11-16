package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_updateDistributionConfigurationCmd = &cobra.Command{
	Use:   "update-distribution-configuration",
	Short: "Updates a new distribution configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_updateDistributionConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(imagebuilder_updateDistributionConfigurationCmd).Standalone()

		imagebuilder_updateDistributionConfigurationCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier you provide to ensure idempotency of the request.")
		imagebuilder_updateDistributionConfigurationCmd.Flags().String("description", "", "The description of the distribution configuration.")
		imagebuilder_updateDistributionConfigurationCmd.Flags().String("distribution-configuration-arn", "", "The Amazon Resource Name (ARN) of the distribution configuration that you want to update.")
		imagebuilder_updateDistributionConfigurationCmd.Flags().String("distributions", "", "The distributions of the distribution configuration.")
		imagebuilder_updateDistributionConfigurationCmd.MarkFlagRequired("client-token")
		imagebuilder_updateDistributionConfigurationCmd.MarkFlagRequired("distribution-configuration-arn")
		imagebuilder_updateDistributionConfigurationCmd.MarkFlagRequired("distributions")
	})
	imagebuilderCmd.AddCommand(imagebuilder_updateDistributionConfigurationCmd)
}
