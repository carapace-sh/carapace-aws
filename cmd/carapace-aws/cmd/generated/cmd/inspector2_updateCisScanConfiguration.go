package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_updateCisScanConfigurationCmd = &cobra.Command{
	Use:   "update-cis-scan-configuration",
	Short: "Updates a CIS scan configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_updateCisScanConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector2_updateCisScanConfigurationCmd).Standalone()

		inspector2_updateCisScanConfigurationCmd.Flags().String("scan-configuration-arn", "", "The CIS scan configuration ARN.")
		inspector2_updateCisScanConfigurationCmd.Flags().String("scan-name", "", "The scan name for the CIS scan configuration.")
		inspector2_updateCisScanConfigurationCmd.Flags().String("schedule", "", "The schedule for the CIS scan configuration.")
		inspector2_updateCisScanConfigurationCmd.Flags().String("security-level", "", "The security level for the CIS scan configuration.")
		inspector2_updateCisScanConfigurationCmd.Flags().String("targets", "", "The targets for the CIS scan configuration.")
		inspector2_updateCisScanConfigurationCmd.MarkFlagRequired("scan-configuration-arn")
	})
	inspector2Cmd.AddCommand(inspector2_updateCisScanConfigurationCmd)
}
