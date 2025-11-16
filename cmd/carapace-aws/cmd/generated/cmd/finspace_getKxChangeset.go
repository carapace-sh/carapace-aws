package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspace_getKxChangesetCmd = &cobra.Command{
	Use:   "get-kx-changeset",
	Short: "Returns information about a kdb changeset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspace_getKxChangesetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(finspace_getKxChangesetCmd).Standalone()

		finspace_getKxChangesetCmd.Flags().String("changeset-id", "", "A unique identifier of the changeset for which you want to retrieve data.")
		finspace_getKxChangesetCmd.Flags().String("database-name", "", "The name of the kdb database.")
		finspace_getKxChangesetCmd.Flags().String("environment-id", "", "A unique identifier for the kdb environment.")
		finspace_getKxChangesetCmd.MarkFlagRequired("changeset-id")
		finspace_getKxChangesetCmd.MarkFlagRequired("database-name")
		finspace_getKxChangesetCmd.MarkFlagRequired("environment-id")
	})
	finspaceCmd.AddCommand(finspace_getKxChangesetCmd)
}
