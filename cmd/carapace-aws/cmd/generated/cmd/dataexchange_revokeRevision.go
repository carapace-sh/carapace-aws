package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dataexchange_revokeRevisionCmd = &cobra.Command{
	Use:   "revoke-revision",
	Short: "This operation revokes subscribers' access to a revision.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dataexchange_revokeRevisionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dataexchange_revokeRevisionCmd).Standalone()

		dataexchange_revokeRevisionCmd.Flags().String("data-set-id", "", "The unique identifier for a data set.")
		dataexchange_revokeRevisionCmd.Flags().String("revision-id", "", "The unique identifier for a revision.")
		dataexchange_revokeRevisionCmd.Flags().String("revocation-comment", "", "A required comment to inform subscribers of the reason their access to the revision was revoked.")
		dataexchange_revokeRevisionCmd.MarkFlagRequired("data-set-id")
		dataexchange_revokeRevisionCmd.MarkFlagRequired("revision-id")
		dataexchange_revokeRevisionCmd.MarkFlagRequired("revocation-comment")
	})
	dataexchangeCmd.AddCommand(dataexchange_revokeRevisionCmd)
}
