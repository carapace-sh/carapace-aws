package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_createLensShareCmd = &cobra.Command{
	Use:   "create-lens-share",
	Short: "Create a lens share.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_createLensShareCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wellarchitected_createLensShareCmd).Standalone()

		wellarchitected_createLensShareCmd.Flags().String("client-request-token", "", "")
		wellarchitected_createLensShareCmd.Flags().String("lens-alias", "", "")
		wellarchitected_createLensShareCmd.Flags().String("shared-with", "", "")
		wellarchitected_createLensShareCmd.MarkFlagRequired("client-request-token")
		wellarchitected_createLensShareCmd.MarkFlagRequired("lens-alias")
		wellarchitected_createLensShareCmd.MarkFlagRequired("shared-with")
	})
	wellarchitectedCmd.AddCommand(wellarchitected_createLensShareCmd)
}
