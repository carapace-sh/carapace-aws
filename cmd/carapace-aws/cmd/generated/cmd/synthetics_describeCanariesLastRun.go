package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var synthetics_describeCanariesLastRunCmd = &cobra.Command{
	Use:   "describe-canaries-last-run",
	Short: "Use this operation to see information from the most recent run of each canary that you have created.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(synthetics_describeCanariesLastRunCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(synthetics_describeCanariesLastRunCmd).Standalone()

		synthetics_describeCanariesLastRunCmd.Flags().String("browser-type", "", "The type of browser to use for the canary run.")
		synthetics_describeCanariesLastRunCmd.Flags().String("max-results", "", "Specify this parameter to limit how many runs are returned each time you use the `DescribeLastRun` operation.")
		synthetics_describeCanariesLastRunCmd.Flags().String("names", "", "Use this parameter to return only canaries that match the names that you specify here.")
		synthetics_describeCanariesLastRunCmd.Flags().String("next-token", "", "A token that indicates that there is more data available.")
	})
	syntheticsCmd.AddCommand(synthetics_describeCanariesLastRunCmd)
}
