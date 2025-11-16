package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_deleteDistributionConfigurationCmd = &cobra.Command{
	Use:   "delete-distribution-configuration",
	Short: "Deletes a distribution configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_deleteDistributionConfigurationCmd).Standalone()

	imagebuilder_deleteDistributionConfigurationCmd.Flags().String("distribution-configuration-arn", "", "The Amazon Resource Name (ARN) of the distribution configuration to delete.")
	imagebuilder_deleteDistributionConfigurationCmd.MarkFlagRequired("distribution-configuration-arn")
	imagebuilderCmd.AddCommand(imagebuilder_deleteDistributionConfigurationCmd)
}
