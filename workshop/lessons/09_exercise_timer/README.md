# cli timer

write a cli timer

e.g.

```
timer 10
```
- will wait 10 seconds


## extended
- flag package

```
	tlsDisabled := flag.Bool("tlsDisabled", false, "disabled tls for the server")
	limitMemory := flag.String("limitMemory", "1G", "memory limit (default 1G)")
	limitCPU := flag.String("limitCPU", "0.5", "cpu limit (default 0.5 cores)")
	requestMemory := flag.String("requestMemory", "512M", "memory request (default 1G)")
	requestCPU := flag.String("requestCPU", "0.05", "cpu request (default 0.1 cores)")
	addr := flag.String("addr", ":8083", "address to bind to")
	sslCert := flag.String("sslCert", "/certs/ssl-cert.pem", "address to bind to")
	sslKey := flag.String("sslKey", "/certs/ssl-key.pem", "address to bind to")
	dryRun := flag.Bool("dry-run", false, "enables dry-run mode, always returning success AdmissionReview")
```

```
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Please enter a duration")
		os.Exit(2)
	} else if len(args) == 1 {
		arg1 := args[0]
		duration, err := time.ParseDuration(arg1)
		if err != nil {
			panic(err)
			os.Exit(2)
		}
		MyTimer(duration, "")
		os.Exit(0)
	} else if len(args) == 2 {
```