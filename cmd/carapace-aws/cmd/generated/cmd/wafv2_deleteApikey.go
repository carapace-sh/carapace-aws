package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafv2_deleteApikeyCmd = &cobra.Command{
	Use:   "delete-apikey",
	Short: "Deletes the specified API key.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafv2_deleteApikeyCmd).Standalone()

	wafv2_deleteApikeyCmd.Flags().String("apikey", "", "The encrypted API key that you want to delete.")
	wafv2_deleteApikeyCmd.Flags().String("scope", "", "Specifies whether this is for a global resource type, such as a Amazon CloudFront distribution.")
	wafv2_deleteApikeyCmd.MarkFlagRequired("apikey")
	wafv2_deleteApikeyCmd.MarkFlagRequired("scope")
	wafv2Cmd.AddCommand(wafv2_deleteApikeyCmd)
}
