package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datasync_describeLocationHdfsCmd = &cobra.Command{
	Use:   "describe-location-hdfs",
	Short: "Provides details about how an DataSync transfer location for a Hadoop Distributed File System (HDFS) is configured.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datasync_describeLocationHdfsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datasync_describeLocationHdfsCmd).Standalone()

		datasync_describeLocationHdfsCmd.Flags().String("location-arn", "", "Specifies the Amazon Resource Name (ARN) of the HDFS location.")
		datasync_describeLocationHdfsCmd.MarkFlagRequired("location-arn")
	})
	datasyncCmd.AddCommand(datasync_describeLocationHdfsCmd)
}
