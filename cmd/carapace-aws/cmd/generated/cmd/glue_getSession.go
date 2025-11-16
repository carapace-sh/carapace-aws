package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getSessionCmd = &cobra.Command{
	Use:   "get-session",
	Short: "Retrieves the session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getSessionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_getSessionCmd).Standalone()

		glue_getSessionCmd.Flags().String("id", "", "The ID of the session.")
		glue_getSessionCmd.Flags().String("request-origin", "", "The origin of the request.")
		glue_getSessionCmd.MarkFlagRequired("id")
	})
	glueCmd.AddCommand(glue_getSessionCmd)
}
