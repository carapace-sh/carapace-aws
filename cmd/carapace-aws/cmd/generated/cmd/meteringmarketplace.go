package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var meteringmarketplaceCmd = &cobra.Command{
	Use:   "meteringmarketplace",
	Short: "Amazon Web Services Marketplace Metering Service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(meteringmarketplaceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(meteringmarketplaceCmd).Standalone()

	})
	rootCmd.AddCommand(meteringmarketplaceCmd)
}
