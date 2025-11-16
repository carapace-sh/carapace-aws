package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_stopSessionCmd = &cobra.Command{
	Use:   "stop-session",
	Short: "Stops the session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_stopSessionCmd).Standalone()

	glue_stopSessionCmd.Flags().String("id", "", "The ID of the session to be stopped.")
	glue_stopSessionCmd.Flags().String("request-origin", "", "The origin of the request.")
	glue_stopSessionCmd.MarkFlagRequired("id")
	glueCmd.AddCommand(glue_stopSessionCmd)
}
