package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_deleteIndexCmd = &cobra.Command{
	Use:   "delete-index",
	Short: "Deletes an Amazon Kendra index.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_deleteIndexCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kendra_deleteIndexCmd).Standalone()

		kendra_deleteIndexCmd.Flags().String("id", "", "The identifier of the index you want to delete.")
		kendra_deleteIndexCmd.MarkFlagRequired("id")
	})
	kendraCmd.AddCommand(kendra_deleteIndexCmd)
}
