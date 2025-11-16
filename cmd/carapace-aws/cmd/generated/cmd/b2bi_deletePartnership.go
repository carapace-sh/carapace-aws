package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var b2bi_deletePartnershipCmd = &cobra.Command{
	Use:   "delete-partnership",
	Short: "Deletes the specified partnership.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(b2bi_deletePartnershipCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(b2bi_deletePartnershipCmd).Standalone()

		b2bi_deletePartnershipCmd.Flags().String("partnership-id", "", "Specifies the unique, system-generated identifier for a partnership.")
		b2bi_deletePartnershipCmd.MarkFlagRequired("partnership-id")
	})
	b2biCmd.AddCommand(b2bi_deletePartnershipCmd)
}
