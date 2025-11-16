package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_createCrawlerCmd = &cobra.Command{
	Use:   "create-crawler",
	Short: "Creates a new crawler with specified targets, role, configuration, and optional schedule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_createCrawlerCmd).Standalone()

	glue_createCrawlerCmd.Flags().String("classifiers", "", "A list of custom classifiers that the user has registered.")
	glue_createCrawlerCmd.Flags().String("configuration", "", "Crawler configuration information.")
	glue_createCrawlerCmd.Flags().String("crawler-security-configuration", "", "The name of the `SecurityConfiguration` structure to be used by this crawler.")
	glue_createCrawlerCmd.Flags().String("database-name", "", "The Glue database where results are written, such as: `arn:aws:daylight:us-east-1::database/sometable/*`.")
	glue_createCrawlerCmd.Flags().String("description", "", "A description of the new crawler.")
	glue_createCrawlerCmd.Flags().String("lake-formation-configuration", "", "Specifies Lake Formation configuration settings for the crawler.")
	glue_createCrawlerCmd.Flags().String("lineage-configuration", "", "Specifies data lineage configuration settings for the crawler.")
	glue_createCrawlerCmd.Flags().String("name", "", "Name of the new crawler.")
	glue_createCrawlerCmd.Flags().String("recrawl-policy", "", "A policy that specifies whether to crawl the entire dataset again, or to crawl only folders that were added since the last crawler run.")
	glue_createCrawlerCmd.Flags().String("role", "", "The IAM role or Amazon Resource Name (ARN) of an IAM role used by the new crawler to access customer resources.")
	glue_createCrawlerCmd.Flags().String("schedule", "", "A `cron` expression used to specify the schedule (see [Time-Based Schedules for Jobs and Crawlers](https://docs.aws.amazon.com/glue/latest/dg/monitor-data-warehouse-schedule.html). For example, to run something every day at 12:15 UTC, you would specify: `cron(15 12 * * ?")
	glue_createCrawlerCmd.Flags().String("schema-change-policy", "", "The policy for the crawler's update and deletion behavior.")
	glue_createCrawlerCmd.Flags().String("table-prefix", "", "The table prefix used for catalog tables that are created.")
	glue_createCrawlerCmd.Flags().String("tags", "", "The tags to use with this crawler request.")
	glue_createCrawlerCmd.Flags().String("targets", "", "A list of collection of targets to crawl.")
	glue_createCrawlerCmd.MarkFlagRequired("name")
	glue_createCrawlerCmd.MarkFlagRequired("role")
	glue_createCrawlerCmd.MarkFlagRequired("targets")
	glueCmd.AddCommand(glue_createCrawlerCmd)
}
