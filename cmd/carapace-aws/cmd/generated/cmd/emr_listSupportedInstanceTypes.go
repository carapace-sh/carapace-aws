package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_listSupportedInstanceTypesCmd = &cobra.Command{
	Use:   "list-supported-instance-types",
	Short: "A list of the instance types that Amazon EMR supports.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_listSupportedInstanceTypesCmd).Standalone()

	emr_listSupportedInstanceTypesCmd.Flags().String("marker", "", "The pagination token that marks the next set of results to retrieve.")
	emr_listSupportedInstanceTypesCmd.Flags().String("release-label", "", "The Amazon EMR release label determines the [versions of open-source application packages](https://docs.aws.amazon.com/emr/latest/ReleaseGuide/emr-release-app-versions-6.x.html) that Amazon EMR has installed on the cluster.")
	emr_listSupportedInstanceTypesCmd.MarkFlagRequired("release-label")
	emrCmd.AddCommand(emr_listSupportedInstanceTypesCmd)
}
