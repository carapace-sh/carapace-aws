package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrailDataCmd = &cobra.Command{
	Use:   "cloudtrail-data",
	Short: "The CloudTrail Data Service lets you ingest events into CloudTrail from any source in your hybrid environments, such as in-house or SaaS applications hosted on-premises or in the cloud, virtual machines, or containers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrailDataCmd).Standalone()

	rootCmd.AddCommand(cloudtrailDataCmd)
}
