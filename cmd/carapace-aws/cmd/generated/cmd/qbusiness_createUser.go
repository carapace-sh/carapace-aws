package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_createUserCmd = &cobra.Command{
	Use:   "create-user",
	Short: "Creates a universally unique identifier (UUID) mapped to a list of local user ids within an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_createUserCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qbusiness_createUserCmd).Standalone()

		qbusiness_createUserCmd.Flags().String("application-id", "", "The identifier of the application for which the user mapping will be created.")
		qbusiness_createUserCmd.Flags().String("client-token", "", "A token that you provide to identify the request to create your Amazon Q Business user mapping.")
		qbusiness_createUserCmd.Flags().String("user-aliases", "", "The list of user aliases in the mapping.")
		qbusiness_createUserCmd.Flags().String("user-id", "", "The user emails attached to a user mapping.")
		qbusiness_createUserCmd.MarkFlagRequired("application-id")
		qbusiness_createUserCmd.MarkFlagRequired("user-id")
	})
	qbusinessCmd.AddCommand(qbusiness_createUserCmd)
}
