package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dataexchange_createRevisionCmd = &cobra.Command{
	Use:   "create-revision",
	Short: "This operation creates a revision for a data set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dataexchange_createRevisionCmd).Standalone()

	dataexchange_createRevisionCmd.Flags().String("comment", "", "An optional comment about the revision.")
	dataexchange_createRevisionCmd.Flags().String("data-set-id", "", "The unique identifier for a data set.")
	dataexchange_createRevisionCmd.Flags().String("tags", "", "A revision tag is an optional label that you can assign to a revision when you create it.")
	dataexchange_createRevisionCmd.MarkFlagRequired("data-set-id")
	dataexchangeCmd.AddCommand(dataexchange_createRevisionCmd)
}
