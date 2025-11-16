package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalizeEvents_putItemsCmd = &cobra.Command{
	Use:   "put-items",
	Short: "Adds one or more items to an Items dataset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalizeEvents_putItemsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(personalizeEvents_putItemsCmd).Standalone()

		personalizeEvents_putItemsCmd.Flags().String("dataset-arn", "", "The Amazon Resource Name (ARN) of the Items dataset you are adding the item or items to.")
		personalizeEvents_putItemsCmd.Flags().String("items", "", "A list of item data.")
		personalizeEvents_putItemsCmd.MarkFlagRequired("dataset-arn")
		personalizeEvents_putItemsCmd.MarkFlagRequired("items")
	})
	personalizeEventsCmd.AddCommand(personalizeEvents_putItemsCmd)
}
