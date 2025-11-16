package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_describeFaqCmd = &cobra.Command{
	Use:   "describe-faq",
	Short: "Gets information about a FAQ.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_describeFaqCmd).Standalone()

	kendra_describeFaqCmd.Flags().String("id", "", "The identifier of the FAQ you want to get information on.")
	kendra_describeFaqCmd.Flags().String("index-id", "", "The identifier of the index for the FAQ.")
	kendra_describeFaqCmd.MarkFlagRequired("id")
	kendra_describeFaqCmd.MarkFlagRequired("index-id")
	kendraCmd.AddCommand(kendra_describeFaqCmd)
}
