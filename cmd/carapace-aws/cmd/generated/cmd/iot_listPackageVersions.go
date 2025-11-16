package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listPackageVersionsCmd = &cobra.Command{
	Use:   "list-package-versions",
	Short: "Lists the software package versions associated to the account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listPackageVersionsCmd).Standalone()

	iot_listPackageVersionsCmd.Flags().String("max-results", "", "The maximum number of results to return at one time.")
	iot_listPackageVersionsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	iot_listPackageVersionsCmd.Flags().String("package-name", "", "The name of the target software package.")
	iot_listPackageVersionsCmd.Flags().String("status", "", "The status of the package version.")
	iot_listPackageVersionsCmd.MarkFlagRequired("package-name")
	iotCmd.AddCommand(iot_listPackageVersionsCmd)
}
