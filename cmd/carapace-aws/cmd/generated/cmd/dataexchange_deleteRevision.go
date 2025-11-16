package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dataexchange_deleteRevisionCmd = &cobra.Command{
	Use:   "delete-revision",
	Short: "This operation deletes a revision.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dataexchange_deleteRevisionCmd).Standalone()

	dataexchange_deleteRevisionCmd.Flags().String("data-set-id", "", "The unique identifier for a data set.")
	dataexchange_deleteRevisionCmd.Flags().String("revision-id", "", "The unique identifier for a revision.")
	dataexchange_deleteRevisionCmd.MarkFlagRequired("data-set-id")
	dataexchange_deleteRevisionCmd.MarkFlagRequired("revision-id")
	dataexchangeCmd.AddCommand(dataexchange_deleteRevisionCmd)
}
