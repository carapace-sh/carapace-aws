package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_deleteInfrastructureConfigurationCmd = &cobra.Command{
	Use:   "delete-infrastructure-configuration",
	Short: "Deletes an infrastructure configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_deleteInfrastructureConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(imagebuilder_deleteInfrastructureConfigurationCmd).Standalone()

		imagebuilder_deleteInfrastructureConfigurationCmd.Flags().String("infrastructure-configuration-arn", "", "The Amazon Resource Name (ARN) of the infrastructure configuration to delete.")
		imagebuilder_deleteInfrastructureConfigurationCmd.MarkFlagRequired("infrastructure-configuration-arn")
	})
	imagebuilderCmd.AddCommand(imagebuilder_deleteInfrastructureConfigurationCmd)
}
