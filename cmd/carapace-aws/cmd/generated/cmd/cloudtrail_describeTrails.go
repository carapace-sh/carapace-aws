package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_describeTrailsCmd = &cobra.Command{
	Use:   "describe-trails",
	Short: "Retrieves settings for one or more trails associated with the current Region for your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_describeTrailsCmd).Standalone()

	cloudtrail_describeTrailsCmd.Flags().Bool("include-shadow-trails", false, "Specifies whether to include shadow trails in the response.")
	cloudtrail_describeTrailsCmd.Flags().Bool("no-include-shadow-trails", false, "Specifies whether to include shadow trails in the response.")
	cloudtrail_describeTrailsCmd.Flags().String("trail-name-list", "", "Specifies a list of trail names, trail ARNs, or both, of the trails to describe.")
	cloudtrail_describeTrailsCmd.Flag("no-include-shadow-trails").Hidden = true
	cloudtrailCmd.AddCommand(cloudtrail_describeTrailsCmd)
}
