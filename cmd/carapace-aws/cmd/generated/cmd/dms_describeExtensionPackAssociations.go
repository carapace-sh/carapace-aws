package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_describeExtensionPackAssociationsCmd = &cobra.Command{
	Use:   "describe-extension-pack-associations",
	Short: "Returns a paginated list of extension pack associations for the specified migration project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_describeExtensionPackAssociationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_describeExtensionPackAssociationsCmd).Standalone()

		dms_describeExtensionPackAssociationsCmd.Flags().String("filters", "", "Filters applied to the extension pack associations described in the form of key-value pairs.")
		dms_describeExtensionPackAssociationsCmd.Flags().String("marker", "", "Specifies the unique pagination token that makes it possible to display the next page of results.")
		dms_describeExtensionPackAssociationsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
		dms_describeExtensionPackAssociationsCmd.Flags().String("migration-project-identifier", "", "The name or Amazon Resource Name (ARN) for the migration project.")
		dms_describeExtensionPackAssociationsCmd.MarkFlagRequired("migration-project-identifier")
	})
	dmsCmd.AddCommand(dms_describeExtensionPackAssociationsCmd)
}
