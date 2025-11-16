package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_describeResizeCmd = &cobra.Command{
	Use:   "describe-resize",
	Short: "Returns information about the last resize operation for the specified cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_describeResizeCmd).Standalone()

	redshift_describeResizeCmd.Flags().String("cluster-identifier", "", "The unique identifier of a cluster whose resize progress you are requesting.")
	redshift_describeResizeCmd.MarkFlagRequired("cluster-identifier")
	redshiftCmd.AddCommand(redshift_describeResizeCmd)
}
