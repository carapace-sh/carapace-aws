package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticbeanstalk_checkDnsavailabilityCmd = &cobra.Command{
	Use:   "check-dnsavailability",
	Short: "Checks if the specified CNAME is available.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticbeanstalk_checkDnsavailabilityCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticbeanstalk_checkDnsavailabilityCmd).Standalone()

		elasticbeanstalk_checkDnsavailabilityCmd.Flags().String("cnameprefix", "", "The prefix used when this CNAME is reserved.")
		elasticbeanstalk_checkDnsavailabilityCmd.MarkFlagRequired("cnameprefix")
	})
	elasticbeanstalkCmd.AddCommand(elasticbeanstalk_checkDnsavailabilityCmd)
}
