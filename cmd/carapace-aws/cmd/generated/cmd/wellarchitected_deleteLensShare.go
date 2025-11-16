package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_deleteLensShareCmd = &cobra.Command{
	Use:   "delete-lens-share",
	Short: "Delete a lens share.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_deleteLensShareCmd).Standalone()

	wellarchitected_deleteLensShareCmd.Flags().String("client-request-token", "", "")
	wellarchitected_deleteLensShareCmd.Flags().String("lens-alias", "", "")
	wellarchitected_deleteLensShareCmd.Flags().String("share-id", "", "")
	wellarchitected_deleteLensShareCmd.MarkFlagRequired("client-request-token")
	wellarchitected_deleteLensShareCmd.MarkFlagRequired("lens-alias")
	wellarchitected_deleteLensShareCmd.MarkFlagRequired("share-id")
	wellarchitectedCmd.AddCommand(wellarchitected_deleteLensShareCmd)
}
