package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var m2_updateApplicationCmd = &cobra.Command{
	Use:   "update-application",
	Short: "Updates an application and creates a new version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(m2_updateApplicationCmd).Standalone()

	m2_updateApplicationCmd.Flags().String("application-id", "", "The unique identifier of the application you want to update.")
	m2_updateApplicationCmd.Flags().String("current-application-version", "", "The current version of the application to update.")
	m2_updateApplicationCmd.Flags().String("definition", "", "The application definition for this application.")
	m2_updateApplicationCmd.Flags().String("description", "", "The description of the application to update.")
	m2_updateApplicationCmd.MarkFlagRequired("application-id")
	m2_updateApplicationCmd.MarkFlagRequired("current-application-version")
	m2Cmd.AddCommand(m2_updateApplicationCmd)
}
