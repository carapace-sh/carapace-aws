package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var efs_describeAccessPointsCmd = &cobra.Command{
	Use:   "describe-access-points",
	Short: "Returns the description of a specific Amazon EFS access point if the `AccessPointId` is provided.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(efs_describeAccessPointsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(efs_describeAccessPointsCmd).Standalone()

		efs_describeAccessPointsCmd.Flags().String("access-point-id", "", "(Optional) Specifies an EFS access point to describe in the response; mutually exclusive with `FileSystemId`.")
		efs_describeAccessPointsCmd.Flags().String("file-system-id", "", "(Optional) If you provide a `FileSystemId`, EFS returns all access points for that file system; mutually exclusive with `AccessPointId`.")
		efs_describeAccessPointsCmd.Flags().String("max-results", "", "(Optional) When retrieving all access points for a file system, you can optionally specify the `MaxItems` parameter to limit the number of objects returned in a response.")
		efs_describeAccessPointsCmd.Flags().String("next-token", "", "`NextToken` is present if the response is paginated.")
	})
	efsCmd.AddCommand(efs_describeAccessPointsCmd)
}
