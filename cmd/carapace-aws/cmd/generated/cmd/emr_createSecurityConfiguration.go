package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_createSecurityConfigurationCmd = &cobra.Command{
	Use:   "create-security-configuration",
	Short: "Creates a security configuration, which is stored in the service and can be specified when a cluster is created.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_createSecurityConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(emr_createSecurityConfigurationCmd).Standalone()

		emr_createSecurityConfigurationCmd.Flags().String("name", "", "The name of the security configuration.")
		emr_createSecurityConfigurationCmd.Flags().String("security-configuration", "", "The security configuration details in JSON format.")
		emr_createSecurityConfigurationCmd.MarkFlagRequired("name")
		emr_createSecurityConfigurationCmd.MarkFlagRequired("security-configuration")
	})
	emrCmd.AddCommand(emr_createSecurityConfigurationCmd)
}
