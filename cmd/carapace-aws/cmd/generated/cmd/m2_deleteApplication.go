package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var m2_deleteApplicationCmd = &cobra.Command{
	Use:   "delete-application",
	Short: "Deletes a specific application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(m2_deleteApplicationCmd).Standalone()

	m2_deleteApplicationCmd.Flags().String("application-id", "", "The unique identifier of the application you want to delete.")
	m2_deleteApplicationCmd.MarkFlagRequired("application-id")
	m2Cmd.AddCommand(m2_deleteApplicationCmd)
}
