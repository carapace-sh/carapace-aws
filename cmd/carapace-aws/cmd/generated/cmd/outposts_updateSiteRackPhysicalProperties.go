package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var outposts_updateSiteRackPhysicalPropertiesCmd = &cobra.Command{
	Use:   "update-site-rack-physical-properties",
	Short: "Update the physical and logistical details for a rack at a site.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(outposts_updateSiteRackPhysicalPropertiesCmd).Standalone()

	outposts_updateSiteRackPhysicalPropertiesCmd.Flags().String("fiber-optic-cable-type", "", "The type of fiber that you will use to attach the Outpost to your network.")
	outposts_updateSiteRackPhysicalPropertiesCmd.Flags().String("maximum-supported-weight-lbs", "", "The maximum rack weight that this site can support.")
	outposts_updateSiteRackPhysicalPropertiesCmd.Flags().String("optical-standard", "", "The type of optical standard that you will use to attach the Outpost to your network.")
	outposts_updateSiteRackPhysicalPropertiesCmd.Flags().String("power-connector", "", "The power connector that Amazon Web Services should plan to provide for connections to the hardware.")
	outposts_updateSiteRackPhysicalPropertiesCmd.Flags().String("power-draw-kva", "", "The power draw, in kVA, available at the hardware placement position for the rack.")
	outposts_updateSiteRackPhysicalPropertiesCmd.Flags().String("power-feed-drop", "", "Indicates whether the power feed comes above or below the rack.")
	outposts_updateSiteRackPhysicalPropertiesCmd.Flags().String("power-phase", "", "The power option that you can provide for hardware.")
	outposts_updateSiteRackPhysicalPropertiesCmd.Flags().String("site-id", "", "The ID or the Amazon Resource Name (ARN) of the site.")
	outposts_updateSiteRackPhysicalPropertiesCmd.Flags().String("uplink-count", "", "Racks come with two Outpost network devices.")
	outposts_updateSiteRackPhysicalPropertiesCmd.Flags().String("uplink-gbps", "", "The uplink speed the rack should support for the connection to the Region.")
	outposts_updateSiteRackPhysicalPropertiesCmd.MarkFlagRequired("site-id")
	outpostsCmd.AddCommand(outposts_updateSiteRackPhysicalPropertiesCmd)
}
