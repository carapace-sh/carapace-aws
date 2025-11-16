package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_describeLaunchConfigurationTemplatesCmd = &cobra.Command{
	Use:   "describe-launch-configuration-templates",
	Short: "Lists all Launch Configuration Templates, filtered by Launch Configuration Template IDs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_describeLaunchConfigurationTemplatesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mgn_describeLaunchConfigurationTemplatesCmd).Standalone()

		mgn_describeLaunchConfigurationTemplatesCmd.Flags().String("launch-configuration-template-ids", "", "Request to filter Launch Configuration Templates list by Launch Configuration Template ID.")
		mgn_describeLaunchConfigurationTemplatesCmd.Flags().String("max-results", "", "Maximum results to be returned in DescribeLaunchConfigurationTemplates.")
		mgn_describeLaunchConfigurationTemplatesCmd.Flags().String("next-token", "", "Next pagination token returned from DescribeLaunchConfigurationTemplates.")
	})
	mgnCmd.AddCommand(mgn_describeLaunchConfigurationTemplatesCmd)
}
