package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_putEnvironmentBlueprintConfigurationCmd = &cobra.Command{
	Use:   "put-environment-blueprint-configuration",
	Short: "Writes the configuration for the specified environment blueprint in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_putEnvironmentBlueprintConfigurationCmd).Standalone()

	datazone_putEnvironmentBlueprintConfigurationCmd.Flags().String("domain-identifier", "", "The identifier of the Amazon DataZone domain.")
	datazone_putEnvironmentBlueprintConfigurationCmd.Flags().String("enabled-regions", "", "Specifies the enabled Amazon Web Services Regions.")
	datazone_putEnvironmentBlueprintConfigurationCmd.Flags().String("environment-blueprint-identifier", "", "The identifier of the environment blueprint.")
	datazone_putEnvironmentBlueprintConfigurationCmd.Flags().String("environment-role-permission-boundary", "", "The environment role permissions boundary.")
	datazone_putEnvironmentBlueprintConfigurationCmd.Flags().String("global-parameters", "", "Region-agnostic environment blueprint parameters.")
	datazone_putEnvironmentBlueprintConfigurationCmd.Flags().String("manage-access-role-arn", "", "The ARN of the manage access role.")
	datazone_putEnvironmentBlueprintConfigurationCmd.Flags().String("provisioning-configurations", "", "The provisioning configuration of a blueprint.")
	datazone_putEnvironmentBlueprintConfigurationCmd.Flags().String("provisioning-role-arn", "", "The ARN of the provisioning role.")
	datazone_putEnvironmentBlueprintConfigurationCmd.Flags().String("regional-parameters", "", "The regional parameters in the environment blueprint.")
	datazone_putEnvironmentBlueprintConfigurationCmd.MarkFlagRequired("domain-identifier")
	datazone_putEnvironmentBlueprintConfigurationCmd.MarkFlagRequired("enabled-regions")
	datazone_putEnvironmentBlueprintConfigurationCmd.MarkFlagRequired("environment-blueprint-identifier")
	datazoneCmd.AddCommand(datazone_putEnvironmentBlueprintConfigurationCmd)
}
