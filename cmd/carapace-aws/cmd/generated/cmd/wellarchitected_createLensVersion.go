package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_createLensVersionCmd = &cobra.Command{
	Use:   "create-lens-version",
	Short: "Create a new lens version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_createLensVersionCmd).Standalone()

	wellarchitected_createLensVersionCmd.Flags().String("client-request-token", "", "")
	wellarchitected_createLensVersionCmd.Flags().String("is-major-version", "", "Set to true if this new major lens version.")
	wellarchitected_createLensVersionCmd.Flags().String("lens-alias", "", "")
	wellarchitected_createLensVersionCmd.Flags().String("lens-version", "", "The version of the lens being created.")
	wellarchitected_createLensVersionCmd.MarkFlagRequired("client-request-token")
	wellarchitected_createLensVersionCmd.MarkFlagRequired("lens-alias")
	wellarchitected_createLensVersionCmd.MarkFlagRequired("lens-version")
	wellarchitectedCmd.AddCommand(wellarchitected_createLensVersionCmd)
}
