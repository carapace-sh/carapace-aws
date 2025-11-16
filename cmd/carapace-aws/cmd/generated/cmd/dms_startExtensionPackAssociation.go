package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_startExtensionPackAssociationCmd = &cobra.Command{
	Use:   "start-extension-pack-association",
	Short: "Applies the extension pack to your target database.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_startExtensionPackAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_startExtensionPackAssociationCmd).Standalone()

		dms_startExtensionPackAssociationCmd.Flags().String("migration-project-identifier", "", "The migration project name or Amazon Resource Name (ARN).")
		dms_startExtensionPackAssociationCmd.MarkFlagRequired("migration-project-identifier")
	})
	dmsCmd.AddCommand(dms_startExtensionPackAssociationCmd)
}
