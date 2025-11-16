package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_describeExportCmd = &cobra.Command{
	Use:   "describe-export",
	Short: "Gets information about a specific export.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_describeExportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexv2Models_describeExportCmd).Standalone()

		lexv2Models_describeExportCmd.Flags().String("export-id", "", "The unique identifier of the export to describe.")
		lexv2Models_describeExportCmd.MarkFlagRequired("export-id")
	})
	lexv2ModelsCmd.AddCommand(lexv2Models_describeExportCmd)
}
