package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var drs_describeLaunchConfigurationTemplatesCmd = &cobra.Command{
	Use:   "describe-launch-configuration-templates",
	Short: "Lists all Launch Configuration Templates, filtered by Launch Configuration Template IDs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(drs_describeLaunchConfigurationTemplatesCmd).Standalone()

	drs_describeLaunchConfigurationTemplatesCmd.Flags().String("launch-configuration-template-ids", "", "Request to filter Launch Configuration Templates list by Launch Configuration Template ID.")
	drs_describeLaunchConfigurationTemplatesCmd.Flags().String("max-results", "", "Maximum results to be returned in DescribeLaunchConfigurationTemplates.")
	drs_describeLaunchConfigurationTemplatesCmd.Flags().String("next-token", "", "The token of the next Launch Configuration Template to retrieve.")
	drsCmd.AddCommand(drs_describeLaunchConfigurationTemplatesCmd)
}
