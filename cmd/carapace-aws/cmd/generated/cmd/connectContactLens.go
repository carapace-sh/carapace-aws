package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectContactLensCmd = &cobra.Command{
	Use:   "connect-contact-lens",
	Short: "- [Contact Lens actions](https://docs.aws.amazon.com/connect/latest/APIReference/API_Operations_Amazon_Connect_Contact_Lens.html)\n- [Contact Lens data types](https://docs.aws.amazon.com/connect/latest/APIReference/API_Types_Amazon_Connect_Contact_Lens.html)\n\nAmazon Connect Contact Lens enables you to analyze conversations between customer and agents, by using speech transcription, natural language processing, and intelligent search capabilities.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectContactLensCmd).Standalone()

	rootCmd.AddCommand(connectContactLensCmd)
}
