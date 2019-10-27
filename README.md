# imagecheck-for-gocon

```bash
$ go get github.com/tomoyamachi/imagecheck-for-gocon
$ imagecheck-for-gocon nginx:latest 
2019/10/27 21:59:02 Start ScanImage...
2019/10/27 21:59:03 etc/nginx/nginx.conf: Expect log format contains "ltsv" but "main;"
exit status 1

$ imagecheck-for-gocon golang:1.13
2019/10/27 21:58:39 Start ScanImage...
2019/10/27 21:58:40 no nginx.conf files
2019/10/27 21:58:40 Finish ScanImage...

$ # create ltsv log_format containers
$ cd $GOPATH/src/github.com/tomoyamachi/imagecheck-for-gocon/containers/ltsv/ && docker build -t nginx-ltsv:v1 .
$ imagecheck-for-gocon nginx-ltsv:v1
2019/10/27 21:58:50 Start ScanImage...
2019/10/27 21:58:50 Finish ScanImage...
```

This is a sample repository for container image scanner. The design of the repository is very simple.

If you try to create better design, please check [aquasecurity/trivy](https://github.com/aquasecurity/trivy) and [goodwithtech/dockle](https://github.com/goodwithtech/dockle).