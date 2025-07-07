package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"os"

	"github.com/hcarriz/embedhash"
)

func main() {
	var (
		hashFlag   = flag.String("hash", "md5", "hash algorithm: md5 or sha256")
		outFlag    = flag.String("out", "embedhash.go", "output file name")
		varFlag    = flag.String("var", "HashesForEmbedded", "output variable name")
		dryRunFlag = flag.Bool("dry", false, "dry run")
	)

	flag.Parse()

	var opts []embedhash.Option

	if notempty(outFlag) {
		opts = append(opts, embedhash.OutputFileName(*outFlag))
	}
	if notempty(varFlag) {
		opts = append(opts, embedhash.OutputVarName(*varFlag))
	}

	if notempty(hashFlag) {
		switch *hashFlag {
		case "sha256":
			opts = append(opts, embedhash.Hasher(sha256.New()))
		}
	}

	results, err := embedhash.New(flag.Arg(0), opts...)
	fail(err)

	for _, x := range results {
		if notempty(dryRunFlag) {
			fmt.Printf("[DRY RUN] %s\n", x.Filename)
			fail(x.Save(os.Stdout))
		} else {

			fmt.Printf("writing to %s\n", x.Filename)

			f, err := os.Create(x.Filename)
			fail(err)

			defer f.Close()

			fail(x.Save(f))

		}
	}

}

func notempty[T comparable](in *T) bool {
	var empty T

	if in == nil {
		return false
	}

	return *in != empty

}

func fail(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
