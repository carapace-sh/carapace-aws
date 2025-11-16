package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_getDomainStatisticsReportCmd = &cobra.Command{
	Use:   "get-domain-statistics-report",
	Short: "Retrieve inbox placement and engagement rates for the domains that you use to send email.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_getDomainStatisticsReportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sesv2_getDomainStatisticsReportCmd).Standalone()

		sesv2_getDomainStatisticsReportCmd.Flags().String("domain", "", "The domain that you want to obtain deliverability metrics for.")
		sesv2_getDomainStatisticsReportCmd.Flags().String("end-date", "", "The last day (in Unix time) that you want to obtain domain deliverability metrics for.")
		sesv2_getDomainStatisticsReportCmd.Flags().String("start-date", "", "The first day (in Unix time) that you want to obtain domain deliverability metrics for.")
		sesv2_getDomainStatisticsReportCmd.MarkFlagRequired("domain")
		sesv2_getDomainStatisticsReportCmd.MarkFlagRequired("end-date")
		sesv2_getDomainStatisticsReportCmd.MarkFlagRequired("start-date")
	})
	sesv2Cmd.AddCommand(sesv2_getDomainStatisticsReportCmd)
}
