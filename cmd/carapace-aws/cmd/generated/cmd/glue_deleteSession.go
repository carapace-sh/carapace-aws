package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_deleteSessionCmd = &cobra.Command{
	Use:   "delete-session",
	Short: "Deletes the session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_deleteSessionCmd).Standalone()

	glue_deleteSessionCmd.Flags().String("id", "", "The ID of the session to be deleted.")
	glue_deleteSessionCmd.Flags().String("request-origin", "", "The name of the origin of the delete session request.")
	glue_deleteSessionCmd.MarkFlagRequired("id")
	glueCmd.AddCommand(glue_deleteSessionCmd)
}
