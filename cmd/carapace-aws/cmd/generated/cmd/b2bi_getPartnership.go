package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var b2bi_getPartnershipCmd = &cobra.Command{
	Use:   "get-partnership",
	Short: "Retrieves the details for a partnership, based on the partner and profile IDs specified.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(b2bi_getPartnershipCmd).Standalone()

	b2bi_getPartnershipCmd.Flags().String("partnership-id", "", "Specifies the unique, system-generated identifier for a partnership.")
	b2bi_getPartnershipCmd.MarkFlagRequired("partnership-id")
	b2biCmd.AddCommand(b2bi_getPartnershipCmd)
}
