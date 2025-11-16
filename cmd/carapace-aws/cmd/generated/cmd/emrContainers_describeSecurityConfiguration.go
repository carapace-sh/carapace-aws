package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emrContainers_describeSecurityConfigurationCmd = &cobra.Command{
	Use:   "describe-security-configuration",
	Short: "Displays detailed information about a specified security configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emrContainers_describeSecurityConfigurationCmd).Standalone()

	emrContainers_describeSecurityConfigurationCmd.Flags().String("id", "", "The ID of the security configuration.")
	emrContainers_describeSecurityConfigurationCmd.MarkFlagRequired("id")
	emrContainersCmd.AddCommand(emrContainers_describeSecurityConfigurationCmd)
}
