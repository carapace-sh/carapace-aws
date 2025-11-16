package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emrContainers_createSecurityConfigurationCmd = &cobra.Command{
	Use:   "create-security-configuration",
	Short: "Creates a security configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emrContainers_createSecurityConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(emrContainers_createSecurityConfigurationCmd).Standalone()

		emrContainers_createSecurityConfigurationCmd.Flags().String("client-token", "", "The client idempotency token to use when creating the security configuration.")
		emrContainers_createSecurityConfigurationCmd.Flags().String("container-provider", "", "The container provider associated with the security configuration.")
		emrContainers_createSecurityConfigurationCmd.Flags().String("name", "", "The name of the security configuration.")
		emrContainers_createSecurityConfigurationCmd.Flags().String("security-configuration-data", "", "Security configuration input for the request.")
		emrContainers_createSecurityConfigurationCmd.Flags().String("tags", "", "The tags to add to the security configuration.")
		emrContainers_createSecurityConfigurationCmd.MarkFlagRequired("client-token")
		emrContainers_createSecurityConfigurationCmd.MarkFlagRequired("name")
		emrContainers_createSecurityConfigurationCmd.MarkFlagRequired("security-configuration-data")
	})
	emrContainersCmd.AddCommand(emrContainers_createSecurityConfigurationCmd)
}
