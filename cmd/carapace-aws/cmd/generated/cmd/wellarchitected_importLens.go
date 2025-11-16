package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_importLensCmd = &cobra.Command{
	Use:   "import-lens",
	Short: "Import a new custom lens or update an existing custom lens.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_importLensCmd).Standalone()

	wellarchitected_importLensCmd.Flags().String("client-request-token", "", "")
	wellarchitected_importLensCmd.Flags().String("jsonstring", "", "The JSON representation of a lens.")
	wellarchitected_importLensCmd.Flags().String("lens-alias", "", "")
	wellarchitected_importLensCmd.Flags().String("tags", "", "Tags to associate to a lens.")
	wellarchitected_importLensCmd.MarkFlagRequired("client-request-token")
	wellarchitected_importLensCmd.MarkFlagRequired("jsonstring")
	wellarchitectedCmd.AddCommand(wellarchitected_importLensCmd)
}
