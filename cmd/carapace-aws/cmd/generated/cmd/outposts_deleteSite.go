package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var outposts_deleteSiteCmd = &cobra.Command{
	Use:   "delete-site",
	Short: "Deletes the specified site.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(outposts_deleteSiteCmd).Standalone()

	outposts_deleteSiteCmd.Flags().String("site-id", "", "The ID or the Amazon Resource Name (ARN) of the site.")
	outposts_deleteSiteCmd.MarkFlagRequired("site-id")
	outpostsCmd.AddCommand(outposts_deleteSiteCmd)
}
