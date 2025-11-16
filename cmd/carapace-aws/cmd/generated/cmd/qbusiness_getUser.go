package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_getUserCmd = &cobra.Command{
	Use:   "get-user",
	Short: "Describes the universally unique identifier (UUID) associated with a local user in a data source.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_getUserCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qbusiness_getUserCmd).Standalone()

		qbusiness_getUserCmd.Flags().String("application-id", "", "The identifier of the application connected to the user.")
		qbusiness_getUserCmd.Flags().String("user-id", "", "The user email address attached to the user.")
		qbusiness_getUserCmd.MarkFlagRequired("application-id")
		qbusiness_getUserCmd.MarkFlagRequired("user-id")
	})
	qbusinessCmd.AddCommand(qbusiness_getUserCmd)
}
