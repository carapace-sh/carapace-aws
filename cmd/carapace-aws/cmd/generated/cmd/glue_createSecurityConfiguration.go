package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_createSecurityConfigurationCmd = &cobra.Command{
	Use:   "create-security-configuration",
	Short: "Creates a new security configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_createSecurityConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_createSecurityConfigurationCmd).Standalone()

		glue_createSecurityConfigurationCmd.Flags().String("encryption-configuration", "", "The encryption configuration for the new security configuration.")
		glue_createSecurityConfigurationCmd.Flags().String("name", "", "The name for the new security configuration.")
		glue_createSecurityConfigurationCmd.MarkFlagRequired("encryption-configuration")
		glue_createSecurityConfigurationCmd.MarkFlagRequired("name")
	})
	glueCmd.AddCommand(glue_createSecurityConfigurationCmd)
}
