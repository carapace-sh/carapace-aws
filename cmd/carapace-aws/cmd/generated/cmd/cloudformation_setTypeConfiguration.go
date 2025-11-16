package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_setTypeConfigurationCmd = &cobra.Command{
	Use:   "set-type-configuration",
	Short: "Specifies the configuration data for a CloudFormation extension, such as a resource or Hook, in the given account and Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_setTypeConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudformation_setTypeConfigurationCmd).Standalone()

		cloudformation_setTypeConfigurationCmd.Flags().String("configuration", "", "The configuration data for the extension in this account and Region.")
		cloudformation_setTypeConfigurationCmd.Flags().String("configuration-alias", "", "An alias by which to refer to this extension configuration data.")
		cloudformation_setTypeConfigurationCmd.Flags().String("type", "", "The type of extension.")
		cloudformation_setTypeConfigurationCmd.Flags().String("type-arn", "", "The Amazon Resource Name (ARN) for the extension in this account and Region.")
		cloudformation_setTypeConfigurationCmd.Flags().String("type-name", "", "The name of the extension.")
		cloudformation_setTypeConfigurationCmd.MarkFlagRequired("configuration")
	})
	cloudformationCmd.AddCommand(cloudformation_setTypeConfigurationCmd)
}
