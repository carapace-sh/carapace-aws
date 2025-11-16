package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_modifyConversionConfigurationCmd = &cobra.Command{
	Use:   "modify-conversion-configuration",
	Short: "Modifies the specified schema conversion configuration using the provided parameters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_modifyConversionConfigurationCmd).Standalone()

	dms_modifyConversionConfigurationCmd.Flags().String("conversion-configuration", "", "The new conversion configuration.")
	dms_modifyConversionConfigurationCmd.Flags().String("migration-project-identifier", "", "The migration project name or Amazon Resource Name (ARN).")
	dms_modifyConversionConfigurationCmd.MarkFlagRequired("conversion-configuration")
	dms_modifyConversionConfigurationCmd.MarkFlagRequired("migration-project-identifier")
	dmsCmd.AddCommand(dms_modifyConversionConfigurationCmd)
}
