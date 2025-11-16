package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_getInfrastructureConfigurationCmd = &cobra.Command{
	Use:   "get-infrastructure-configuration",
	Short: "Gets an infrastructure configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_getInfrastructureConfigurationCmd).Standalone()

	imagebuilder_getInfrastructureConfigurationCmd.Flags().String("infrastructure-configuration-arn", "", "The Amazon Resource Name (ARN) of the infrastructure configuration that you want to retrieve.")
	imagebuilder_getInfrastructureConfigurationCmd.MarkFlagRequired("infrastructure-configuration-arn")
	imagebuilderCmd.AddCommand(imagebuilder_getInfrastructureConfigurationCmd)
}
