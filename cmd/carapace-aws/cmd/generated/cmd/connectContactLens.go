package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectContactLensCmd = &cobra.Command{
	Use:   "connect-contact-lens",
	Short: "- [Contact Lens actions](https://docs.aws.amazon.com/connect/latest/APIReference/API_Operations_Amazon_Connect_Contact_Lens.html)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectContactLensCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectContactLensCmd).Standalone()

	})
	rootCmd.AddCommand(connectContactLensCmd)
}
