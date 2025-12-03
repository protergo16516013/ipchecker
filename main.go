package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/protergo16516013/abuseipdb"
	conf "github.com/protergo16516013/config"
)

//go:embed helper/usage
var usage string

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	config := conf.New()
	action := os.Args[1]

	if action == "config" {
		config.Setup()
		config.Save()
		os.Exit(0)
	} else {
		config.Load()
	}

	client := abuseipdb.NewClient(config.Apikey)

	switch action {
	case "check":
		handleCheck(client, os.Args[2:])
	case "report":
		handleReport(client, os.Args[2:])
	case "reports":
		handleReports(client, os.Args[2:])
	case "blacklist":
		handleBlacklist(client, os.Args[2:])
	case "check-block":
		handleCheckBlock(client, os.Args[2:])
	case "bulk-report":
		handleBulkReport(client, os.Args[2:])
	case "clear-address":
		handleClearAddress(client, os.Args[2:])
	default:
		fmt.Printf("Unknown action: %s\n", action)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println(usage)
}

func handleCheck(client *abuseipdb.Client, args []string) {
	fs := flag.NewFlagSet("check", flag.ExitOnError)
	ip := fs.String("ip", "", "IP address to check (required)")
	maxAge := fs.Int("max-age", 0, "Max age in days")
	verbose := fs.Bool("verbose", false, "Verbose output")

	fs.Parse(args)
	if *ip == "" {
		log.Fatal("--ip flag is required")
	}

	response, err := client.Check(*ip, *maxAge, *verbose)
	if err != nil {
		log.Fatal(err)
	}

	client.PrettyPrint(response)
}

func handleReport(client *abuseipdb.Client, args []string) {
	// fs := flag.NewFlagSet("report", flag.ExitOnError)
	// ip := fs.String("ip", "", "IP address to report (required)")
	// comment := fs.String("comment", "", "Report comment")
	// categories := fs.String("categories", "", "Comma-separated categories")

	// fs.Parse(args)

	// if *ip == "" {
	// 	log.Fatal("--ip flag is required")
	// }

	// // TODO: Implement Report method in client
	// fmt.Printf("Reporting IP: %s\n", *ip)
	// fmt.Printf("Comment: %s\n", *comment)
	// fmt.Printf("Categories: %s\n", *categories)
	println("Report function is still under construction!")
}

func handleReports(client *abuseipdb.Client, args []string) {
	fs := flag.NewFlagSet("reports", flag.ExitOnError)
	ip := fs.String("ip", "", "IP address (required)")
	page := fs.Int("page", 0, "Page number")
	perPage := fs.Int("per-page", 25, "Results per page")

	fs.Parse(args)

	if *ip == "" {
		log.Fatal("--ip flag is required")
	}

	response, err := client.Reports(*ip, *page, *perPage)
	if err != nil {
		log.Fatal(err)
	}

	client.PrettyPrint(response)
}

func handleBlacklist(client *abuseipdb.Client, args []string) {
	// fs := flag.NewFlagSet("blacklist", flag.ExitOnError)
	// plaintext := fs.Bool("plaintext", false, "Output as plaintext")

	// fs.Parse(args)

	// // TODO: Implement Blacklist method in client
	// fmt.Printf("Getting blacklist (plaintext: %v)\n", *plaintext)
	println("Blacklist function is still under construction!")
}

func handleCheckBlock(client *abuseipdb.Client, args []string) {
	// fs := flag.NewFlagSet("check-block", flag.ExitOnError)
	// network := fs.String("network", "", "Network address (required)")

	// fs.Parse(args)

	// if *network == "" {
	// 	log.Fatal("--network flag is required")
	// }

	// // TODO: Implement CheckBlock method in client
	// fmt.Printf("Checking network block: %s\n", *network)
	println("Check-Block function is still under construction!")
}

func handleBulkReport(client *abuseipdb.Client, args []string) {
	// fs := flag.NewFlagSet("bulk-report", flag.ExitOnError)
	// file := fs.String("file", "", "CSV file with reports (required)")

	// fs.Parse(args)

	// if *file == "" {
	// 	log.Fatal("--file flag is required")
	// }

	// // TODO: Implement BulkReport method in client
	// fmt.Printf("Bulk reporting from file: %s\n", *file)
	// println("Bulk-Report function is still under construction!")
}

func handleClearAddress(client *abuseipdb.Client, args []string) {
	// fs := flag.NewFlagSet("clear-address", flag.ExitOnError)
	// ip := fs.String("ip", "", "IP address (required)")

	// fs.Parse(args)

	// if *ip == "" {
	// 	log.Fatal("--ip flag is required")
	// }

	// // TODO: Implement ClearAddress method in client
	// fmt.Printf("Clearing address: %s\n", *ip)
	println("Clear-Address function is still under construction!")
}
