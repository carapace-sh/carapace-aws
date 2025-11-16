package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dataexchange_updateRevisionCmd = &cobra.Command{
	Use:   "update-revision",
	Short: "This operation updates a revision.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dataexchange_updateRevisionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dataexchange_updateRevisionCmd).Standalone()

		dataexchange_updateRevisionCmd.Flags().String("comment", "", "An optional comment about the revision.")
		dataexchange_updateRevisionCmd.Flags().String("data-set-id", "", "The unique identifier for a data set.")
		dataexchange_updateRevisionCmd.Flags().String("finalized", "", "Finalizing a revision tells AWS Data Exchange that your changes to the assets in the revision are complete.")
		dataexchange_updateRevisionCmd.Flags().String("revision-id", "", "The unique identifier for a revision.")
		dataexchange_updateRevisionCmd.MarkFlagRequired("data-set-id")
		dataexchange_updateRevisionCmd.MarkFlagRequired("revision-id")
	})
	dataexchangeCmd.AddCommand(dataexchange_updateRevisionCmd)
}
