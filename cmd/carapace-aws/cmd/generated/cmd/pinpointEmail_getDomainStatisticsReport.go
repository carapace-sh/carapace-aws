package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointEmail_getDomainStatisticsReportCmd = &cobra.Command{
	Use:   "get-domain-statistics-report",
	Short: "Retrieve inbox placement and engagement rates for the domains that you use to send email.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointEmail_getDomainStatisticsReportCmd).Standalone()

	pinpointEmail_getDomainStatisticsReportCmd.Flags().String("domain", "", "The domain that you want to obtain deliverability metrics for.")
	pinpointEmail_getDomainStatisticsReportCmd.Flags().String("end-date", "", "The last day (in Unix time) that you want to obtain domain deliverability metrics for.")
	pinpointEmail_getDomainStatisticsReportCmd.Flags().String("start-date", "", "The first day (in Unix time) that you want to obtain domain deliverability metrics for.")
	pinpointEmail_getDomainStatisticsReportCmd.MarkFlagRequired("domain")
	pinpointEmail_getDomainStatisticsReportCmd.MarkFlagRequired("end-date")
	pinpointEmail_getDomainStatisticsReportCmd.MarkFlagRequired("start-date")
	pinpointEmailCmd.AddCommand(pinpointEmail_getDomainStatisticsReportCmd)
}
