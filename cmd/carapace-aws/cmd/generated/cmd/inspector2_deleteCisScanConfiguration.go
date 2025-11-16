package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_deleteCisScanConfigurationCmd = &cobra.Command{
	Use:   "delete-cis-scan-configuration",
	Short: "Deletes a CIS scan configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_deleteCisScanConfigurationCmd).Standalone()

	inspector2_deleteCisScanConfigurationCmd.Flags().String("scan-configuration-arn", "", "The ARN of the CIS scan configuration.")
	inspector2_deleteCisScanConfigurationCmd.MarkFlagRequired("scan-configuration-arn")
	inspector2Cmd.AddCommand(inspector2_deleteCisScanConfigurationCmd)
}
