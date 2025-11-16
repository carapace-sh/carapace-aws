package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dataexchange_getRevisionCmd = &cobra.Command{
	Use:   "get-revision",
	Short: "This operation returns information about a revision.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dataexchange_getRevisionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dataexchange_getRevisionCmd).Standalone()

		dataexchange_getRevisionCmd.Flags().String("data-set-id", "", "The unique identifier for a data set.")
		dataexchange_getRevisionCmd.Flags().String("revision-id", "", "The unique identifier for a revision.")
		dataexchange_getRevisionCmd.MarkFlagRequired("data-set-id")
		dataexchange_getRevisionCmd.MarkFlagRequired("revision-id")
	})
	dataexchangeCmd.AddCommand(dataexchange_getRevisionCmd)
}
