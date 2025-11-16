package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dataexchange_acceptDataGrantCmd = &cobra.Command{
	Use:   "accept-data-grant",
	Short: "This operation accepts a data grant.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dataexchange_acceptDataGrantCmd).Standalone()

	dataexchange_acceptDataGrantCmd.Flags().String("data-grant-arn", "", "The Amazon Resource Name (ARN) of the data grant to accept.")
	dataexchange_acceptDataGrantCmd.MarkFlagRequired("data-grant-arn")
	dataexchangeCmd.AddCommand(dataexchange_acceptDataGrantCmd)
}
