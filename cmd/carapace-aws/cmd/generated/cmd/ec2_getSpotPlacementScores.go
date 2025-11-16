package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getSpotPlacementScoresCmd = &cobra.Command{
	Use:   "get-spot-placement-scores",
	Short: "Calculates the Spot placement score for a Region or Availability Zone based on the specified target capacity and compute requirements.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getSpotPlacementScoresCmd).Standalone()

	ec2_getSpotPlacementScoresCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_getSpotPlacementScoresCmd.Flags().String("instance-requirements-with-metadata", "", "The attributes for the instance types.")
	ec2_getSpotPlacementScoresCmd.Flags().String("instance-types", "", "The instance types.")
	ec2_getSpotPlacementScoresCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
	ec2_getSpotPlacementScoresCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
	ec2_getSpotPlacementScoresCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_getSpotPlacementScoresCmd.Flags().Bool("no-single-availability-zone", false, "Specify `true` so that the response returns a list of scored Availability Zones.")
	ec2_getSpotPlacementScoresCmd.Flags().String("region-names", "", "The Regions used to narrow down the list of Regions to be scored.")
	ec2_getSpotPlacementScoresCmd.Flags().Bool("single-availability-zone", false, "Specify `true` so that the response returns a list of scored Availability Zones.")
	ec2_getSpotPlacementScoresCmd.Flags().String("target-capacity", "", "The target capacity.")
	ec2_getSpotPlacementScoresCmd.Flags().String("target-capacity-unit-type", "", "The unit for the target capacity.")
	ec2_getSpotPlacementScoresCmd.Flag("no-dry-run").Hidden = true
	ec2_getSpotPlacementScoresCmd.Flag("no-single-availability-zone").Hidden = true
	ec2_getSpotPlacementScoresCmd.MarkFlagRequired("target-capacity")
	ec2Cmd.AddCommand(ec2_getSpotPlacementScoresCmd)
}
