package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_getDeployablePatchSnapshotForInstanceCmd = &cobra.Command{
	Use:   "get-deployable-patch-snapshot-for-instance",
	Short: "Retrieves the current snapshot for the patch baseline the managed node uses.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_getDeployablePatchSnapshotForInstanceCmd).Standalone()

	ssm_getDeployablePatchSnapshotForInstanceCmd.Flags().String("baseline-override", "", "Defines the basic information about a patch baseline override.")
	ssm_getDeployablePatchSnapshotForInstanceCmd.Flags().String("instance-id", "", "The ID of the managed node for which the appropriate patch snapshot should be retrieved.")
	ssm_getDeployablePatchSnapshotForInstanceCmd.Flags().Bool("no-use-s3-dual-stack-endpoint", false, "Specifies whether to use S3 dualstack endpoints for the patch snapshot download URL.")
	ssm_getDeployablePatchSnapshotForInstanceCmd.Flags().String("snapshot-id", "", "The snapshot ID provided by the user when running `AWS-RunPatchBaseline`.")
	ssm_getDeployablePatchSnapshotForInstanceCmd.Flags().Bool("use-s3-dual-stack-endpoint", false, "Specifies whether to use S3 dualstack endpoints for the patch snapshot download URL.")
	ssm_getDeployablePatchSnapshotForInstanceCmd.MarkFlagRequired("instance-id")
	ssm_getDeployablePatchSnapshotForInstanceCmd.Flag("no-use-s3-dual-stack-endpoint").Hidden = true
	ssm_getDeployablePatchSnapshotForInstanceCmd.MarkFlagRequired("snapshot-id")
	ssmCmd.AddCommand(ssm_getDeployablePatchSnapshotForInstanceCmd)
}
