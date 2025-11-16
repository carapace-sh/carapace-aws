package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_describeFleetAttributesCmd = &cobra.Command{
	Use:   "describe-fleet-attributes",
	Short: "**This API works with the following fleet types:** EC2, Anywhere, Container\n\nRetrieves core fleet-wide properties for fleets in an Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_describeFleetAttributesCmd).Standalone()

	gamelift_describeFleetAttributesCmd.Flags().String("fleet-ids", "", "A list of unique fleet identifiers to retrieve attributes for.")
	gamelift_describeFleetAttributesCmd.Flags().String("limit", "", "The maximum number of results to return.")
	gamelift_describeFleetAttributesCmd.Flags().String("next-token", "", "A token that indicates the start of the next sequential page of results.")
	gameliftCmd.AddCommand(gamelift_describeFleetAttributesCmd)
}
