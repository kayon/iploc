package main

import (
	"fmt"
	"os"

	"github.com/kayon/iploc"
	flag "github.com/spf13/pflag"
)

var (
	ip       string
	version  bool
	detailed bool
	help     bool
)

func init() {

	flag.BoolVarP(&help, "help", "h", false, "this help")
	flag.BoolVarP(&version, "version", "v", false, "show version and exit")
	flag.BoolVarP(&detailed, "detailed", "d", false, "show the detailed results")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: iploc [IPv4] [arguments]\n%12s [command]\nOptions:\n", "iploc")
		flag.PrintDefaults()
	}
	flag.Parse()

	if version {
		fmt.Fprintln(os.Stderr, "iploc", iploc.Version)
		os.Exit(0)
	} else if flag.Arg(0) == "" || help {
		flag.Usage()
		if help {
			os.Exit(0)
		}
		os.Exit(1)
	}
	ip = flag.Arg(0)

}

func main() {
	var err error
	detail := IPLoc.Find(ip)
	if detail == nil {
		fmt.Fprintf(os.Stderr, "iploc: invalid IP %q\n", ip)
		os.Exit(1)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if !detailed {
		fmt.Printf("%s %s\n", detail.IP, detail.String())
	} else {
		fmt.Printf("IP: %s\n网段: %s - %s\n位置: %s\n", detail.IP, detail.Start, detail.End, detail)
	}
}
