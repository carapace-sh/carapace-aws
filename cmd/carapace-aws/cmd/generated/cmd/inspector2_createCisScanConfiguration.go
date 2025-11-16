package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_createCisScanConfigurationCmd = &cobra.Command{
	Use:   "create-cis-scan-configuration",
	Short: "Creates a CIS scan configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_createCisScanConfigurationCmd).Standalone()

	inspector2_createCisScanConfigurationCmd.Flags().String("scan-name", "", "The scan name for the CIS scan configuration.")
	inspector2_createCisScanConfigurationCmd.Flags().String("schedule", "", "The schedule for the CIS scan configuration.")
	inspector2_createCisScanConfigurationCmd.Flags().String("security-level", "", "The security level for the CIS scan configuration.")
	inspector2_createCisScanConfigurationCmd.Flags().String("tags", "", "The tags for the CIS scan configuration.")
	inspector2_createCisScanConfigurationCmd.Flags().String("targets", "", "The targets for the CIS scan configuration.")
	inspector2_createCisScanConfigurationCmd.MarkFlagRequired("scan-name")
	inspector2_createCisScanConfigurationCmd.MarkFlagRequired("schedule")
	inspector2_createCisScanConfigurationCmd.MarkFlagRequired("security-level")
	inspector2_createCisScanConfigurationCmd.MarkFlagRequired("targets")
	inspector2Cmd.AddCommand(inspector2_createCisScanConfigurationCmd)
}
