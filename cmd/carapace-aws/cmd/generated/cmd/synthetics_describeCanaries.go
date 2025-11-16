package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var synthetics_describeCanariesCmd = &cobra.Command{
	Use:   "describe-canaries",
	Short: "This operation returns a list of the canaries in your account, along with full details about each canary.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(synthetics_describeCanariesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(synthetics_describeCanariesCmd).Standalone()

		synthetics_describeCanariesCmd.Flags().String("max-results", "", "Specify this parameter to limit how many canaries are returned each time you use the `DescribeCanaries` operation.")
		synthetics_describeCanariesCmd.Flags().String("names", "", "Use this parameter to return only canaries that match the names that you specify here.")
		synthetics_describeCanariesCmd.Flags().String("next-token", "", "A token that indicates that there is more data available.")
	})
	syntheticsCmd.AddCommand(synthetics_describeCanariesCmd)
}
