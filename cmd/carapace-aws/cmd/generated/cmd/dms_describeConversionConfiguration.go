package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_describeConversionConfigurationCmd = &cobra.Command{
	Use:   "describe-conversion-configuration",
	Short: "Returns configuration parameters for a schema conversion project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_describeConversionConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_describeConversionConfigurationCmd).Standalone()

		dms_describeConversionConfigurationCmd.Flags().String("migration-project-identifier", "", "The name or Amazon Resource Name (ARN) for the schema conversion project to describe.")
		dms_describeConversionConfigurationCmd.MarkFlagRequired("migration-project-identifier")
	})
	dmsCmd.AddCommand(dms_describeConversionConfigurationCmd)
}
