package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var b2bi_updatePartnershipCmd = &cobra.Command{
	Use:   "update-partnership",
	Short: "Updates some of the parameters for a partnership between a customer and trading partner.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(b2bi_updatePartnershipCmd).Standalone()

	b2bi_updatePartnershipCmd.Flags().String("capabilities", "", "List of the capabilities associated with this partnership.")
	b2bi_updatePartnershipCmd.Flags().String("capability-options", "", "To update, specify the structure that contains the details for the associated capabilities.")
	b2bi_updatePartnershipCmd.Flags().String("name", "", "The name of the partnership, used to identify it.")
	b2bi_updatePartnershipCmd.Flags().String("partnership-id", "", "Specifies the unique, system-generated identifier for a partnership.")
	b2bi_updatePartnershipCmd.MarkFlagRequired("partnership-id")
	b2biCmd.AddCommand(b2bi_updatePartnershipCmd)
}
