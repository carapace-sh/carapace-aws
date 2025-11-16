package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_describeInstanceProfilesCmd = &cobra.Command{
	Use:   "describe-instance-profiles",
	Short: "Returns a paginated list of instance profiles for your account in the current region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_describeInstanceProfilesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_describeInstanceProfilesCmd).Standalone()

		dms_describeInstanceProfilesCmd.Flags().String("filters", "", "Filters applied to the instance profiles described in the form of key-value pairs.")
		dms_describeInstanceProfilesCmd.Flags().String("marker", "", "Specifies the unique pagination token that makes it possible to display the next page of results.")
		dms_describeInstanceProfilesCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	})
	dmsCmd.AddCommand(dms_describeInstanceProfilesCmd)
}
