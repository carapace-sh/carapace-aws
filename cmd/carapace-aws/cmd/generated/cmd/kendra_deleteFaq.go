package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_deleteFaqCmd = &cobra.Command{
	Use:   "delete-faq",
	Short: "Removes a FAQ from an index.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_deleteFaqCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kendra_deleteFaqCmd).Standalone()

		kendra_deleteFaqCmd.Flags().String("id", "", "The identifier of the FAQ you want to remove.")
		kendra_deleteFaqCmd.Flags().String("index-id", "", "The identifier of the index for the FAQ.")
		kendra_deleteFaqCmd.MarkFlagRequired("id")
		kendra_deleteFaqCmd.MarkFlagRequired("index-id")
	})
	kendraCmd.AddCommand(kendra_deleteFaqCmd)
}
