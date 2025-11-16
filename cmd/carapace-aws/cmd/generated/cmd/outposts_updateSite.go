package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var outposts_updateSiteCmd = &cobra.Command{
	Use:   "update-site",
	Short: "Updates the specified site.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(outposts_updateSiteCmd).Standalone()

	outposts_updateSiteCmd.Flags().String("description", "", "")
	outposts_updateSiteCmd.Flags().String("name", "", "")
	outposts_updateSiteCmd.Flags().String("notes", "", "Notes about a site.")
	outposts_updateSiteCmd.Flags().String("site-id", "", "The ID or the Amazon Resource Name (ARN) of the site.")
	outposts_updateSiteCmd.MarkFlagRequired("site-id")
	outpostsCmd.AddCommand(outposts_updateSiteCmd)
}
