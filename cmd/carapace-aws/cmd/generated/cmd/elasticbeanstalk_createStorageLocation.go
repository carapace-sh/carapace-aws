package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticbeanstalk_createStorageLocationCmd = &cobra.Command{
	Use:   "create-storage-location",
	Short: "Creates a bucket in Amazon S3 to store application versions, logs, and other files used by Elastic Beanstalk environments.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticbeanstalk_createStorageLocationCmd).Standalone()

	elasticbeanstalkCmd.AddCommand(elasticbeanstalk_createStorageLocationCmd)
}
