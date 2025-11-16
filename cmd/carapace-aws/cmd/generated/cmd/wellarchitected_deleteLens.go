package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_deleteLensCmd = &cobra.Command{
	Use:   "delete-lens",
	Short: "Delete an existing lens.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_deleteLensCmd).Standalone()

	wellarchitected_deleteLensCmd.Flags().String("client-request-token", "", "")
	wellarchitected_deleteLensCmd.Flags().String("lens-alias", "", "")
	wellarchitected_deleteLensCmd.Flags().String("lens-status", "", "The status of the lens to be deleted.")
	wellarchitected_deleteLensCmd.MarkFlagRequired("client-request-token")
	wellarchitected_deleteLensCmd.MarkFlagRequired("lens-alias")
	wellarchitected_deleteLensCmd.MarkFlagRequired("lens-status")
	wellarchitectedCmd.AddCommand(wellarchitected_deleteLensCmd)
}
