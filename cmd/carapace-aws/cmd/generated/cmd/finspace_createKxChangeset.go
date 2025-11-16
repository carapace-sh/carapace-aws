package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspace_createKxChangesetCmd = &cobra.Command{
	Use:   "create-kx-changeset",
	Short: "Creates a changeset for a kdb database.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspace_createKxChangesetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(finspace_createKxChangesetCmd).Standalone()

		finspace_createKxChangesetCmd.Flags().String("change-requests", "", "A list of change request objects that are run in order.")
		finspace_createKxChangesetCmd.Flags().String("client-token", "", "A token that ensures idempotency.")
		finspace_createKxChangesetCmd.Flags().String("database-name", "", "The name of the kdb database.")
		finspace_createKxChangesetCmd.Flags().String("environment-id", "", "A unique identifier of the kdb environment.")
		finspace_createKxChangesetCmd.MarkFlagRequired("change-requests")
		finspace_createKxChangesetCmd.MarkFlagRequired("client-token")
		finspace_createKxChangesetCmd.MarkFlagRequired("database-name")
		finspace_createKxChangesetCmd.MarkFlagRequired("environment-id")
	})
	finspaceCmd.AddCommand(finspace_createKxChangesetCmd)
}
