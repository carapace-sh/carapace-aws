package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var outposts_getSiteCmd = &cobra.Command{
	Use:   "get-site",
	Short: "Gets information about the specified Outpost site.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(outposts_getSiteCmd).Standalone()

	outposts_getSiteCmd.Flags().String("site-id", "", "The ID or the Amazon Resource Name (ARN) of the site.")
	outposts_getSiteCmd.MarkFlagRequired("site-id")
	outpostsCmd.AddCommand(outposts_getSiteCmd)
}
