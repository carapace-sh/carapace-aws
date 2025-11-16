package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_listWavesCmd = &cobra.Command{
	Use:   "list-waves",
	Short: "Retrieves all waves or multiple waves by ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_listWavesCmd).Standalone()

	mgn_listWavesCmd.Flags().String("account-id", "", "Request account ID.")
	mgn_listWavesCmd.Flags().String("filters", "", "Waves list filters.")
	mgn_listWavesCmd.Flags().String("max-results", "", "Maximum results to return when listing waves.")
	mgn_listWavesCmd.Flags().String("next-token", "", "Request next token.")
	mgnCmd.AddCommand(mgn_listWavesCmd)
}
