# Real Debrid CLI

This tool is a simple command-line interface for Real Debrid. You can use it to manage and download files instead of the website.

## Installation

Install the latest binary from the [releases section](https://github.com/cgund98/realdebrid-cli/releases).

```bash
# Extract tarfile
tar -xf realdebrid-cli-v0.0.7-linux-amd64.tar.gz

# Install somewhere in $PATH
sudo mv realdebrid /usr/local/bin/realdebrid
```

## Example Usage

### Download from a restricted link
```bash
# Set API Token
export REAL_DEBRID_API_TOKEN="myapitoken"

# set file to download
export LINK="https://link.to/my/file.zip"

# Download a file or folder and store results to ~/Downloads/debrid
realdebrid downloads fetch $LINK \
    -o $HOME/Downloads/debrid

# Alternatively set token with a flag
realdebrid downloads fetch $LINK \
    -o $HOME/Downloads/debrid
    --token my-api-token
```

### Check if a restricted link is valid and downloadable
```bash
realdebrid downloads check $LINK
```

### List all downloads
```bash
realdebrid downloads list
```

### Remove all downloads
```bash
realdebrid downloads clean
```