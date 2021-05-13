# xch-downloader - download plots from xch-plotter.com

Helper tool (daemon) to continuously download plots from https://xch-plotter.com service and mark them for deletion.

# Features

* parallel downloads
* supports resuming download if a file already exists in the working directory
* sends DELETE request to the server once plot is downloader
* can run continuously as a daemon

# Installation

Go to [RELEASES](https://github.com/zytek/xch-downloader/releases) page and download the newest version. You can build this from source using simple `go build` command

# Usage

```
Usage of ./xch-downloader:
  -d	daemonize (check and download continuously
  -delete
    	send DELETE requests after downloading plot
  -dir string
    	plot directory (default ".")
  -key string
    	farmer key
  -n int
    	download n files simultaneously (default 10)
  -url string
    	plotter url (default "https://xch-plotter.com")
```

Basic usage. Download all plots from the server:

```
./xch-downloader -key FARMER_KEY -dir .
```

Download all plots continuously. Delete them from server once done:

```
./xch-downloader -key FARMER_KEY -dir . -delete -d
```

