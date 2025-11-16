package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getSecurityConfigurationCmd = &cobra.Command{
	Use:   "get-security-configuration",
	Short: "Retrieves a specified security configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getSecurityConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_getSecurityConfigurationCmd).Standalone()

		glue_getSecurityConfigurationCmd.Flags().String("name", "", "The name of the security configuration to retrieve.")
		glue_getSecurityConfigurationCmd.MarkFlagRequired("name")
	})
	glueCmd.AddCommand(glue_getSecurityConfigurationCmd)
}
