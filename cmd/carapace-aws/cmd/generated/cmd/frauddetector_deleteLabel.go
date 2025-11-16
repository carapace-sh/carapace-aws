package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_deleteLabelCmd = &cobra.Command{
	Use:   "delete-label",
	Short: "Deletes a label.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_deleteLabelCmd).Standalone()

	frauddetector_deleteLabelCmd.Flags().String("name", "", "The name of the label to delete.")
	frauddetector_deleteLabelCmd.MarkFlagRequired("name")
	frauddetectorCmd.AddCommand(frauddetector_deleteLabelCmd)
}
