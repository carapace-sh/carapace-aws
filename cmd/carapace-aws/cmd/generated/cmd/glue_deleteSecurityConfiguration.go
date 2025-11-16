package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_deleteSecurityConfigurationCmd = &cobra.Command{
	Use:   "delete-security-configuration",
	Short: "Deletes a specified security configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_deleteSecurityConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_deleteSecurityConfigurationCmd).Standalone()

		glue_deleteSecurityConfigurationCmd.Flags().String("name", "", "The name of the security configuration to delete.")
		glue_deleteSecurityConfigurationCmd.MarkFlagRequired("name")
	})
	glueCmd.AddCommand(glue_deleteSecurityConfigurationCmd)
}
