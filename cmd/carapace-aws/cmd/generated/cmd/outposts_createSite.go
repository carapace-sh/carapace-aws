package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var outposts_createSiteCmd = &cobra.Command{
	Use:   "create-site",
	Short: "Creates a site for an Outpost.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(outposts_createSiteCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(outposts_createSiteCmd).Standalone()

		outposts_createSiteCmd.Flags().String("description", "", "")
		outposts_createSiteCmd.Flags().String("name", "", "")
		outposts_createSiteCmd.Flags().String("notes", "", "Additional information that you provide about site access requirements, electrician scheduling, personal protective equipment, or regulation of equipment materials that could affect your installation process.")
		outposts_createSiteCmd.Flags().String("operating-address", "", "The location to install and power on the hardware.")
		outposts_createSiteCmd.Flags().String("rack-physical-properties", "", "Information about the physical and logistical details for the rack at this site.")
		outposts_createSiteCmd.Flags().String("shipping-address", "", "The location to ship the hardware.")
		outposts_createSiteCmd.Flags().String("tags", "", "The tags to apply to a site.")
		outposts_createSiteCmd.MarkFlagRequired("name")
	})
	outpostsCmd.AddCommand(outposts_createSiteCmd)
}
