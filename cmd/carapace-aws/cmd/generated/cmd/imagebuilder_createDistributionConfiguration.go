package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_createDistributionConfigurationCmd = &cobra.Command{
	Use:   "create-distribution-configuration",
	Short: "Creates a new distribution configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_createDistributionConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(imagebuilder_createDistributionConfigurationCmd).Standalone()

		imagebuilder_createDistributionConfigurationCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier you provide to ensure idempotency of the request.")
		imagebuilder_createDistributionConfigurationCmd.Flags().String("description", "", "The description of the distribution configuration.")
		imagebuilder_createDistributionConfigurationCmd.Flags().String("distributions", "", "The distributions of the distribution configuration.")
		imagebuilder_createDistributionConfigurationCmd.Flags().String("name", "", "The name of the distribution configuration.")
		imagebuilder_createDistributionConfigurationCmd.Flags().String("tags", "", "The tags of the distribution configuration.")
		imagebuilder_createDistributionConfigurationCmd.MarkFlagRequired("client-token")
		imagebuilder_createDistributionConfigurationCmd.MarkFlagRequired("distributions")
		imagebuilder_createDistributionConfigurationCmd.MarkFlagRequired("name")
	})
	imagebuilderCmd.AddCommand(imagebuilder_createDistributionConfigurationCmd)
}
